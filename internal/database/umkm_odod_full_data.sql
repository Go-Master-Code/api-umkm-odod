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

-- Dumping data for table umkm_odod.catalog_categories: ~3 rows (approximately)
INSERT INTO `catalog_categories` (`id`, `tenant_id`, `name`, `created_at`, `updated_at`, `deleted_at`) VALUES
	('08e850d4-6590-4cef-91bd-df349c271019', '11111111-1111-1111-1111-111111111111', 'Kue Basah', '2026-05-20 07:54:18', '2026-05-20 07:54:18', NULL),
	('44444444-4444-4444-4444-444444444441', '11111111-1111-1111-1111-111111111111', 'Snack', '2026-05-18 09:21:24', '2026-05-20 07:22:22', NULL),
	('44444444-4444-4444-4444-444444444442', '11111111-1111-1111-1111-111111111111', 'Minuman', '2026-05-18 09:21:24', '2026-05-20 06:50:33', NULL);

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

-- Dumping data for table umkm_odod.catalog_items: ~4 rows (approximately)
INSERT INTO `catalog_items` (`id`, `tenant_id`, `category_id`, `name`, `description`, `is_active`, `created_at`, `updated_at`, `deleted_at`) VALUES
	('55555555-5555-5555-5555-555555555551', '11111111-1111-1111-1111-111111111111', '44444444-4444-4444-4444-444444444441', 'Keripik Singkong', 'Keripik singkong premium', 1, '2026-05-18 09:21:44', '2026-05-20 07:34:53', NULL),
	('55555555-5555-5555-5555-555555555552', '11111111-1111-1111-1111-111111111111', '44444444-4444-4444-4444-444444444441', 'Keripik Pisang', 'Keripik pisang renyah', 1, '2026-05-18 09:21:44', '2026-05-20 07:34:23', NULL),
	('592fcb7b-4451-46e9-ba1d-3039a70ee61d', 'f27e441f-5385-4b8d-b2e2-88b8615a4634', '44444444-4444-4444-4444-444444444442', 'Fruit tea', 'Rasa Black Currant', 0, '2026-05-20 07:37:28', '2026-05-20 07:49:08', NULL),
	('c1124d3b-2894-420a-8878-556f96af06db', '11111111-1111-1111-1111-111111111111', '08e850d4-6590-4cef-91bd-df349c271019', 'Nagasari', 'Nagasari pisang ambon', 1, '2026-05-20 07:55:51', '2026-05-20 07:55:51', NULL);

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

-- Dumping data for table umkm_odod.item_attribute_values: ~4 rows (approximately)
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
  `minimum_stock` decimal(18,2) NOT NULL DEFAULT (0),
  `cost_price` decimal(18,2) NOT NULL DEFAULT '0.00',
  `selling_price` decimal(18,2) NOT NULL DEFAULT '0.00',
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

-- Dumping data for table umkm_odod.item_variants: ~3 rows (approximately)
INSERT INTO `item_variants` (`id`, `tenant_id`, `item_id`, `sku`, `barcode`, `variant_name`, `minimum_stock`, `cost_price`, `selling_price`, `is_active`, `created_at`, `updated_at`, `deleted_at`) VALUES
	('76a3c5f8-4340-42a4-8aa0-db591a791c58', '11111111-1111-1111-1111-111111111111', '55555555-5555-5555-5555-555555555551', 'KS-PB-500', '899100000003', 'Pedas Balado', 5.00, 13000.00, 15000.00, 1, '2026-05-20 10:45:44', '2026-06-02 08:23:37', NULL),
	('88888888-8888-8888-8888-888888888881', '11111111-1111-1111-1111-111111111111', '55555555-5555-5555-5555-555555555551', 'KS-PDJ-250', '899100000001', 'Pedas Daun Jeruk 250gr', 5.00, 13000.00, 15000.00, 1, '2026-05-18 09:22:19', '2026-06-02 08:23:38', NULL),
	('88888888-8888-8888-8888-888888888882', '11111111-1111-1111-1111-111111111111', '55555555-5555-5555-5555-555555555551', 'KS-AB-500', '899100000002', 'Asin Bawang 500gr', 5.00, 12000.00, 14000.00, 1, '2026-05-18 09:22:19', '2026-06-02 08:23:39', NULL);

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

-- Dumping structure for table umkm_odod.purchases
CREATE TABLE IF NOT EXISTS `purchases` (
  `id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `tenant_id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `purchase_number` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `supplier_id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `invoice_number` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `subtotal` decimal(18,2) NOT NULL DEFAULT '0.00',
  `discount_amount` decimal(18,2) NOT NULL DEFAULT '0.00',
  `tax_amount` decimal(18,2) NOT NULL DEFAULT '0.00',
  `grand_total` decimal(18,2) NOT NULL DEFAULT '0.00',
  `notes` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `created_by` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uk_purchase_number` (`purchase_number`) USING BTREE,
  KEY `idx_purchase_tenant` (`tenant_id`) USING BTREE,
  KEY `idx_purchase_supplier` (`supplier_id`) USING BTREE,
  KEY `idx_purchase_created_by` (`created_by`) USING BTREE,
  KEY `idx_purchase_created_at` (`created_at`) USING BTREE,
  CONSTRAINT `fk_purchase_created_by` FOREIGN KEY (`created_by`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_purchase_supplier` FOREIGN KEY (`supplier_id`) REFERENCES `suppliers` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_purchase_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table umkm_odod.purchases: ~0 rows (approximately)

-- Dumping structure for table umkm_odod.purchase_items
CREATE TABLE IF NOT EXISTS `purchase_items` (
  `id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `tenant_id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `purchase_id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `item_variant_id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `item_name_snapshot` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `variant_name_snapshot` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `sku_snapshot` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `qty` decimal(18,2) NOT NULL,
  `cost_price` decimal(18,2) NOT NULL,
  `discount_amount` decimal(18,2) NOT NULL DEFAULT '0.00',
  `subtotal` decimal(18,2) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_purchase_item_tenant` (`tenant_id`) USING BTREE,
  KEY `idx_purchase_item_purchase` (`purchase_id`) USING BTREE,
  KEY `idx_purchase_item_variant` (`item_variant_id`) USING BTREE,
  CONSTRAINT `fk_purchase_item_purchase` FOREIGN KEY (`purchase_id`) REFERENCES `purchases` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_purchase_item_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_purchase_item_variant` FOREIGN KEY (`item_variant_id`) REFERENCES `item_variants` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table umkm_odod.purchase_items: ~0 rows (approximately)

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

-- Dumping data for table umkm_odod.roles: ~5 rows (approximately)
INSERT INTO `roles` (`id`, `tenant_id`, `name`, `created_at`, `updated_at`, `deleted_at`) VALUES
	('22222222-2222-2222-2222-222222222221', '11111111-1111-1111-1111-111111111112', 'OWNER', '2026-05-18 09:20:36', '2026-05-18 16:35:19', NULL),
	('22222222-2222-2222-2222-222222222222', '11111111-1111-1111-1111-111111111111', 'CASHIER', '2026-05-18 09:20:36', '2026-05-18 09:20:36', NULL),
	('22222222-2222-2222-2222-222222222223', '11111111-1111-1111-1111-111111111111', 'ADMIN', '2026-06-02 10:01:02', '2026-06-02 10:01:02', NULL),
	('ab88f666-b2d6-482c-bdcd-be29915e5395', 'f27e441f-5385-4b8d-b2e2-88b8615a4634', 'SUPER_ADMIN', '2026-05-19 02:02:01', '2026-05-28 10:32:27', NULL),
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

-- Dumping data for table umkm_odod.sales: ~3 rows (approximately)
INSERT INTO `sales` (`id`, `tenant_id`, `invoice_number`, `customer_name`, `cashier_id`, `subtotal`, `discount_amount`, `tax_amount`, `grand_total`, `payment_method`, `payment_status`, `notes`, `created_at`, `updated_at`) VALUES
	('4cf1c87a-0014-47e6-9b46-d3f4ec980377', '11111111-1111-1111-1111-111111111111', 'INV-1779948538', 'Budi', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', 27000.00, 500.00, 2700.00, 29200.00, 'CASH', 'PAID', 'Pembelian 1 item', '2026-05-28 06:08:58', '2026-05-28 06:08:58'),
	('bcec37b4-a8bb-44a2-a039-cbe8cb06a530', '11111111-1111-1111-1111-111111111111', 'INV-1779941000', 'Budi Santoso', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', 104000.00, 5000.00, 10400.00, 109400.00, 'QRIS', 'PAID', 'Pembelian sore hari', '2026-05-28 04:03:21', '2026-05-28 04:03:21'),
	('ecf9ac1a-6c21-4881-80d5-b741e62ad205', '11111111-1111-1111-1111-111111111111', 'INV-1779948742', 'Andi', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', 437000.00, 10000.00, 43700.00, 470700.00, 'TRANSFER', 'UNPAID', 'Pembelian grosir', '2026-05-28 06:12:22', '2026-05-28 06:51:41');

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

-- Dumping data for table umkm_odod.sale_items: ~6 rows (approximately)
INSERT INTO `sale_items` (`id`, `tenant_id`, `sale_id`, `item_variant_id`, `item_name_snapshot`, `variant_name_snapshot`, `sku_snapshot`, `qty`, `unit_price`, `discount_amount`, `subtotal`, `created_at`) VALUES
	('13574740-8adf-4885-842c-7009b8567392', '11111111-1111-1111-1111-111111111111', 'ecf9ac1a-6c21-4881-80d5-b741e62ad205', '88888888-8888-8888-8888-888888888882', 'Keripik Singkong', 'Asin Bawang 500gr', 'KS-AB-500', 10.00, 14000.00, 2000.00, 138000.00, '2026-05-28 06:12:22'),
	('1d10f8cb-e35d-4cab-9035-762eeee4cfdb', '11111111-1111-1111-1111-111111111111', 'bcec37b4-a8bb-44a2-a039-cbe8cb06a530', '88888888-8888-8888-8888-888888888881', 'Keripik Singkong', 'Pedas Daun Jeruk 250gr', 'KS-PDJ-250', 5.00, 15000.00, 0.00, 75000.00, '2026-05-28 04:03:21'),
	('1f17845a-f5b6-48b2-a93c-b45daae9ccfa', '11111111-1111-1111-1111-111111111111', 'ecf9ac1a-6c21-4881-80d5-b741e62ad205', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'Keripik Singkong', 'Pedas Balado', 'KS-PB-500', 7.00, 15000.00, 1000.00, 104000.00, '2026-05-28 06:12:22'),
	('8bca2936-1417-4d3a-b2d4-4cd119637451', '11111111-1111-1111-1111-111111111111', 'bcec37b4-a8bb-44a2-a039-cbe8cb06a530', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'Keripik Singkong', 'Pedas Balado', 'KS-PB-500', 2.00, 15000.00, 1000.00, 29000.00, '2026-05-28 04:03:21'),
	('adcb2348-f629-4dda-a8c0-535039e03942', '11111111-1111-1111-1111-111111111111', '4cf1c87a-0014-47e6-9b46-d3f4ec980377', '88888888-8888-8888-8888-888888888882', 'Keripik Singkong', 'Asin Bawang 500gr', 'KS-AB-500', 2.00, 14000.00, 1000.00, 27000.00, '2026-05-28 06:08:58'),
	('beddda53-3faa-46b7-a4dc-6fb00befb751', '11111111-1111-1111-1111-111111111111', 'ecf9ac1a-6c21-4881-80d5-b741e62ad205', '88888888-8888-8888-8888-888888888881', 'Keripik Singkong', 'Pedas Daun Jeruk 250gr', 'KS-PDJ-250', 13.00, 15000.00, 0.00, 195000.00, '2026-05-28 06:12:22');

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

-- Dumping data for table umkm_odod.stock_movements: ~51 rows (approximately)
INSERT INTO `stock_movements` (`id`, `tenant_id`, `item_variant_id`, `movement_type`, `qty`, `reference_type`, `reference_id`, `notes`, `created_by`, `created_at`) VALUES
	('1012475d-2a7f-4fd5-88dc-8e399e3ee3c6', '11111111-1111-1111-1111-111111111111', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'SALE', -10.00, '', '', 'Bapak Udin', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-21 09:23:40'),
	('119de780-c9bb-48fc-bbad-3fd8ab8b2118', '11111111-1111-1111-1111-111111111111', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'SALE', -7.00, 'SALE', 'ecf9ac1a-6c21-4881-80d5-b741e62ad205', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-28 06:12:22'),
	('1776a9be-4f3a-4a9e-9c89-ee71eaee709f', '11111111-1111-1111-1111-111111111111', '88888888-8888-8888-8888-888888888881', 'SALE', -5.00, 'SALE', '0f79d701-b78f-452e-a0cc-62c7d2ac35e8', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-26 13:20:06'),
	('1e38c273-cb4b-47db-be33-9a1f53739337', '11111111-1111-1111-1111-111111111111', '88888888-8888-8888-8888-888888888881', 'SALE', -5.00, 'SALE', 'bcec37b4-a8bb-44a2-a039-cbe8cb06a530', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-28 04:03:21'),
	('1f880299-0e64-4c4b-a3e8-15be0ab2052a', '11111111-1111-1111-1111-111111111111', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'REDUCE', -3.00, 'ADJUSTMENT', '', 'Reason: kencing tikus | Notes: plastik berbau', '64ee38a5-23c6-4a82-b1ad-d488ccc0d8e6', '2026-05-29 02:06:48'),
	('204020ba-96b8-4394-8033-26d0cd6cd286', '11111111-1111-1111-1111-111111111111', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'SALE', -2.00, 'SALE', '2d61cc53-394b-4ab8-9354-dbd8f7212150', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-26 13:23:39'),
	('219bc491-cc1e-4dcd-be2b-f57016e12bb5', '11111111-1111-1111-1111-111111111111', '88888888-8888-8888-8888-888888888881', 'SALE', -1.00, 'SALE', '9b34dc3c-db00-400b-a09a-d0eb9ef21d52', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-26 13:06:49'),
	('251daba8-f80e-4f18-8bfc-1fca2da08c50', '11111111-1111-1111-1111-111111111111', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'SALE', -2.00, 'SALE', '0565b794-0d7f-4ce4-92d6-13aa5cd6007f', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-26 13:19:06'),
	('26ffdab0-e6b7-4103-8110-48e44b38eebe', '11111111-1111-1111-1111-111111111111', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'SALE', -3.00, 'SALE', '609c139e-4a73-4453-8462-c6e3d39150ba', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-26 12:35:57'),
	('27514e4a-a357-45ba-adc3-9275289b22e8', '11111111-1111-1111-1111-111111111111', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'SALE', -2.00, 'SALE', 'b5c9ed63-63f0-46c1-95e3-c77fb096049d', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-26 13:26:25'),
	('27c2aa35-82aa-4410-b218-43213e8e9aa0', '11111111-1111-1111-1111-111111111111', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'SALE', -5.00, 'SALE', '58171b7c-b6f8-49ee-b4a3-41d8fa84f5f9', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-26 12:05:04'),
	('28bc8dbc-c94a-403e-bbaf-d8318a6617f7', '11111111-1111-1111-1111-111111111111', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'SALE', -2.00, 'SALE', '1500c864-83a8-49e0-a704-f962ee94fe57', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-26 13:25:09'),
	('28c420a2-846c-45e0-9ed3-4f74c6799d2d', '11111111-1111-1111-1111-111111111111', '88888888-8888-8888-8888-888888888881', 'PURCHASE', 150.00, '', '', 'Stock Awal', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-28 06:11:41'),
	('2d6f8aac-8b17-4368-b744-a1b14b917823', '11111111-1111-1111-1111-111111111111', '88888888-8888-8888-8888-888888888881', 'SALE', -5.00, 'SALE', '1500c864-83a8-49e0-a704-f962ee94fe57', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-26 13:25:09'),
	('33c0109a-37fc-4b85-a02c-501e193c5e35', '11111111-1111-1111-1111-111111111111', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'SALE', -10.00, 'SALE', '55f23f83-a2d0-45a2-a913-078c617d04c8', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-26 12:32:07'),
	('3d73a734-e367-4b76-a120-499f73fb1b4e', '11111111-1111-1111-1111-111111111111', '88888888-8888-8888-8888-888888888881', 'SALE', -5.00, 'SALE', '974dd35c-53fe-4a40-ae94-2f8ca4229c48', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-26 13:21:27'),
	('4de4596b-b553-43ee-9414-af7f78a1625c', '11111111-1111-1111-1111-111111111111', '88888888-8888-8888-8888-888888888881', 'SALE', -1.00, 'SALE', 'db3c1825-a8b3-449a-a999-f7e5dc8c1e6e', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-26 13:05:13'),
	('54e130e8-f38a-4612-86d5-c4a1df869fbe', '11111111-1111-1111-1111-111111111111', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'PURCHASE', 15.00, '', '', 'Repack kedua', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-21 09:15:48'),
	('58481476-b28d-469b-88bb-88328d67568e', '11111111-1111-1111-1111-111111111111', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'SALE', -2.00, 'SALE', 'c4585fcf-5364-479d-b6dd-5ab2d414fd10', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-26 13:12:06'),
	('5c044814-78e5-4ce0-8589-72377b417899', '11111111-1111-1111-1111-111111111111', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'SALE', -2.00, 'SALE', '49e45d2b-7a13-4906-a292-b7d870bb1a8e', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-26 13:15:35'),
	('5c2fb1ed-7339-473c-ba3f-4608a94d6ad1', '11111111-1111-1111-1111-111111111111', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'SALE', -13.00, 'SALE', 'b50b0a06-422e-446c-865c-643a43beab27', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-26 12:22:20'),
	('5fa4dd52-b554-4eec-8ddd-1e08156b6425', '11111111-1111-1111-1111-111111111111', '88888888-8888-8888-8888-888888888881', 'SALE', -13.00, 'SALE', 'ecf9ac1a-6c21-4881-80d5-b741e62ad205', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-28 06:12:22'),
	('64cddab1-8093-41f9-9df3-790b76eb8a74', '11111111-1111-1111-1111-111111111111', '88888888-8888-8888-8888-888888888881', 'SALE', -1.00, 'SALE', 'c4585fcf-5364-479d-b6dd-5ab2d414fd10', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-26 13:12:06'),
	('6eace372-d348-422a-8335-a6c7bb44690d', '11111111-1111-1111-1111-111111111111', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'SALE', -2.00, 'SALE', 'db3c1825-a8b3-449a-a999-f7e5dc8c1e6e', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-26 13:05:13'),
	('704bf63c-9bfb-4cd2-af8b-726897576591', '11111111-1111-1111-1111-111111111111', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'SALE', -4.00, 'SALE', 'c8a4629d-c6fc-46c8-bc9d-608bec8c554e', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-26 12:16:02'),
	('705e9eed-9def-4588-98b6-66cbd32e83d4', '11111111-1111-1111-1111-111111111111', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'SALE', -2.00, 'SALE', '9b34dc3c-db00-400b-a09a-d0eb9ef21d52', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-26 13:06:49'),
	('745b00da-60ad-48dd-8016-737c9af4a588', '11111111-1111-1111-1111-111111111111', '88888888-8888-8888-8888-888888888881', 'SALE', -5.00, 'SALE', '2d61cc53-394b-4ab8-9354-dbd8f7212150', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-26 13:23:39'),
	('7bb76694-03f8-4b12-b00c-ebf66213c8ff', '11111111-1111-1111-1111-111111111111', '88888888-8888-8888-8888-888888888882', 'SALE', -10.00, 'SALE', 'ecf9ac1a-6c21-4881-80d5-b741e62ad205', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-28 06:12:22'),
	('81b5b7b6-a4b7-40b5-8762-b70c338e1d19', '11111111-1111-1111-1111-111111111111', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'SALE', -2.00, 'SALE', '974dd35c-53fe-4a40-ae94-2f8ca4229c48', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-26 13:21:27'),
	('884ede0d-7385-4224-81a5-9e9ffc23fe17', '11111111-1111-1111-1111-111111111111', '88888888-8888-8888-8888-888888888881', 'SALE', -5.00, 'SALE', 'b5c9ed63-63f0-46c1-95e3-c77fb096049d', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-26 13:26:25'),
	('8df813f3-3e51-49c2-9396-c244c53ccc4c', '11111111-1111-1111-1111-111111111111', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'SALE', -2.00, 'SALE', '7564b59d-35cd-4b52-afe4-deddd48456c9', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-26 13:24:41'),
	('9ed20b5d-9f8c-42e8-b36a-d853c75bfb68', '11111111-1111-1111-1111-111111111111', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'SALE', -11.00, 'SALE', '98dc6b96-260b-4b08-9bc3-6ef84025706d', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-26 12:59:47'),
	('a3ab7610-6093-4ea4-9a45-6c3330d33ad0', '11111111-1111-1111-1111-111111111111', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'SALE', -1.00, 'SALE', 'c0988bf7-cd16-49ce-a33d-9ecb54d506db', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-26 12:12:46'),
	('a47a1f1a-246b-408c-8402-b206118ce58d', '11111111-1111-1111-1111-111111111111', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'SALE', -3.00, '', '', 'Haji Sapari', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-21 09:25:59'),
	('a702ef07-ac1a-4765-9ccf-1dd300616435', '11111111-1111-1111-1111-111111111111', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'SALE', -2.00, 'SALE', '6a208d10-09a5-4e0b-8abe-c293602f0abe', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-26 11:58:26'),
	('a718cc1a-59de-438d-b5db-a14a5b2a8d6f', '11111111-1111-1111-1111-111111111111', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'SALE', -2.00, 'SALE', '0f79d701-b78f-452e-a0cc-62c7d2ac35e8', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-26 13:20:06'),
	('ac65475f-ad0f-4294-8781-b0928654459d', '11111111-1111-1111-1111-111111111111', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'PURCHASE', 13.00, '', '', 'Hasil repack', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-21 09:14:27'),
	('b8d6362f-4cb1-47ff-b704-29609046b15f', '11111111-1111-1111-1111-111111111111', '88888888-8888-8888-8888-888888888882', 'SALE', -2.00, 'SALE', '4cf1c87a-0014-47e6-9b46-d3f4ec980377', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-28 06:08:58'),
	('bc8b4ad8-dc92-43b4-b09f-11e0094bbf39', '11111111-1111-1111-1111-111111111111', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'SALE', -13.00, 'SALE', 'abbeae5a-20a5-44d5-883c-065773b4e524', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-26 12:53:06'),
	('c82aec49-2d32-43a8-9f6a-fe4397fc9394', '11111111-1111-1111-1111-111111111111', '88888888-8888-8888-8888-888888888881', 'SALE', -5.00, 'SALE', '0565b794-0d7f-4ce4-92d6-13aa5cd6007f', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-26 13:19:06'),
	('cccccccc-cccc-cccc-cccc-ccccccccccc1', '11111111-1111-1111-1111-111111111111', '88888888-8888-8888-8888-888888888881', 'OPENING', 100.00, 'SYSTEM', NULL, 'Stok awal', '2416854f-55e6-423c-a2ec-8154c9431cd6', '2026-05-18 09:24:30'),
	('cccccccc-cccc-cccc-cccc-ccccccccccc2', '11111111-1111-1111-1111-111111111111', '88888888-8888-8888-8888-888888888881', 'SALE', -1.00, 'SALE', 'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa', 'Penjualan invoice INV-20260518-0001', '2416854f-55e6-423c-a2ec-8154c9431cd6', '2026-05-18 09:24:30'),
	('d0f869a8-ed0c-4c53-8798-f7a2098d3baa', '11111111-1111-1111-1111-111111111111', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'REDUCE', -2.00, 'ADJUSTMENT', '', 'plastik terbuka', '64ee38a5-23c6-4a82-b1ad-d488ccc0d8e6', '2026-05-28 16:24:38'),
	('d426f189-813d-4501-804d-456710fa3543', '11111111-1111-1111-1111-111111111111', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'SALE', -20.00, '', '', 'Ibu Linda', '64ee38a5-23c6-4a82-b1ad-d488ccc0d8e6', '2026-06-02 08:25:23'),
	('d908d5e0-5ff5-486b-874b-47df2622b965', '11111111-1111-1111-1111-111111111111', '88888888-8888-8888-8888-888888888882', 'PURCHASE', 500.00, '', '', 'Stock Awal', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-28 06:08:47'),
	('e0b28d2a-51f7-4d05-a243-11cea30dd8c2', '11111111-1111-1111-1111-111111111111', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'PURCHASE', 99.00, '', '', 'Stock awal', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-21 09:02:13'),
	('e331dfb6-6a53-4eea-8869-3447d116b043', '11111111-1111-1111-1111-111111111111', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'ADD', 10.00, 'ADJUSTMENT', '', 'barang masuk angin', '64ee38a5-23c6-4a82-b1ad-d488ccc0d8e6', '2026-05-28 16:23:07'),
	('e7bc843e-7178-4148-8ea0-417db75819f2', '11111111-1111-1111-1111-111111111111', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'REDUCE', -1.00, 'ADJUSTMENT', '', 'plastik terbuka', '64ee38a5-23c6-4a82-b1ad-d488ccc0d8e6', '2026-05-29 01:56:02'),
	('eae61992-c31c-4b65-99f4-fbe568378f1f', '11111111-1111-1111-1111-111111111111', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'SALE', -2.00, 'SALE', 'b4a64b8a-c2ac-4c57-8a67-14e65058a7e4', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-26 13:08:45'),
	('f3c363d4-0425-44a0-a1ea-80475000f3b2', '11111111-1111-1111-1111-111111111111', '88888888-8888-8888-8888-888888888881', 'SALE', -1.00, 'SALE', 'b4a64b8a-c2ac-4c57-8a67-14e65058a7e4', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-26 13:08:45'),
	('f7802881-51b8-4cd5-8e80-86aa73110fd8', '11111111-1111-1111-1111-111111111111', '88888888-8888-8888-8888-888888888881', 'SALE', -5.00, 'SALE', '49e45d2b-7a13-4906-a292-b7d870bb1a8e', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-26 13:15:35'),
	('f8dea795-4049-415d-8e19-f0d0e8d7e45b', '11111111-1111-1111-1111-111111111111', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'SALE', -2.00, 'SALE', 'bcec37b4-a8bb-44a2-a039-cbe8cb06a530', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-28 04:03:21'),
	('f90c5b7d-497f-4cd1-9385-4583a19b7aab', '11111111-1111-1111-1111-111111111111', '88888888-8888-8888-8888-888888888881', 'SALE', -5.00, 'SALE', '7564b59d-35cd-4b52-afe4-deddd48456c9', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-26 13:24:41');

-- Dumping structure for table umkm_odod.suppliers
CREATE TABLE IF NOT EXISTS `suppliers` (
  `id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `tenant_id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` varchar(150) COLLATE utf8mb4_unicode_ci NOT NULL,
  `phone` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `address` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `is_active` tinyint(1) NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT (now()),
  `updated_at` timestamp NULL DEFAULT (now()) ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_supplier_name_per_tenant` (`tenant_id`,`name`),
  KEY `idx_suppliers_tenant` (`tenant_id`),
  KEY `idx_suppliers_name` (`name`),
  KEY `idx_suppliers_active` (`is_active`),
  CONSTRAINT `fk_suppliers_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table umkm_odod.suppliers: ~3 rows (approximately)
INSERT INTO `suppliers` (`id`, `tenant_id`, `name`, `phone`, `address`, `is_active`, `created_at`, `updated_at`, `deleted_at`) VALUES
	('4e0c76e8-e258-440d-9ed3-a8831f3c3296', '11111111-1111-1111-1111-111111111111', 'Lusi Oktaviani', '089981451237', 'Jl. Kalidam No. 45', 1, '2026-06-02 08:32:51', '2026-06-02 08:32:51', NULL),
	('4efba307-0fbd-4501-bb8d-3f9212c14fb9', '11111111-1111-1111-1111-111111111111', 'Keripik Ma Icih', '081248150369', 'Jl. Pasirkaliki No. 33', 1, '2026-06-02 10:04:20', '2026-06-02 10:04:20', NULL),
	('98f6686a-48b9-478e-b709-e8177a483bb8', '11111111-1111-1111-1111-111111111111', 'Amir Machmud', '081547841302', 'Jl. Cibaduyut No. 121', 1, '2026-06-02 11:10:00', '2026-06-02 11:25:59', '2026-06-02 11:25:59');

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

-- Dumping data for table umkm_odod.tenants: ~4 rows (approximately)
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

-- Dumping data for table umkm_odod.users: ~4 rows (approximately)
INSERT INTO `users` (`id`, `tenant_id`, `role_id`, `full_name`, `username`, `password`, `phone`, `is_active`, `last_login_at`, `created_at`, `updated_at`, `deleted_at`) VALUES
	('2416854f-55e6-423c-a2ec-8154c9431cd6', 'f27e441f-5385-4b8d-b2e2-88b8615a4634', 'dc064501-5798-4c91-8308-0198561ceae3', 'Budi Karmana', 'budi.admin', '$2a$10$NgbIGU78TzIZ.n2An5RIJe557VhXcHLzlQy1vOl1lVIUx4h4CVKnm', '087718273837', 1, NULL, '2026-05-20 06:35:42', '2026-05-20 06:35:42', NULL),
	('375b98d3-9d07-483c-b9ea-7b2ce088ae48', 'f27e441f-5385-4b8d-b2e2-88b8615a4634', 'ab88f666-b2d6-482c-bdcd-be29915e5395', 'Kevin SA', 'kevin.sa', '$2a$10$ERZtRhd/MuxKn6pvMkFuS.8TfxCdyC4B8Zn4C4PYsOvspoPbCAzxW', '089945712017', 1, NULL, '2026-05-28 10:26:35', '2026-05-28 10:26:35', NULL),
	('64ee38a5-23c6-4a82-b1ad-d488ccc0d8e6', '11111111-1111-1111-1111-111111111111', '22222222-2222-2222-2222-222222222223', 'Kiki Sutisna', 'kiki.admin', '$2a$10$kN/7usrkKbIjHPPNZnLYCenuD7fLF3BKBbRTiLiIwLfP2yyVQfOVu', '081321114109', 1, NULL, '2026-05-28 09:39:41', '2026-06-02 10:01:17', NULL),
	('9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '11111111-1111-1111-1111-111111111111', '22222222-2222-2222-2222-222222222221', 'Heru Wibowo', 'heru.owner', '$2a$10$wuBaDA4Tqk4Tv7pSscEPwed2mB.KovlUW5Mf6xf979MWu3Zm0kTBm', '081345712017', 1, NULL, '2026-05-20 06:47:35', '2026-05-20 07:53:10', NULL);

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
