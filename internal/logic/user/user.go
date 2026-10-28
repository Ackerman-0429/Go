// Package user 实现用户模块的业务逻辑：注册、登录、查询。
package user

import (
	"context"
	"fmt"
	"strings"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"golang.org/x/crypto/bcrypt"

	"goshop/internal/consts"
	"goshop/internal/dao"
	"goshop/internal/model"
	"goshop/internal/model/do"
	"goshop/internal/model/entity"
	"goshop/internal/service"
)

// sUser 是 service.IUser 的实现
type sUser struct{}

func init() {
	service.RegisterUser(&sUser{})
}

// Register 注册新用户。密码用 bcrypt 哈希后入库，绝不存明文。
func (s *sUser) Register(ctx context.Context, username, password, nickname string) (uint64, error) {
	count, err := dao.User.Ctx(ctx).Where(dao.User.Columns().Username, username).Count()
	if err != nil {
		return 0, err
	}
	if count > 0 {
		return 0, gerror.NewCode(gcode.CodeValidationFailed, "用户名已被占用")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	if nickname == "" {
		nickname = username
	}

	result, err := dao.User.Ctx(ctx).Data(do.User{
		Username: username,
		Password: string(hashed),
		Nickname: nickname,
		Status:   1,
	}).Insert()
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(id), nil
}

// Login 登录校验。用户名不存在和密码错误统一提示，避免暴露账号是否存在。
func (s *sUser) Login(ctx context.Context, username, password string) (*entity.User, error) {
	var user *entity.User
	if err := dao.User.Ctx(ctx).Where(dao.User.Columns().Username, username).Scan(&user); err != nil {
		return nil, err
	}
	if user == nil {
		return nil, gerror.NewCode(gcode.CodeNotAuthorized, "用户名或密码错误")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, gerror.NewCode(gcode.CodeNotAuthorized, "用户名或密码错误")
	}

	if user.Status != 1 {
		return nil, gerror.NewCode(gcode.CodeNotAuthorized, "账号已被禁用")
	}

	return user, nil
}

// GetById 按主键查询用户
func (s *sUser) GetById(ctx context.Context, id uint64) (*entity.User, error) {
	var user *entity.User
	if err := dao.User.Ctx(ctx).WherePri(id).Scan(&user); err != nil {
		return nil, err
	}
	if user == nil {
		return nil, gerror.NewCode(gcode.CodeNotFound, "用户不存在")
	}
	return user, nil
}

// AdminList 后台用户列表。支持关键字（用户名/昵称/邮箱/手机号）、角色、状态筛选，
// 并附带订单数与累计消费，方便判断这个账号值不值得留着。
func (s *sUser) AdminList(ctx context.Context, in model.UserAdminListInput) ([]model.UserAdminItem, int, error) {
	cols := dao.User.Columns()
	m := dao.User.Ctx(ctx)

	if in.Status >= 0 {
		m = m.Where(cols.Status, in.Status)
	}
	if in.Role >= 0 {
		m = m.Where(cols.Role, in.Role)
	}
	if in.Keyword != "" {
		like := "%" + in.Keyword + "%"
		m = m.Where(
			"("+cols.Username+" LIKE ? OR "+cols.Nickname+" LIKE ? OR "+cols.Email+" LIKE ? OR "+cols.Phone+" LIKE ?)",
			like, like, like, like,
		)
	}

	var (
		users []entity.User
		total int
	)
	// 后台按 ID 倒序，最新注册的排前面
	if err := m.Page(in.Page, in.PageSize).
		OrderDesc(cols.Id).
		ScanAndCount(&users, &total, false); err != nil {
		return nil, 0, err
	}
	if len(users) == 0 {
		return []model.UserAdminItem{}, total, nil
	}

	ids := make([]uint64, 0, len(users))
	for _, u := range users {
		ids = append(ids, u.Id)
	}
	stats, err := orderStatsByUser(ctx, ids)
	if err != nil {
		return nil, 0, err
	}

	out := make([]model.UserAdminItem, 0, len(users))
	for _, u := range users {
		item := model.UserAdminItem{
			ID:        u.Id,
			Username:  u.Username,
			Nickname:  u.Nickname,
			Email:     u.Email,
			Phone:     u.Phone,
			Role:      u.Role,
			Status:    u.Status,
			CreatedAt: u.CreatedAt,
		}
		if st := stats[u.Id]; st != nil {
			item.OrderCount = st.OrderCount
			item.TotalSpent = st.TotalSpent
		}
		out = append(out, item)
	}
	return out, total, nil
}

// orderStats 单个用户的订单汇总
type orderStats struct {
	OrderCount int
	TotalSpent float64
}

// orderStatsByUser 按用户批量汇总订单数，一次查询解决，避免 N+1。
// 订单数含全部状态；消费额只统计已付款及之后，待付款和已取消不算实际收入。
func orderStatsByUser(ctx context.Context, userIDs []uint64) (map[uint64]*orderStats, error) {
	var rows []struct {
		UserId uint64  `orm:"user_id"`
		Total  int     `orm:"total"`
		Amount float64 `orm:"amount"`
	}

	cols := dao.Orders.Columns()
	paid := fmt.Sprintf("%d,%d,%d",
		consts.OrderStatusPaid, consts.OrderStatusShipped, consts.OrderStatusFinished)
	amountExpr := fmt.Sprintf(
		"SUM(CASE WHEN %s IN (%s) THEN %s ELSE 0 END) AS amount",
		cols.Status, paid, cols.TotalAmount,
	)

	if err := dao.Orders.Ctx(ctx).
		Fields(cols.UserId, "COUNT(*) AS total", amountExpr).
		WhereIn(cols.UserId, userIDs).
		Group(cols.UserId).
		Scan(&rows); err != nil {
		return nil, err
	}

	out := make(map[uint64]*orderStats, len(rows))
	for _, r := range rows {
		out[r.UserId] = &orderStats{OrderCount: r.Total, TotalSpent: r.Amount}
	}
	return out, nil
}

// UpdateProfile 用户自助修改资料：昵称、邮箱、手机号。
// 角色、状态、用户名都不在可改范围内，所以这个接口即便被构造调用也越不了权。
func (s *sUser) UpdateProfile(ctx context.Context, in model.UserProfileInput) (*entity.User, error) {
	cols := dao.User.Columns()

	var user *entity.User
	if err := dao.User.Ctx(ctx).WherePri(in.ID).Scan(&user); err != nil {
		return nil, err
	}
	if user == nil {
		return nil, gerror.NewCode(gcode.CodeNotFound, "用户不存在")
	}
	// 这个接口不在 Auth 中间件之后（登录/注册也在同一个控制器上），
	// 所以禁用状态要在这里自己拦，避免老 Session 还能改资料。
	if user.Status != consts.UserStatusNormal {
		return nil, gerror.NewCode(gcode.CodeNotAuthorized, "账号已被禁用")
	}

	email, phone := strings.TrimSpace(in.Email), strings.TrimSpace(in.Phone)
	// 排除自己：原样提交未改动的邮箱/手机号不算重复
	if err := checkContactUnique(ctx, email, phone, user.Id); err != nil {
		return nil, err
	}

	nickname := strings.TrimSpace(in.Nickname)
	if nickname == "" {
		nickname = user.Username
	}

	if _, err := dao.User.Ctx(ctx).WherePri(user.Id).Data(g.Map{
		cols.Nickname: nickname,
		cols.Email:    nullableContact(email),
		cols.Phone:    nullableContact(phone),
	}).Update(); err != nil {
		return nil, err
	}

	// 回填返回值，前端拿到就能刷新顶栏昵称，省一次 /user/info
	user.Nickname = nickname
	user.Email = email
	user.Phone = phone
	return user, nil
}

// ChangePassword 用户自助改密码。必须验原密码——只有会话还不够，
// 否则别人碰一下没锁屏的浏览器就能把密码换掉。
// 改完不动 Session：当前这台设备继续可用，其他设备下次登录才需要新密码。
func (s *sUser) ChangePassword(ctx context.Context, in model.UserPasswordInput) error {
	cols := dao.User.Columns()

	var user *entity.User
	if err := dao.User.Ctx(ctx).WherePri(in.ID).Scan(&user); err != nil {
		return err
	}
	if user == nil {
		return gerror.NewCode(gcode.CodeNotFound, "用户不存在")
	}
	if user.Status != consts.UserStatusNormal {
		return gerror.NewCode(gcode.CodeNotAuthorized, "账号已被禁用")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(in.OldPassword)); err != nil {
		return gerror.NewCode(gcode.CodeValidationFailed, "原密码不正确")
	}

	// 新旧相同等于没改，直接拦掉，免得用户以为改成功了
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(in.NewPassword)) == nil {
		return gerror.NewCode(gcode.CodeValidationFailed, "新密码不能与原密码相同")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(in.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = dao.User.Ctx(ctx).WherePri(user.Id).Data(g.Map{cols.Password: string(hashed)}).Update()
	return err
}

// AdminSave 编辑用户：昵称、角色、状态。
// operatorID 是当前登录的管理员，用来拦住「把自己锁在门外」的操作。
func (s *sUser) AdminSave(ctx context.Context, in model.UserSaveInput, operatorID uint64) (uint64, error) {
	cols := dao.User.Columns()

	var target *entity.User
	if err := dao.User.Ctx(ctx).WherePri(in.ID).Scan(&target); err != nil {
		return 0, err
	}
	if target == nil {
		return 0, gerror.NewCode(gcode.CodeNotFound, "用户不存在")
	}

	// 管理员给自己降级或禁用，会立刻失去后台权限，这里直接拦住
	if target.Id == operatorID && (in.Role != consts.UserRoleAdmin || in.Status != consts.UserStatusNormal) {
		return 0, gerror.NewCode(gcode.CodeInvalidParameter, "不能降级或禁用当前登录的账号")
	}

	// 至少保留一名可用管理员，否则以后没人能进后台
	demoting := target.Role == consts.UserRoleAdmin &&
		(in.Role != consts.UserRoleAdmin || in.Status != consts.UserStatusNormal)
	if demoting {
		n, err := dao.User.Ctx(ctx).
			Where(cols.Role, consts.UserRoleAdmin).
			Where(cols.Status, consts.UserStatusNormal).
			WhereNot(cols.Id, target.Id).
			Count()
		if err != nil {
			return 0, err
		}
		if n == 0 {
			return 0, gerror.NewCode(gcode.CodeInvalidParameter, "系统至少需要保留一名可用管理员")
		}
	}

	email, phone := strings.TrimSpace(in.Email), strings.TrimSpace(in.Phone)
	if err := checkContactUnique(ctx, email, phone, target.Id); err != nil {
		return 0, err
	}

	nickname := strings.TrimSpace(in.Nickname)
	if nickname == "" {
		nickname = target.Username
	}

	if _, err := dao.User.Ctx(ctx).WherePri(in.ID).Data(g.Map{
		cols.Nickname: nickname,
		cols.Email:    nullableContact(email),
		cols.Phone:    nullableContact(phone),
		cols.Role:     in.Role,
		cols.Status:   in.Status,
	}).Update(); err != nil {
		return 0, err
	}
	return in.ID, nil
}

// AdminCreate 后台直接新建用户。复用注册那套校验，额外放开角色与状态，
// 让管理员可以一次性建出「已启用的管理员」或「先建好但禁用」的账号。
func (s *sUser) AdminCreate(ctx context.Context, in model.UserCreateInput) (uint64, error) {
	cols := dao.User.Columns()
	username := strings.TrimSpace(in.Username)

	count, err := dao.User.Ctx(ctx).Where(cols.Username, username).Count()
	if err != nil {
		return 0, err
	}
	if count > 0 {
		return 0, gerror.NewCode(gcode.CodeValidationFailed, "用户名已被占用")
	}

	email, phone := strings.TrimSpace(in.Email), strings.TrimSpace(in.Phone)
	if err = checkContactUnique(ctx, email, phone, 0); err != nil {
		return 0, err
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	nickname := strings.TrimSpace(in.Nickname)
	if nickname == "" {
		nickname = username
	}

	result, err := dao.User.Ctx(ctx).Data(g.Map{
		cols.Username: username,
		cols.Password: string(hashed),
		cols.Nickname: nickname,
		cols.Email:    nullableContact(email),
		cols.Phone:    nullableContact(phone),
		cols.Role:     in.Role,
		cols.Status:   in.Status,
	}).Insert()
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(id), nil
}

// nullableContact 邮箱/手机号留空时返回 nil，落库即 NULL。
// 这两列是 UNIQUE 且允许 NULL：多个 NULL 互不冲突，而空字符串只能存在一个，
// 第二个不填邮箱的用户就会撞唯一键。
func nullableContact(v string) any {
	if v == "" {
		return nil
	}
	return v
}

// checkContactUnique 校验邮箱/手机号未被其他账号占用。
// excludeID 用于编辑场景排除自己，否则原样保存会被判成重复。
func checkContactUnique(ctx context.Context, email, phone string, excludeID uint64) error {
	cols := dao.User.Columns()

	if email != "" {
		m := dao.User.Ctx(ctx).Where(cols.Email, email)
		if excludeID > 0 {
			m = m.WhereNot(cols.Id, excludeID)
		}
		n, err := m.Count()
		if err != nil {
			return err
		}
		if n > 0 {
			return gerror.NewCode(gcode.CodeValidationFailed, "该邮箱已被其他账号使用")
		}
	}

	if phone != "" {
		m := dao.User.Ctx(ctx).Where(cols.Phone, phone)
		if excludeID > 0 {
			m = m.WhereNot(cols.Id, excludeID)
		}
		n, err := m.Count()
		if err != nil {
			return err
		}
		if n > 0 {
			return gerror.NewCode(gcode.CodeValidationFailed, "该手机号已被其他账号使用")
		}
	}
	return nil
}

// AdminResetPassword 管理员重置密码，不需要知道原密码。
func (s *sUser) AdminResetPassword(ctx context.Context, id uint64, password string) error {
	cols := dao.User.Columns()

	count, err := dao.User.Ctx(ctx).WherePri(id).Count()
	if err != nil {
		return err
	}
	if count == 0 {
		return gerror.NewCode(gcode.CodeNotFound, "用户不存在")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = dao.User.Ctx(ctx).WherePri(id).Data(g.Map{cols.Password: string(hashed)}).Update()
	return err
}
