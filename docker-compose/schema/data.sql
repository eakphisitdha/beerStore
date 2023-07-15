CREATE DATABASE IF NOT EXISTS `data` CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE `data`;

CREATE TABLE IF NOT EXISTS  `beer`
(
    `id` int PRIMARY KEY AUTO_INCREMENT,
    `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
    `type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,	
    `detail` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
    `url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
    INDEX (name)

) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

INSERT INTO `beer` (`name`, `type`, `detail`, `url`)
VALUES
    ('Heineken', 'Pale Lager', 'เบียร์แบรนด์ Heineken', 'https://example.com/heineken.jpg'),
    ('Guinness', 'Dry Stout', 'เบียร์แบรนด์ Guinness', 'https://example.com/guinness.jpg'),
    ('Budweiser', 'Pale Lager', 'เบียร์แบรนด์ Budweiser', 'https://example.com/budweiser.jpg'),
    ('Stella Artois', 'Pale Lager', 'เบียร์แบรนด์ Stella Artois', 'https://example.com/stella-artois.jpg'),
    ('Corona', 'Light Lager', 'เบียร์แบรนด์ Corona', 'https://example.com/corona.jpg'),
    ('IPA', 'India Pale Ale', 'เบียร์แบรนด์ IPA', 'https://example.com/ipa.jpg'),
    ('Singha', 'Lager', 'เบียร์แบรนด์ Singha', 'https://example.com/singha.jpg'),
    ('Tiger', 'Lager', 'เบียร์แบรนด์ Tiger', 'https://example.com/tiger.jpg'),
    ('Pilsner Urquell', 'Pilsner', 'เบียร์แบรนด์ Pilsner Urquell', 'https://example.com/pilsner-urquell.jpg'),
    ('Chang', 'Lager', 'เบียร์แบรนด์ Chang', 'https://example.com/chang.jpg'),
    ('Hoegaarden', 'Wheat Beer', 'เบียร์แบรนด์ Hoegaarden', 'https://example.com/hoegaarden.jpg'),
    ('Sapporo', 'Lager', 'เบียร์แบรนด์ Sapporo', 'https://example.com/sapporo.jpg'),
    ('Leffe', 'Abbey Beer', 'เบียร์แบรนด์ Leffe', 'https://example.com/leffe.jpg'),
    ('Carlsberg', 'Lager', 'เบียร์แบรนด์ Carlsberg', 'https://example.com/carlsberg.jpg'),
    ('Blue Moon', 'Belgian White', 'เบียร์แบรนด์ Blue Moon', 'https://example.com/blue-moon.jpg'),
    ('Asahi', 'Lager', 'เบียร์แบรนด์ Asahi', 'https://example.com/asahi.jpg'),
    ('Modelo Especial', 'Pale Lager', 'เบียร์แบรนด์ Modelo Especial', 'https://example.com/modelo-especial.jpg'),
    ('Harbin', 'Lager', 'เบียร์แบรนด์ Harbin', 'https://example.com/harbin.jpg'),
    ('Kronenbourg 1664', 'Lager', 'เบียร์แบรนด์ Kronenbourg 1664', 'https://example.com/kronenbourg-1664.jpg'),
    ('Redd\'s Apple Ale', 'Fruit Beer', 'เบียร์แบรนด์ Redd\'s Apple Ale', 'https://example.com/redds-apple-ale.jpg'),
    ('Hoegaarden Rosée', 'Fruit Beer', 'เบียร์แบรนด์ Hoegaarden Rosée', 'https://example.com/hoegaarden-rosee.jpg'),
    ('Grimbergen', 'Abbey Beer', 'เบียร์แบรนด์ Grimbergen', 'https://example.com/grimbergen.jpg'),
    ('Paulaner', 'Hefeweizen', 'เบียร์แบรนด์ Paulaner', 'https://example.com/paulaner.jpg'),
    ('König Ludwig', 'Weissbier', 'เบียร์แบรนด์ König Ludwig', 'https://example.com/konig-ludwig.jpg'),
    ('Duvel', 'Strong Pale Ale', 'เบียร์แบรนด์ Duvel', 'https://example.com/duvel.jpg'),
    ('Sierra Nevada Pale Ale', 'Pale Ale', 'เบียร์แบรนด์ Sierra Nevada Pale Ale', 'https://example.com/sierra-nevada-pale-ale.jpg'),
    ('Chimay', 'Trappist Beer', 'เบียร์แบรนด์ Chimay', 'https://example.com/chimay.jpg'),
    ('Kasteel Tripel', 'Tripel', 'เบียร์แบรนด์ Kasteel Tripel', 'https://example.com/kasteel-tripel.jpg'),
    ('Weihenstephaner Hefeweissbier', 'Hefeweizen', 'เบียร์แบรนด์ Weihenstephaner Hefeweissbier', 'https://example.com/weihenstephaner-hefeweissbier.jpg'),
    ('La Chouffe', 'Blonde Ale', 'เบียร์แบรนด์ La Chouffe', 'https://example.com/la-chouffe.jpg');
