# GoShop · GoFrame 网购系统（MVP）

基于 GoFrame v2 的电商后端 MVP，已跑通「注册 → 浏览商品 → 加购物车 → 下单 → 支付」闭环，
并附带一个原生 JS 的展示页。

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
- 接口文档：http://127.0.0.1:8000/swagger/
- OpenAPI：http://127.0.0.1:8000/api.json

演示账号：`alice` / `bob`，密码 `123456`。

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
resource/public/        静态资源（商城展示页）
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

统一响应体：`{"code":0,"message":"OK","data":{...}}`，由 `ghttp.MiddlewareHandlerResponse` 输出。

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

## 说明

静态资源当前从 `resource/public` 目录读取（便于前端热改）。如需单文件部署，
可执行 `gf pack resource internal/packed/packed.go` 并在 `main.go` 空白导入 `internal/packed`，
把资源内嵌进二进制——代价是前端改动后需要重新打包。
