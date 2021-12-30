CREATE SCHEMA IF NOT EXISTS `blog_db` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_bin ;

CREATE TABLE IF NOT EXISTS blog_db.articles (
  `id` int NOT NULL AUTO_INCREMENT,
  `text` text,
  `link` text,
  `service` varchar(255) DEFAULT NULL,
  `article_id` varchar(255) DEFAULT NULL,
  `data_created` varchar(255) DEFAULT NULL,
  `created_at` varchar(255) NOT NULL,
  PRIMARY KEY (`id`, `created_at`)
);
