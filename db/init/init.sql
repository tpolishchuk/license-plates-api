DROP DATABASE IF EXISTS `license-plates`;
CREATE DATABASE `license-plates`;
USE `license-plates`;

CREATE TABLE `plates` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `abbreviation` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `country` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `city_or_region` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `state` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_plates_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
