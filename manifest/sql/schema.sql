-- ============================================================
-- 网购系统 数据库设计（MVP v1）
-- 数据库：gfstudy      引擎：InnoDB      字符集：utf8mb4
--
-- 表清单：
--   user        用户
--   category    商品分类
--   goods       商品
--   cart        购物车
--   orders      订单（注意：order 是 MySQL 保留字，故用复数）
--   order_item  订单明细
--
-- 执行：mysql -u root -p gfstudy < schema.sql
-- ============================================================

SET NAMES utf8mb4;

-- ------------------------------------------------------------
-- 用户
-- ------------------------------------------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id`         bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `username`   varchar(50)     NOT NULL                COMMENT '用户名，登录用',
  `password`   varchar(255)    NOT NULL                COMMENT '密码哈希(bcrypt)，禁止存明文',
  `nickname`   varchar(50)     NOT NULL DEFAULT ''     COMMENT '昵称',
  `email`      varchar(100)             DEFAULT NULL   COMMENT '邮箱，可为空但不可重复',
  `phone`      varchar(20)              DEFAULT NULL   COMMENT '手机号，可为空但不可重复',
  `status`     tinyint unsigned NOT NULL DEFAULT 1     COMMENT '状态：1正常 0禁用',
  `role`       tinyint unsigned NOT NULL DEFAULT 0     COMMENT '角色：0普通用户 1管理员',
  `created_at` datetime        NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime        NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_username` (`username`),
  UNIQUE KEY `uk_email` (`email`),
  UNIQUE KEY `uk_phone` (`phone`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户';

-- ------------------------------------------------------------
-- 商品分类（支持二级：parent_id=0 为顶级）
-- ------------------------------------------------------------
DROP TABLE IF EXISTS `category`;
CREATE TABLE `category` (
  `id`         int unsigned    NOT NULL AUTO_INCREMENT COMMENT '分类ID',
  `parent_id`  int unsigned    NOT NULL DEFAULT 0      COMMENT '父分类ID，0表示顶级',
  `name`       varchar(50)     NOT NULL                COMMENT '分类名称',
  `sort`       int             NOT NULL DEFAULT 0      COMMENT '排序值，越小越靠前',
  `status`     tinyint unsigned NOT NULL DEFAULT 1     COMMENT '状态：1启用 0停用',
  `created_at` datetime        NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime        NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_parent` (`parent_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='商品分类';

-- ------------------------------------------------------------
-- 商品
-- ------------------------------------------------------------
DROP TABLE IF EXISTS `goods`;
CREATE TABLE `goods` (
  `id`          bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '商品ID',
  `category_id` int unsigned    NOT NULL                COMMENT '所属分类ID',
  `name`        varchar(200)    NOT NULL                COMMENT '商品名称',
  `subtitle`    varchar(255)    NOT NULL DEFAULT ''     COMMENT '副标题/卖点',
  `price`       decimal(10,2)   NOT NULL DEFAULT 0.00   COMMENT '售价，必须用DECIMAL',
  `stock`       int unsigned    NOT NULL DEFAULT 0      COMMENT '库存',
  `cover`       varchar(255)    NOT NULL DEFAULT ''     COMMENT '封面图URL',
  `images`      text                                    COMMENT '多图URL，JSON数组',
  `detail`      text                                    COMMENT '商品详情',
  `sales`       int unsigned    NOT NULL DEFAULT 0      COMMENT '销量',
  `status`      tinyint unsigned NOT NULL DEFAULT 1     COMMENT '状态：1上架 0下架',
  `created_at`  datetime        NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at`  datetime        NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_category` (`category_id`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='商品';

-- ------------------------------------------------------------
-- 购物车（同一用户同一商品只保留一条，加购则累加数量）
-- ------------------------------------------------------------
DROP TABLE IF EXISTS `cart`;
CREATE TABLE `cart` (
  `id`         bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '购物车项ID',
  `user_id`    bigint unsigned NOT NULL                COMMENT '用户ID',
  `goods_id`   bigint unsigned NOT NULL                COMMENT '商品ID',
  `quantity`   int unsigned    NOT NULL DEFAULT 1      COMMENT '数量',
  `created_at` datetime        NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime        NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_user_goods` (`user_id`, `goods_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='购物车';

-- ------------------------------------------------------------
-- 订单（order 为保留字，故用 orders）
-- ------------------------------------------------------------
DROP TABLE IF EXISTS `orders`;
CREATE TABLE `orders` (
  `id`               bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '订单ID',
  `order_no`         varchar(32)     NOT NULL                COMMENT '订单号，对外展示',
  `user_id`          bigint unsigned NOT NULL                COMMENT '下单用户ID',
  `total_amount`     decimal(10,2)   NOT NULL DEFAULT 0.00   COMMENT '订单总金额',
  `status`           tinyint unsigned NOT NULL DEFAULT 0     COMMENT '状态：0待付款 1已付款 2已发货 3已完成 4已取消',
  `receiver_name`    varchar(50)     NOT NULL DEFAULT ''     COMMENT '收货人',
  `receiver_phone`   varchar(20)     NOT NULL DEFAULT ''     COMMENT '收货电话',
  `receiver_address` varchar(255)    NOT NULL DEFAULT ''     COMMENT '收货地址',
  `remark`           varchar(255)    NOT NULL DEFAULT ''     COMMENT '订单备注',
  `paid_at`          datetime                 DEFAULT NULL   COMMENT '付款时间',
  `created_at`       datetime        NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at`       datetime        NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_order_no` (`order_no`),
  KEY `idx_user` (`user_id`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='订单';

-- ------------------------------------------------------------
-- 订单明细（必须冗余商品名与价格快照）
-- ------------------------------------------------------------
DROP TABLE IF EXISTS `order_item`;
CREATE TABLE `order_item` (
  `id`          bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '明细ID',
  `order_id`    bigint unsigned NOT NULL                COMMENT '订单ID',
  `goods_id`    bigint unsigned NOT NULL                COMMENT '商品ID',
  `goods_name`  varchar(200)    NOT NULL                COMMENT '下单时商品名快照',
  `goods_cover` varchar(255)    NOT NULL DEFAULT ''     COMMENT '下单时封面图快照',
  `goods_price` decimal(10,2)   NOT NULL                COMMENT '下单时单价快照',
  `quantity`    int unsigned    NOT NULL                COMMENT '购买数量',
  `subtotal`    decimal(10,2)   NOT NULL                COMMENT '小计 = 单价 × 数量',
  `created_at`  datetime        NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_order` (`order_id`),
  KEY `idx_goods` (`goods_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='订单明细';
