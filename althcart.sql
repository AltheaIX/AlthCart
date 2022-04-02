-- --------------------------------------------------------
-- Host:                         127.0.0.1
-- Server version:               5.7.33 - MySQL Community Server (GPL)
-- Server OS:                    Win64
-- HeidiSQL Version:             11.2.0.6213
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


-- Dumping database structure for althcart
CREATE DATABASE IF NOT EXISTS `althcart` /*!40100 DEFAULT CHARACTER SET latin1 COLLATE latin1_general_ci */;
USE `althcart`;

-- Dumping structure for table althcart.products
CREATE TABLE IF NOT EXISTS `products` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) COLLATE latin1_general_ci NOT NULL DEFAULT '0',
  `desc` text COLLATE latin1_general_ci NOT NULL,
  `image` text COLLATE latin1_general_ci NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=latin1 COLLATE=latin1_general_ci COMMENT='List products that will be indexed at website';

-- Dumping data for table althcart.products: ~5 rows (approximately)
/*!40000 ALTER TABLE `products` DISABLE KEYS */;
REPLACE INTO `products` (`id`, `name`, `desc`, `image`) VALUES
	(1, 'Apple Watch', 'Apple Smart Watch dengan chip terbaru dan layar OLED Retina buatan Apple yang memanjakan mata.', 'rachit-tank-2cFZ_FB08UM-unsplash.jpg'),
	(2, 'Fujica Mirrorless', 'Kamera Nikon generasi terbaru yang memiliki Image Processing sangat baik, yang membuat kamera ini sangat laku dipasaran.', 'marcus-urbenz-4J77m6D9430-unsplash.jpg'),
	(3, 'Nike Airmax 200', 'Sepatu dengan mobilitas yang baik dan outsole yang sudah dilengkapi XDR untuk bermain di outdoor.', 'luis-villasmil-SmCIRo1QCpo-unsplash.jpg'),
	(4, 'Drone', 'Drone dengan Panorama View Camera yang memudahkan anda untuk memotret pemandangan dari atas dengan mudah.', 'george-kroeker-96bacP6Vba8-unsplash.jpg');
/*!40000 ALTER TABLE `products` ENABLE KEYS */;

-- Dumping structure for table althcart.users
CREATE TABLE IF NOT EXISTS `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(50) COLLATE latin1_general_ci NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1 COLLATE=latin1_general_ci;

-- Dumping data for table althcart.users: ~1 rows (approximately)
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
REPLACE INTO `users` (`id`, `username`) VALUES
	(1, 'Malik');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;

-- Dumping structure for table althcart.users_cart
CREATE TABLE IF NOT EXISTS `users_cart` (
  `id` int(15) NOT NULL AUTO_INCREMENT,
  `product_id` int(15) NOT NULL DEFAULT '0',
  `username` varchar(50) COLLATE latin1_general_ci NOT NULL,
  `quantity` int(15) NOT NULL DEFAULT '1',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=latin1 COLLATE=latin1_general_ci;

-- Dumping data for table althcart.users_cart: ~5 rows (approximately)
/*!40000 ALTER TABLE `users_cart` DISABLE KEYS */;
REPLACE INTO `users_cart` (`id`, `product_id`, `username`, `quantity`) VALUES
	(21, 4, 'Malik', 2);
/*!40000 ALTER TABLE `users_cart` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
