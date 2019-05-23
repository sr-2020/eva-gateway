-- MySQL dump 10.13  Distrib 5.7.25, for Linux (x86_64)
--
-- Host: localhost    Database: lumen
-- ------------------------------------------------------
-- Server version	5.7.25-0ubuntu0.18.04.2

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

CREATE DATABASE `eva-auth`;
GRANT ALL PRIVILEGES ON *.* TO 'app'@'%';

DROP TABLE IF EXISTS `eva-auth`.`users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `eva-auth`.`users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `admin` tinyint(1) NOT NULL DEFAULT '0',
  `beacon_id` int(11) DEFAULT NULL,
  `location_id` int(11) DEFAULT NULL,
  `status` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `email` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `password` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'test',
  `api_key` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `amount` int(11) NOT NULL DEFAULT '0',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `users_name_unique` (`name`),
  UNIQUE KEY `users_email_unique` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `eva-auth`.`users` WRITE;
/*!40000 ALTER TABLE `eva-auth`.`users` DISABLE KEYS */;
INSERT INTO `eva-auth`.`users` VALUES (1,1,NULL,NULL,'white','Мистер X','admin@evarun.ru','$2y$10$TRxK0F1twgCAbGygKTX1E.2wk0KNT3fPcfASlLtdFvU2AU8XsxZfG','TkRVem4yTERSQTNQRHFxcmo4SUozNWZp',223,'2019-03-30 04:57:38','2019-03-30 04:57:38');
/*!40000 ALTER TABLE `eva-auth`.`users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

CREATE DATABASE `eva-position`;

--
-- Table structure for table `eva-position`.`beacons`
--

DROP TABLE IF EXISTS `eva-position`.`beacons`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `eva-position`.`beacons` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `ssid` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `bssid` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `location_id` int(11) DEFAULT NULL,
  `label` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `beacons_bssid_unique` (`bssid`),
  KEY `beacons_location_id_index` (`location_id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `eva-position`.`beacons`
--

LOCK TABLES `eva-position`.`beacons` WRITE;
/*!40000 ALTER TABLE `eva-position`.`beacons` DISABLE KEYS */;
INSERT INTO `eva-position`.`beacons` VALUES (1,'E9:DC:0E:20:E3:DC','E9:DC:0E:20:E3:DC',1,NULL,NULL,NULL),(2,'D2:7E:91:02:AB:64','D2:7E:91:02:AB:64',1,NULL,NULL,NULL),(3,'F3:86:35:4C:6E:03','F3:86:35:4C:6E:03',1,NULL,NULL,NULL),(4,'C0:DA:B3:09:A9:FB','C0:DA:B3:09:A9:FB',2,NULL,NULL,NULL),(5,'F6:A3:B4:E1:D1:15','F6:A3:B4:E1:D1:15',2,NULL,NULL,NULL),(6,'F3:8F:DE:2F:66:C9','F3:8F:DE:2F:66:C9',3,NULL,NULL,NULL);
/*!40000 ALTER TABLE `eva-position`.`beacons` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `eva-position`.`locations`
--

DROP TABLE IF EXISTS `eva-position`.`locations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `eva-position`.`locations` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `label` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `eva-position`.`locations`
--

LOCK TABLES `eva-position`.`locations` WRITE;
/*!40000 ALTER TABLE `eva-position`.`locations` DISABLE KEYS */;
INSERT INTO `eva-position`.`locations` VALUES (1,'Танц-фойе Рим, 2 этаж',NULL,NULL),(2,'Концертный зал Москва',NULL,NULL),(3,'Левый коридор, 2 этаж',NULL,NULL);
/*!40000 ALTER TABLE `eva-position`.`locations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `eva-position`.`migrations`
--

DROP TABLE IF EXISTS `eva-position`.`migrations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `eva-position`.`migrations` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `migration` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `batch` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `eva-position`.`migrations`
--

LOCK TABLES `eva-position`.`migrations` WRITE;
/*!40000 ALTER TABLE `eva-position`.`migrations` DISABLE KEYS */;
INSERT INTO `eva-position`.`migrations` VALUES (7,'2018_10_14_000000_create_users_table',1),(8,'2019_01_03_053129_create_positions_table',1),(9,'2019_01_30_175738_create_beacons_table',1),(10,'2019_02_10_132148_create_paths_table',1),(11,'2019_03_03_135919_create_locations_table',1),(12,'2019_04_23_084050_create_positions_beacons_table',1);
/*!40000 ALTER TABLE `eva-position`.`migrations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `eva-position`.`paths`
--

DROP TABLE IF EXISTS `eva-position`.`paths`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `eva-position`.`paths` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `location_id` int(11) NOT NULL DEFAULT '0',
  `beacon_id` int(11) NOT NULL DEFAULT '0',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `eva-position`.`paths`
--

LOCK TABLES `eva-position`.`paths` WRITE;
/*!40000 ALTER TABLE `eva-position`.`paths` DISABLE KEYS */;
/*!40000 ALTER TABLE `eva-position`.`paths` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `eva-position`.`positions`
--

DROP TABLE IF EXISTS `eva-position`.`positions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `eva-position`.`positions` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) DEFAULT NULL,
  `beacons` json DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `eva-position`.`positions`
--

LOCK TABLES `eva-position`.`positions` WRITE;
/*!40000 ALTER TABLE `eva-position`.`positions` DISABLE KEYS */;
/*!40000 ALTER TABLE `eva-position`.`positions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `eva-position`.`positions_beacons`
--

DROP TABLE IF EXISTS `eva-position`.`positions_beacons`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `eva-position`.`positions_beacons` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `position_id` int(11) NOT NULL,
  `beacon_id` int(11) DEFAULT NULL,
  `bssid` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `level` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `positions_beacons_position_id_index` (`position_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `eva-position`.`positions_beacons`
--

LOCK TABLES `eva-position`.`positions_beacons` WRITE;
/*!40000 ALTER TABLE `eva-position`.`positions_beacons` DISABLE KEYS */;
/*!40000 ALTER TABLE `eva-position`.`positions_beacons` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `eva-position`.`users`
--

DROP TABLE IF EXISTS `eva-position`.`users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `eva-position`.`users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `beacon_id` int(11) DEFAULT NULL,
  `location_id` int(11) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `eva-position`.`users`
--

LOCK TABLES `eva-position`.`users` WRITE;
/*!40000 ALTER TABLE `eva-position`.`users` DISABLE KEYS */;
INSERT INTO `eva-position`.`users` VALUES (1,NULL,NULL,'2019-05-14 18:05:48','2019-05-14 18:05:48');
/*!40000 ALTER TABLE `eva-position`.`users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

CREATE DATABASE `billing`;

CREATE TABLE `billing`.`transactions` (
  `id` int(11),
  `created_at` datetime NOT NULL,
  `sin_from` int(11) NOT NULL,
  `sin_to` int(11) NOT NULL,
  `amount` int(11) NOT NULL,
  `comment` text,
  `recurrent_payment_id` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2019-03-30  7:58:37
