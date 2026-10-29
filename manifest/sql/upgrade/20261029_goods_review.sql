-- ============================================================
-- 升级脚本：新增 goods_review 商品评价表
-- 适用：已按旧版 schema.sql 建库的环境
-- 执行：mysql -h 127.0.0.1 -u root -p gfstudy < 20261029_goods_review.sql
-- ============================================================

SET NAMES utf8mb4;

CREATE TABLE IF NOT EXISTS `goods_review` (
  `id`            bigint unsigned  NOT NULL AUTO_INCREMENT COMMENT '评价ID',
  `goods_id`      bigint unsigned  NOT NULL                COMMENT '商品ID',
  `order_id`      bigint unsigned  NOT NULL                COMMENT '订单ID',
  `order_item_id` bigint unsigned  NOT NULL                COMMENT '订单明细ID，一条明细只能评一次',
  `user_id`       bigint unsigned  NOT NULL                COMMENT '评价用户ID',
  `rating`        tinyint unsigned NOT NULL DEFAULT 5      COMMENT '评分 1-5 星',
  `content`       varchar(500)     NOT NULL DEFAULT ''     COMMENT '评价内容',
  `anonymous`     tinyint unsigned NOT NULL DEFAULT 0      COMMENT '是否匿名：1 是 0 否',
  `status`        tinyint unsigned NOT NULL DEFAULT 1      COMMENT '状态：1 显示 0 隐藏',
  `created_at`    datetime         NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  -- 同一条订单明细只能评价一次，重复提交靠唯一键兜底
  UNIQUE KEY `uk_order_item` (`order_item_id`),
  KEY `idx_goods` (`goods_id`, `status`),
  KEY `idx_user` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='商品评价';
