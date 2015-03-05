# ************************************************************
# Sequel Pro SQL dump
# Version 4135
#
# http://www.sequelpro.com/
# http://code.google.com/p/sequel-pro/
#
# Host: 127.0.0.1 (MySQL 5.5.38)
# Database: shades
# Generation Time: 2015-03-05 10:34:59 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table shades
# ------------------------------------------------------------

DROP TABLE IF EXISTS `shades`;

CREATE TABLE `shades` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Grey` tinyint(1) NOT NULL,
  `Hex` tinytext NOT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `shades` WRITE;
/*!40000 ALTER TABLE `shades` DISABLE KEYS */;

INSERT INTO `shades` (`Id`, `Grey`, `Hex`)
VALUES
	(1,1,'#4a4a49'),
	(2,1,'#575756'),
	(3,1,'#646363'),
	(4,1,'#727271'),
	(5,1,'#838382'),
	(6,1,'#909090'),
	(7,1,'#9d9d9c'),
	(8,1,'#a8a8a7'),
	(9,1,'#b4b4b4'),
	(10,1,'#c2c2c2'),
	(11,1,'#4d4d4c'),
	(12,1,'#5a5a59'),
	(13,1,'#666665'),
	(14,1,'#757474'),
	(15,1,'#858584'),
	(16,1,'#939393'),
	(17,1,'#9f9f9f'),
	(18,1,'#aaaaa9'),
	(19,1,'#b7b6b6'),
	(20,1,'#c4c4c4'),
	(21,1,'#504f4e'),
	(22,1,'#5c5c5b'),
	(23,1,'#686867'),
	(24,1,'#777776'),
	(25,1,'#878787'),
	(26,1,'#949494'),
	(27,1,'#a1a1a1'),
	(28,1,'#acacac'),
	(29,1,'#b9b8b8'),
	(30,1,'#c6c6c6'),
	(31,1,'#525251'),
	(32,1,'#5f5e5e'),
	(33,1,'#6b6a6a'),
	(34,1,'#797979'),
	(35,1,'#8a8989'),
	(36,1,'#979696'),
	(37,1,'#a4a3a3'),
	(38,1,'#aeaeae'),
	(39,1,'#bbbaba'),
	(40,1,'#c8c8c8'),
	(41,1,'#555454'),
	(42,1,'#616160'),
	(43,1,'#6d6d6c'),
	(44,1,'#7e7e7d'),
	(45,1,'#8c8c8b'),
	(46,1,'#999998'),
	(47,1,'#a6a5a5'),
	(48,1,'#b0b0b0'),
	(49,1,'#bebebe'),
	(50,1,'#cacaca');

/*!40000 ALTER TABLE `shades` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
