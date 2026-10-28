# GoShop · GoFrame 网购系统（MVP）

基于 GoFrame v2 的电商后端 MVP，已跑通「注册 → 浏览商品 → 加购物车 → 下单 → 支付」闭环，
并附带两个原生 JS 页面：商城前台与后台管理。

## 环境要求

| 组件 | 版本 |
| --- | --- |
| Go | 1.27 |
| MySQL | 8.4 LTS（utf8mb4，时区 +08:00） |
| GoFrame CLI | v2.10.3（与框架版本保持一致） |

```bash
gf version     # 版本不一致时：go install github.com/gogf/gf/cmd/gf/v2@latest
```

## 快速开始

1. 建库并导入结构与演示数据：

```bash
mysql -h 127.0.0.1 -uroot -p gfstudy < manifest/sql/schema.sql
mysql -h 127.0.0.1 -uroot -p gfstudy < manifest/sql/seed.sql
```

2. 准备本地配置（含密码，不入库）：

```bash
cp manifest/config/config.example.yaml manifest/config/config.yaml
# 修改数据库账号密码
```

3. 启动：

```bash
go run .                 # 或双击 scripts/start-dev.bat
```

- 商城首页：http://127.0.0.1:8000/
- 后台管理：http://127.0.0.1:8000/admin.html
- 接口文档：http://127.0.0.1:8000/swagger/
- OpenAPI：http://127.0.0.1:8000/api.json

演示账号：`admin`（管理员）/ `alice`、`bob`（普通用户），密码 `123456`。

## 目录结构（GoFrame 官方分层）

```
api/                    接口契约：Req/Res + g.Meta 路由标签
  └── <模块>/v1/        各版本入参出参
internal/
  ├── cmd/              路由装配、中间件挂载
  ├── consts/           项目常量：订单状态、商品上下架状态
  ├── controller/       控制器：只做参数接收与响应组装
  ├── dao/              数据访问（gf gen dao 生成）
  ├── logic/            业务实现（init() 中向 service 注册）
  ├── middleware/       HTTP 中间件：登录态校验
  ├── model/            业务 DTO（跨层复用）
  │   ├── do/           写入模型（生成）
  │   └── entity/       表结构实体（生成）
  └── service/          业务接口（gf gen service 生成）
manifest/
  ├── config/           运行配置
  ├── deploy/           k8s 部署模板
  └── sql/              建表与种子数据
resource/public/        前端页面：index.html 商城、admin.html 后台
utility/                通用工具（金额分换算等）
```

依赖方向：`controller → service（接口） → logic（实现） → dao`，controller 不直接碰 dao。

## 接口一览

| 模块 | 接口 | 说明 |
| --- | --- | --- |
| 用户 | `POST /user/register`、`POST /user/login`、`POST /user/logout`、`GET /user/info` | 注册/登录写入 Session |
| 分类 | `GET /category/list` | 二级分类，深度优先排序 |
| 商品 | `GET /goods/list`、`GET /goods/detail` | 仅返回上架商品 |
| 购物车 | `POST /cart/add|update|delete|clear`、`GET /cart/list` | 需登录 |
| 订单 | `POST /order/create|pay|cancel`、`GET /order/list|detail` | 需登录 |
| 后台 | 见下方「后台管理」 | 需登录 + 管理员 |

统一响应体：`{"code":0,"message":"OK","data":{...}}`，由 `ghttp.MiddlewareHandlerResponse` 输出。

## 后台管理

入口 `http://127.0.0.1:8000/admin.html`，商城页右上角也会给管理员显示「管理后台」入口。

**权限模型**：`user.role` 字段（0 普通用户 / 1 管理员）。角色存在数据库里，不写 Session，
所以管理员被降权后立即生效，不用等会话过期。三道关口：

| 情形 | 返回 |
| --- | --- |
| 未登录 | `code: 61`（请先登录） |
| 已登录但非管理员 | `code: 40301`（需要管理员权限） |
| 管理员 | 正常数据 |

中间件挂载顺序在 `internal/cmd/cmd.go`：`MiddlewareHandlerResponse` → `Auth` → `Admin`。

`Auth` 不只看 Session 里有没有用户 ID，还会回查一次 `user` 表：账号被删除就清掉脏 Session，
被禁用就返回「账号已被禁用」，并把角色写进上下文供 `Admin` 复用。
代价是每个需要登录的请求多一次主键查询，换来的是**禁用/降权立即生效**——
否则管理员封了某个账号，对方拿着旧 Session 还能继续下单。

**接口**（统一前缀 `/admin`）：

| 方法 | 路径 | 说明 |
| --- | --- | --- |
| GET | `/admin/stats` | 概览：今日订单/销售额、待发货、累计订单、在售商品、用户数 |
| GET | `/admin/goods/list` | 商品列表，含下架商品，支持关键字/分类/状态筛选 |
| POST | `/admin/goods/save` | 新建或编辑商品（`id` 为空即新建） |
| POST | `/admin/goods/status` | 上下架 |
| POST | `/admin/goods/delete` | 删除商品（同时清理购物车引用，订单快照不受影响） |
| GET | `/admin/category/list` | 分类列表，含停用分类，带商品数与子分类数 |
| POST | `/admin/category/save` | 新建或编辑分类（`id` 为空即新建，仅支持两级） |
| POST | `/admin/category/delete` | 删除分类（有子分类或商品时拒绝，返回具体原因） |
| GET | `/admin/order/list` | 全站订单，支持状态/订单号/收货人筛选 |
| POST | `/admin/order/ship` | 发货：已付款 → 已发货 |
| POST | `/admin/order/finish` | 完成：已发货 → 已完成 |
| GET | `/admin/user/list` | 用户列表，支持关键字/角色/状态筛选，带订单数与累计消费 |
| POST | `/admin/user/create` | 新建用户：用户名、初始密码、昵称、邮箱、手机号、角色、状态 |
| POST | `/admin/user/save` | 编辑用户：昵称、邮箱、手机号、角色、启用/禁用 |
| POST | `/admin/user/reset-password` | 重置密码（管理员无需知道原密码） |

**给已有环境升级**：

```bash
# 旧库没有 role 字段时先加列
mysql -h 127.0.0.1 -u root -p gfstudy < manifest/sql/upgrade/20261028_user_role.sql

# 建立 admin 管理员账号，并收回 alice 的管理员权限
mysql -h 127.0.0.1 -u root -p gfstudy < manifest/sql/upgrade/20261028_admin_role.sql
```

第二个脚本用 `INSERT ... ON DUPLICATE KEY UPDATE`，所以 admin 已存在时不会覆盖它的密码，
只纠正 `role`/`status`；脚本里的哈希对应演示密码，仅适合本地开发环境。

手工指定管理员：`UPDATE user SET role = 1 WHERE username = 'xxx';`

**分类管理的约束**（都在 logic 层校验，不靠前端拦）：

- 只支持两级。父分类必须存在且自身是顶级分类；已经有子分类的分类不能被挂到别的分类下。
- 同一父级下不允许重名（避免前台出现两个一样的筛选项）。
- 删除前检查子分类与商品：有就拒绝并说明具体数量，避免产生孤儿分类或空分类的商品。
- 顶级分类的「商品数」会累加子分类，否则父分类永远显示 0 件，看起来像统计错误。
- 停用只是后台标记：前台 `/category/list` 不再返回，但商品仍保留在该分类下，后台也能看到。

**用户管理的约束**：

- **邮箱与手机号是可选但唯一的**。留空时写入 `NULL` 而不是空字符串——这两列是
  `UNIQUE` 且允许 `NULL`，多个 `NULL` 互不冲突；若存空串，第二个不填邮箱的账号就会撞唯一键。
  因此保存时统一走 `nullableContact()` 转换，编辑里清空字段能真正清掉。
- 页面上的「快捷启用/禁用」会把昵称、邮箱、手机号原样回传：`save` 是整体覆盖，
  漏传等于把这些字段清空。
- 新建用户与注册接口共用同一套用户名/密码长度校验，邮箱格式用 `email` 规则、
  手机号用 `phone-loose`（`1[3-9]` 开头 11 位，不写死号段，新号段不用改代码）。
- 编辑时用户名只读——它是登录身份，改名会让对方无法用原账号登录。
- 不能降级或禁用**当前登录的账号**，否则会把自己锁在后台外面。
- 不能让系统失去最后一名可用管理员（降级/禁用前会数一遍剩余管理员）。
- 昵称留空时自动回退成用户名，避免列表里出现空白行。
- 累计消费只统计已付款、已发货、已完成的订单，待付款和已取消不计入。
- 重置密码后原密码立即失效，但不会踢掉已登录的会话；要立刻断开请先禁用再启用该账号。

## 代码生成约定

以下文件由 CLI 生成维护，**不要手改**（改动会被覆盖）：

```bash
gf gen ctrl      # api/<模块>/<模块>.go 接口 + api/<模块>/v1 骨架
gf gen dao       # internal/dao、internal/model/do、internal/model/entity
gf gen service   # internal/service 接口 + internal/logic/logic.go 聚合注册
```

- 新增接口：先在 `api/<模块>/v1/` 写 Req/Res（含 `g.Meta` 路由标签），再 `gf gen ctrl` 生成骨架，最后填 controller 与 logic。
- `gf gen service` 按 `logic` 包的 `sXxx` 结构体方法生成接口，改接口签名后必须重跑。
- `internal/service`、`internal/logic` 之外的文件属于手写区。

## 关键实现约定

- **金额用整数分**：`utility/money` 做元/分换算，避免 float64 累加误差。
- **防超卖**：扣库存用 `Where("stock >= ?", n).Decrement(...)` 条件更新，影响行数为 0 即回滚事务。
- **订单快照**：`order_item` 存商品名/封面/单价，商品后续改价不影响历史订单。
- **越权防护**：所有写操作都带 `user_id` 条件，查不到统一返回「不存在」。
- **错误码**：业务错误用 `gerror.NewCode(gcode.Code...)`，由响应中间件统一格式化。
- **时区**：DSN 带 `loc=Local&parseTime=true`，否则 datetime 会差 8 小时。
- **状态流转靠条件更新**：发货/完成/取消都带 `WHERE status = 原状态`，重复点击第二次影响 0 行，
  天然幂等，不需要额外加锁。

## 说明

静态资源当前从 `resource/public` 目录读取（便于前端热改）。如需单文件部署，
可执行 `gf pack resource internal/packed/packed.go` 并在 `main.go` 空白导入 `internal/packed`，
把资源内嵌进二进制——代价是前端改动后需要重新打包。
