-- ============================================================
-- 升级脚本：user 表增加 role 角色字段（后台管理）
-- 适用：已按旧版 schema.sql 建库的环境
-- 执行：mysql -h 127.0.0.1 -u root -p gfstudy < 20261028_user_role.sql
-- ============================================================

SET NAMES utf8mb4;

ALTER TABLE `user`
  ADD COLUMN `role` tinyint unsigned NOT NULL DEFAULT 0
  COMMENT '角色：0普通用户 1管理员' AFTER `status`;

-- 演示环境：把已有账号 alice 提为管理员
UPDATE `user` SET `role` = 1 WHERE `username` = 'alice';
