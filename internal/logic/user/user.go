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

	nickname := strings.TrimSpace(in.Nickname)
	if nickname == "" {
		nickname = target.Username
	}

	if _, err := dao.User.Ctx(ctx).WherePri(in.ID).Data(g.Map{
		cols.Nickname: nickname,
		cols.Role:     in.Role,
		cols.Status:   in.Status,
	}).Update(); err != nil {
		return 0, err
	}
	return in.ID, nil
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
