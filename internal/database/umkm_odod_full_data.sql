-- --------------------------------------------------------
-- Host:                         127.0.0.1
-- Server version:               8.4.6 - MySQL Community Server - GPL
-- Server OS:                    Win64
-- HeidiSQL Version:             12.11.0.7065
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


-- Dumping database structure for umkm_odod
CREATE DATABASE IF NOT EXISTS `umkm_odod` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `umkm_odod`;

-- Dumping structure for table umkm_odod.catalog_categories
CREATE TABLE IF NOT EXISTS `catalog_categories` (
  `id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `tenant_id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_categories_name_per_tenant` (`tenant_id`,`name`),
  CONSTRAINT `fk_categories_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table umkm_odod.catalog_categories: ~2 rows (approximately)
INSERT INTO `catalog_categories` (`id`, `tenant_id`, `name`, `created_at`, `updated_at`, `deleted_at`) VALUES
	('44444444-4444-4444-4444-444444444441', '11111111-1111-1111-1111-111111111111', 'Keripik Singkong', '2026-05-18 09:21:24', '2026-05-18 09:21:24', NULL),
	('44444444-4444-4444-4444-444444444442', '11111111-1111-1111-1111-111111111111', 'Keripik Pisang', '2026-05-18 09:21:24', '2026-05-18 09:21:24', NULL);

-- Dumping structure for table umkm_odod.catalog_items
CREATE TABLE IF NOT EXISTS `catalog_items` (
  `id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `tenant_id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `category_id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `name` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `is_active` tinyint(1) NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_items_tenant` (`tenant_id`),
  KEY `fk_items_category` (`category_id`),
  CONSTRAINT `fk_items_category` FOREIGN KEY (`category_id`) REFERENCES `catalog_categories` (`id`),
  CONSTRAINT `fk_items_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table umkm_odod.catalog_items: ~2 rows (approximately)
INSERT INTO `catalog_items` (`id`, `tenant_id`, `category_id`, `name`, `description`, `is_active`, `created_at`, `updated_at`, `deleted_at`) VALUES
	('55555555-5555-5555-5555-555555555551', '11111111-1111-1111-1111-111111111111', '44444444-4444-4444-4444-444444444441', 'Keripik Singkong', 'Keripik singkong premium', 1, '2026-05-18 09:21:44', '2026-05-18 09:21:44', NULL),
	('55555555-5555-5555-5555-555555555552', '11111111-1111-1111-1111-111111111111', '44444444-4444-4444-4444-444444444442', 'Keripik Pisang', 'Keripik pisang renyah', 1, '2026-05-18 09:21:44', '2026-05-18 09:21:44', NULL);

-- Dumping structure for table umkm_odod.item_attributes
CREATE TABLE IF NOT EXISTS `item_attributes` (
  `id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `tenant_id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_item_attributes_name_per_tenant` (`tenant_id`,`name`),
  CONSTRAINT `fk_item_attributes_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table umkm_odod.item_attributes: ~2 rows (approximately)
INSERT INTO `item_attributes` (`id`, `tenant_id`, `name`, `created_at`, `updated_at`, `deleted_at`) VALUES
	('66666666-6666-6666-6666-666666666661', '11111111-1111-1111-1111-111111111111', 'Rasa', '2026-05-18 09:21:58', '2026-05-18 09:21:58', NULL),
	('66666666-6666-6666-6666-666666666662', '11111111-1111-1111-1111-111111111111', 'Berat', '2026-05-18 09:21:58', '2026-05-18 09:21:58', NULL);

-- Dumping structure for table umkm_odod.item_attribute_values
CREATE TABLE IF NOT EXISTS `item_attribute_values` (
  `id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `tenant_id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `attribute_id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `value` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_attribute_values_tenant` (`tenant_id`),
  KEY `fk_attribute_values_attribute` (`attribute_id`),
  CONSTRAINT `fk_attribute_values_attribute` FOREIGN KEY (`attribute_id`) REFERENCES `item_attributes` (`id`),
  CONSTRAINT `fk_attribute_values_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table umkm_odod.item_attribute_values: ~3 rows (approximately)
INSERT INTO `item_attribute_values` (`id`, `tenant_id`, `attribute_id`, `value`, `created_at`, `updated_at`, `deleted_at`) VALUES
	('77777777-7777-7777-7777-777777777771', '11111111-1111-1111-1111-111111111111', '66666666-6666-6666-6666-666666666661', 'Pedas Daun Jeruk', '2026-05-18 09:22:08', '2026-05-18 09:22:08', NULL),
	('77777777-7777-7777-7777-777777777772', '11111111-1111-1111-1111-111111111111', '66666666-6666-6666-6666-666666666661', 'Asin Bawang', '2026-05-18 09:22:08', '2026-05-18 09:22:08', NULL),
	('77777777-7777-7777-7777-777777777773', '11111111-1111-1111-1111-111111111111', '66666666-6666-6666-6666-666666666662', '250 Gram', '2026-05-18 09:22:08', '2026-05-18 09:22:08', NULL),
	('77777777-7777-7777-7777-777777777774', '11111111-1111-1111-1111-111111111111', '66666666-6666-6666-6666-666666666662', '500 Gram', '2026-05-18 09:22:08', '2026-05-18 09:22:08', NULL);

-- Dumping structure for table umkm_odod.item_variants
CREATE TABLE IF NOT EXISTS `item_variants` (
  `id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `tenant_id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `item_id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `sku` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `barcode` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `variant_name` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `cost_price` decimal(18,2) NOT NULL DEFAULT '0.00',
  `is_active` tinyint(1) NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_variants_sku_per_tenant` (`tenant_id`,`sku`),
  KEY `fk_variants_item` (`item_id`),
  CONSTRAINT `fk_variants_item` FOREIGN KEY (`item_id`) REFERENCES `catalog_items` (`id`),
  CONSTRAINT `fk_variants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table umkm_odod.item_variants: ~2 rows (approximately)
INSERT INTO `item_variants` (`id`, `tenant_id`, `item_id`, `sku`, `barcode`, `variant_name`, `cost_price`, `is_active`, `created_at`, `updated_at`, `deleted_at`) VALUES
	('88888888-8888-8888-8888-888888888881', '11111111-1111-1111-1111-111111111111', '55555555-5555-5555-5555-555555555551', 'KS-PDJ-250', '899100000001', 'Pedas Daun Jeruk 250gr', 12000.00, 1, '2026-05-18 09:22:19', '2026-05-18 09:22:19', NULL),
	('88888888-8888-8888-8888-888888888882', '11111111-1111-1111-1111-111111111111', '55555555-5555-5555-5555-555555555551', 'KS-AB-500', '899100000002', 'Asin Bawang 500gr', 18000.00, 1, '2026-05-18 09:22:19', '2026-05-18 09:22:19', NULL);

-- Dumping structure for table umkm_odod.item_variant_attribute_values
CREATE TABLE IF NOT EXISTS `item_variant_attribute_values` (
  `variant_id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `attribute_value_id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `tenant_id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`variant_id`,`attribute_value_id`),
  KEY `fk_variant_attribute_value` (`attribute_value_id`),
  KEY `FK_item_variant_attribute_values_tenants` (`tenant_id`),
  CONSTRAINT `FK_item_variant_attribute_values_tenants` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`),
  CONSTRAINT `fk_variant_attribute_value` FOREIGN KEY (`attribute_value_id`) REFERENCES `item_attribute_values` (`id`),
  CONSTRAINT `fk_variant_attribute_variant` FOREIGN KEY (`variant_id`) REFERENCES `item_variants` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table umkm_odod.item_variant_attribute_values: ~4 rows (approximately)
INSERT INTO `item_variant_attribute_values` (`variant_id`, `attribute_value_id`, `tenant_id`) VALUES
	('88888888-8888-8888-8888-888888888881', '77777777-7777-7777-7777-777777777771', '11111111-1111-1111-1111-111111111111'),
	('88888888-8888-8888-8888-888888888881', '77777777-7777-7777-7777-777777777773', '11111111-1111-1111-1111-111111111111'),
	('88888888-8888-8888-8888-888888888882', '77777777-7777-7777-7777-777777777772', '11111111-1111-1111-1111-111111111111'),
	('88888888-8888-8888-8888-888888888882', '77777777-7777-7777-7777-777777777774', '11111111-1111-1111-1111-111111111111');

-- Dumping structure for table umkm_odod.price_histories
CREATE TABLE IF NOT EXISTS `price_histories` (
  `id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `tenant_id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `item_variant_id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `price_type` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `price` decimal(18,2) NOT NULL,
  `effective_date` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `created_by` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `fk_prices_tenant` (`tenant_id`),
  KEY `fk_prices_variant` (`item_variant_id`),
  KEY `fk_prices_created_by` (`created_by`),
  CONSTRAINT `fk_prices_created_by` FOREIGN KEY (`created_by`) REFERENCES `users` (`id`),
  CONSTRAINT `fk_prices_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`),
  CONSTRAINT `fk_prices_variant` FOREIGN KEY (`item_variant_id`) REFERENCES `item_variants` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table umkm_odod.price_histories: ~2 rows (approximately)
INSERT INTO `price_histories` (`id`, `tenant_id`, `item_variant_id`, `price_type`, `price`, `effective_date`, `created_by`, `created_at`) VALUES
	('99999999-9999-9999-9999-999999999991', '11111111-1111-1111-1111-111111111111', '88888888-8888-8888-8888-888888888881', 'SELLING', 18000.00, '2026-05-20 06:36:19', '2416854f-55e6-423c-a2ec-8154c9431cd6', '2026-05-18 09:23:03'),
	('99999999-9999-9999-9999-999999999992', '11111111-1111-1111-1111-111111111111', '88888888-8888-8888-8888-888888888882', 'SELLING', 25000.00, '2026-05-20 06:36:20', '2416854f-55e6-423c-a2ec-8154c9431cd6', '2026-05-18 09:23:03');

-- Dumping structure for table umkm_odod.roles
CREATE TABLE IF NOT EXISTS `roles` (
  `id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `tenant_id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_roles_name_per_tenant` (`tenant_id`,`name`),
  CONSTRAINT `fk_roles_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table umkm_odod.roles: ~4 rows (approximately)
INSERT INTO `roles` (`id`, `tenant_id`, `name`, `created_at`, `updated_at`, `deleted_at`) VALUES
	('22222222-2222-2222-2222-222222222221', '11111111-1111-1111-1111-111111111112', 'OWNER', '2026-05-18 09:20:36', '2026-05-18 16:35:19', NULL),
	('22222222-2222-2222-2222-222222222222', '11111111-1111-1111-1111-111111111111', 'CASHIER', '2026-05-18 09:20:36', '2026-05-18 09:20:36', NULL),
	('ab88f666-b2d6-482c-bdcd-be29915e5395', 'f27e441f-5385-4b8d-b2e2-88b8615a4634', 'SUPER ADMIN', '2026-05-19 02:02:01', '2026-05-19 02:02:01', NULL),
	('dc064501-5798-4c91-8308-0198561ceae3', 'f27e441f-5385-4b8d-b2e2-88b8615a4634', 'ADMIN', '2026-05-19 01:56:21', '2026-05-19 02:17:01', NULL);

-- Dumping structure for table umkm_odod.sales
CREATE TABLE IF NOT EXISTS `sales` (
  `id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `tenant_id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `invoice_number` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `customer_name` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `cashier_id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `subtotal` decimal(18,2) NOT NULL DEFAULT '0.00',
  `discount_amount` decimal(18,2) NOT NULL DEFAULT '0.00',
  `tax_amount` decimal(18,2) NOT NULL DEFAULT '0.00',
  `grand_total` decimal(18,2) NOT NULL DEFAULT '0.00',
  `payment_method` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `payment_status` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `notes` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_sales_invoice_per_tenant` (`tenant_id`,`invoice_number`),
  KEY `fk_sales_cashier` (`cashier_id`),
  CONSTRAINT `fk_sales_cashier` FOREIGN KEY (`cashier_id`) REFERENCES `users` (`id`),
  CONSTRAINT `fk_sales_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table umkm_odod.sales: ~1 rows (approximately)
INSERT INTO `sales` (`id`, `tenant_id`, `invoice_number`, `customer_name`, `cashier_id`, `subtotal`, `discount_amount`, `tax_amount`, `grand_total`, `payment_method`, `payment_status`, `notes`, `created_at`, `updated_at`) VALUES
	('aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa', '11111111-1111-1111-1111-111111111111', 'INV-20260518-0001', 'Andi', '2416854f-55e6-423c-a2ec-8154c9431cd6', 43000.00, 3000.00, 0.00, 40000.00, 'QRIS', 'PAID', 'Pembelian sore hari', '2026-05-18 09:23:17', '2026-05-20 06:37:42');

-- Dumping structure for table umkm_odod.sale_items
CREATE TABLE IF NOT EXISTS `sale_items` (
  `id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `tenant_id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `sale_id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `item_variant_id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `item_name_snapshot` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `variant_name_snapshot` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `sku_snapshot` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `qty` decimal(18,2) NOT NULL,
  `unit_price` decimal(18,2) NOT NULL,
  `discount_amount` decimal(18,2) NOT NULL DEFAULT '0.00',
  `subtotal` decimal(18,2) NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `fk_sale_items_tenant` (`tenant_id`),
  KEY `fk_sale_items_sale` (`sale_id`),
  KEY `fk_sale_items_variant` (`item_variant_id`),
  CONSTRAINT `fk_sale_items_sale` FOREIGN KEY (`sale_id`) REFERENCES `sales` (`id`),
  CONSTRAINT `fk_sale_items_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`),
  CONSTRAINT `fk_sale_items_variant` FOREIGN KEY (`item_variant_id`) REFERENCES `item_variants` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table umkm_odod.sale_items: ~2 rows (approximately)
INSERT INTO `sale_items` (`id`, `tenant_id`, `sale_id`, `item_variant_id`, `item_name_snapshot`, `variant_name_snapshot`, `sku_snapshot`, `qty`, `unit_price`, `discount_amount`, `subtotal`, `created_at`) VALUES
	('bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbb1', '11111111-1111-1111-1111-111111111111', 'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa', '88888888-8888-8888-8888-888888888881', 'Keripik Singkong', 'Pedas Daun Jeruk 250gr', 'KS-PDJ-250', 1.00, 18000.00, 0.00, 18000.00, '2026-05-18 09:24:03'),
	('bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbb2', '11111111-1111-1111-1111-111111111111', 'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa', '88888888-8888-8888-8888-888888888882', 'Keripik Singkong', 'Asin Bawang 500gr', 'KS-AB-500', 1.00, 25000.00, 3000.00, 22000.00, '2026-05-18 09:24:03');

-- Dumping structure for table umkm_odod.stock_movements
CREATE TABLE IF NOT EXISTS `stock_movements` (
  `id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `tenant_id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `item_variant_id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `movement_type` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `qty` decimal(18,2) NOT NULL,
  `reference_type` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `reference_id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `notes` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `created_by` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `fk_stock_tenant` (`tenant_id`),
  KEY `fk_stock_variant` (`item_variant_id`),
  KEY `fk_stock_created_by` (`created_by`),
  CONSTRAINT `fk_stock_created_by` FOREIGN KEY (`created_by`) REFERENCES `users` (`id`),
  CONSTRAINT `fk_stock_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`),
  CONSTRAINT `fk_stock_variant` FOREIGN KEY (`item_variant_id`) REFERENCES `item_variants` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table umkm_odod.stock_movements: ~2 rows (approximately)
INSERT INTO `stock_movements` (`id`, `tenant_id`, `item_variant_id`, `movement_type`, `qty`, `reference_type`, `reference_id`, `notes`, `created_by`, `created_at`) VALUES
	('cccccccc-cccc-cccc-cccc-ccccccccccc1', '11111111-1111-1111-1111-111111111111', '88888888-8888-8888-8888-888888888881', 'OPENING', 100.00, 'SYSTEM', NULL, 'Stok awal', '2416854f-55e6-423c-a2ec-8154c9431cd6', '2026-05-18 09:24:30'),
	('cccccccc-cccc-cccc-cccc-ccccccccccc2', '11111111-1111-1111-1111-111111111111', '88888888-8888-8888-8888-888888888881', 'SALE', -1.00, 'SALE', 'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa', 'Penjualan invoice INV-20260518-0001', '2416854f-55e6-423c-a2ec-8154c9431cd6', '2026-05-18 09:24:30');

-- Dumping structure for table umkm_odod.tenants
CREATE TABLE IF NOT EXISTS `tenants` (
  `id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `phone` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `address` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table umkm_odod.tenants: ~0 rows (approximately)
INSERT INTO `tenants` (`id`, `name`, `phone`, `address`, `created_at`, `updated_at`, `deleted_at`) VALUES
	('11111111-1111-1111-1111-111111111111', 'Keripik Mang Odod', '081234567890', 'Jalan Kademangan No. 25, Pojok, Kelurahan Setiamanah, Kota Cimahi', '2026-05-18 09:20:02', '2026-05-18 09:28:00', NULL),
	('11111111-1111-1111-1111-111111111112', 'Bakso Petruk', '082233445566', 'Jl. Sudirman No. 12, Bandung', '2026-05-18 09:57:40', '2026-05-18 14:51:21', NULL),
	('b83ed04b-f5b0-4109-a54d-de2916da7e0b', 'Aay Guevara', '081172638373', 'Jl. Kamarung no 12', '2026-05-18 15:08:31', '2026-05-18 15:09:48', NULL),
	('f27e441f-5385-4b8d-b2e2-88b8615a4634', 'Bakso Solo', '123456788', 'Jl. Singgasana Pradana No.111', '2026-05-18 10:29:53', '2026-05-18 15:09:49', NULL);

-- Dumping structure for table umkm_odod.users
CREATE TABLE IF NOT EXISTS `users` (
  `id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `tenant_id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `role_id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `full_name` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `username` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `password` char(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '0',
  `phone` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `is_active` tinyint(1) NOT NULL DEFAULT '1',
  `last_login_at` timestamp NULL DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_users_username_per_tenant` (`tenant_id`,`username`),
  KEY `fk_users_role` (`role_id`),
  CONSTRAINT `fk_users_role` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`),
  CONSTRAINT `fk_users_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table umkm_odod.users: ~2 rows (approximately)
INSERT INTO `users` (`id`, `tenant_id`, `role_id`, `full_name`, `username`, `password`, `phone`, `is_active`, `last_login_at`, `created_at`, `updated_at`, `deleted_at`) VALUES
	('2416854f-55e6-423c-a2ec-8154c9431cd6', 'f27e441f-5385-4b8d-b2e2-88b8615a4634', 'dc064501-5798-4c91-8308-0198561ceae3', 'Budi Karmana', 'budi.admin', '$2a$10$NgbIGU78TzIZ.n2An5RIJe557VhXcHLzlQy1vOl1lVIUx4h4CVKnm', '087718273837', 1, NULL, '2026-05-20 06:35:42', '2026-05-20 06:35:42', NULL),
	('9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', 'f27e441f-5385-4b8d-b2e2-88b8615a4634', '22222222-2222-2222-2222-222222222221', 'Heru Wibowo', 'heru.owner', '$2a$10$wuBaDA4Tqk4Tv7pSscEPwed2mB.KovlUW5Mf6xf979MWu3Zm0kTBm', '081345712017', 1, NULL, '2026-05-20 06:47:35', '2026-05-20 06:47:35', NULL);

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
