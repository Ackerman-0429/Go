-- ============================================================
-- 测试数据（开发调试用，勿在生产执行）
-- 执行：mysql -u root -p gfstudy < seed.sql
-- ============================================================

SET NAMES utf8mb4;

TRUNCATE TABLE `goods`;
TRUNCATE TABLE `category`;

INSERT INTO `category` (`id`, `parent_id`, `name`, `sort`, `status`) VALUES
(1, 0, '手机数码', 1, 1),
(2, 0, '电脑办公', 2, 1),
(3, 0, '服饰鞋包', 3, 1),
(4, 1, '智能手机', 1, 1),
(5, 2, '笔记本电脑', 1, 1);

INSERT INTO `goods`
(`category_id`, `name`, `subtitle`, `price`, `stock`, `cover`, `images`, `detail`, `sales`, `status`) VALUES
(4, 'GoFrame 旗舰手机 Pro',   '6.7英寸 2K 屏 / 12GB+512GB',   4999.00, 100,
 'https://picsum.photos/seed/phone1/400/400',
 '["https://picsum.photos/seed/phone1/800/800","https://picsum.photos/seed/phone1b/800/800"]',
 '旗舰级影像系统，支持 120W 快充。这是一段较长的商品详情文案，用于验证 detail 字段的存储与返回。', 128, 1),

(4, 'GoFrame 轻享手机',       '6.1英寸 / 8GB+256GB',         2999.00, 200,
 'https://picsum.photos/seed/phone2/400/400',
 '["https://picsum.photos/seed/phone2/800/800"]',
 '轻巧机身，长续航，适合日常使用。', 356, 1),

(5, 'GoFrame 创作本 16',      '16英寸 4K / 32GB+1TB',        8999.00, 50,
 'https://picsum.photos/seed/laptop1/400/400',
 '["https://picsum.photos/seed/laptop1/800/800","https://picsum.photos/seed/laptop1b/800/800"]',
 '面向设计师的高性能笔记本，独显加持。', 47, 1),

(5, 'GoFrame 轻薄本 14',      '14英寸 / 16GB+512GB',         5999.00, 80,
 'https://picsum.photos/seed/laptop2/400/400',
 '["https://picsum.photos/seed/laptop2/800/800"]',
 '1.2kg 轻薄便携，适合通勤办公。', 93, 1),

(3, '纯棉基础款 T 恤',        '100% 新疆长绒棉',              99.00, 500,
 'https://picsum.photos/seed/shirt1/400/400',
 '["https://picsum.photos/seed/shirt1/800/800"]',
 '四季通用基础款，多色可选。', 1280, 1),

(3, '复古牛仔外套',           '水洗做旧 / 男女同款',          299.00, 120,
 'https://picsum.photos/seed/jacket1/400/400',
 '["https://picsum.photos/seed/jacket1/800/800"]',
 '经典版型，耐穿百搭。', 210, 1),

(1, '65W 氮化镓充电器',       '双口输出 / 小巧便携',           129.00, 300,
 'https://picsum.photos/seed/charger/400/400',
 '["https://picsum.photos/seed/charger/800/800"]',
 '一个充电器搞定手机和笔记本。', 640, 1),

(1, '已经下架的商品',         '这个商品用于验证列表过滤',       19.90, 10,
 'https://picsum.photos/seed/offline/400/400',
 '[]',
 'status=0，用户端列表与详情都不应看到。', 0, 0);
