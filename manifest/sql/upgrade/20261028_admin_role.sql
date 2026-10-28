-- ============================================================
-- 管理员账号调整（2026-10-28）
--   1) 确保 admin 账号存在，并设为管理员（已存在则保留原密码）
--   2) 收回 alice 的管理员权限，只留 admin 一个管理员
-- 执行：mysql -h 127.0.0.1 -u root -p gfstudy < 20261028_admin_role.sql
-- 说明：脚本里的哈希对应演示密码，仅供本地开发环境使用。
-- ============================================================

SET NAMES utf8mb4;

-- 已存在则不动密码，只纠正角色与状态，避免覆盖用户自己设置的密码
INSERT INTO `user` (`username`, `nickname`, `password`, `role`, `status`)
VALUES ('admin', '管理员', '$2a$10$wrAgeZpLBqGGXq9Na9BsLe9MtoCimDww.hYPZaJ1Tlg6VtLSO6y6q', 1, 1)
ON DUPLICATE KEY UPDATE `role` = 1, `status` = 1;

-- 收回 alice 的管理员权限（普通用户仍可正常登录下单）
UPDATE `user` SET `role` = 0 WHERE `username` = 'alice';

SELECT `id`, `username`, `role` FROM `user` ORDER BY `id`;
