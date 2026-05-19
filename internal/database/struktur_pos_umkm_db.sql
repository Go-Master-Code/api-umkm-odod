-- --------------------------------------------------------
-- Host:                         localhost
-- Server version:               5.7.43-log - MySQL Community Server (GPL)
-- Server OS:                    Win64
-- HeidiSQL Version:             12.8.0.6908
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
CREATE DATABASE IF NOT EXISTS `umkm_odod` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci */;
USE `umkm_odod`;

-- Dumping structure for table umkm_odod.catalog_categories
CREATE TABLE IF NOT EXISTS `catalog_categories` (
  `id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `tenant_id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_categories_name_per_tenant` (`tenant_id`,`name`),
  CONSTRAINT `fk_categories_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Data exporting was unselected.

-- Dumping structure for table umkm_odod.catalog_items
CREATE TABLE IF NOT EXISTS `catalog_items` (
  `id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `tenant_id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `category_id` char(36) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `name` varchar(200) COLLATE utf8mb4_unicode_ci NOT NULL,
  `description` text COLLATE utf8mb4_unicode_ci,
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

-- Data exporting was unselected.

-- Dumping structure for table umkm_odod.item_attributes
CREATE TABLE IF NOT EXISTS `item_attributes` (
  `id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `tenant_id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_item_attributes_name_per_tenant` (`tenant_id`,`name`),
  CONSTRAINT `fk_item_attributes_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Data exporting was unselected.

-- Dumping structure for table umkm_odod.item_attribute_values
CREATE TABLE IF NOT EXISTS `item_attribute_values` (
  `id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `tenant_id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `attribute_id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `value` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_attribute_values_tenant` (`tenant_id`),
  KEY `fk_attribute_values_attribute` (`attribute_id`),
  CONSTRAINT `fk_attribute_values_attribute` FOREIGN KEY (`attribute_id`) REFERENCES `item_attributes` (`id`),
  CONSTRAINT `fk_attribute_values_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Data exporting was unselected.

-- Dumping structure for table umkm_odod.item_variants
CREATE TABLE IF NOT EXISTS `item_variants` (
  `id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `tenant_id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `item_id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `sku` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `barcode` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `variant_name` varchar(150) COLLATE utf8mb4_unicode_ci NOT NULL,
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

-- Data exporting was unselected.

-- Dumping structure for table umkm_odod.item_variant_attribute_values
CREATE TABLE IF NOT EXISTS `item_variant_attribute_values` (
  `variant_id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `attribute_value_id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `tenant_id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`variant_id`,`attribute_value_id`),
  KEY `fk_variant_attribute_value` (`attribute_value_id`),
  KEY `FK_item_variant_attribute_values_tenants` (`tenant_id`),
  CONSTRAINT `FK_item_variant_attribute_values_tenants` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION,
  CONSTRAINT `fk_variant_attribute_value` FOREIGN KEY (`attribute_value_id`) REFERENCES `item_attribute_values` (`id`),
  CONSTRAINT `fk_variant_attribute_variant` FOREIGN KEY (`variant_id`) REFERENCES `item_variants` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Data exporting was unselected.

-- Dumping structure for table umkm_odod.price_histories
CREATE TABLE IF NOT EXISTS `price_histories` (
  `id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `tenant_id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `item_variant_id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `price_type` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `price` decimal(18,2) NOT NULL,
  `effective_date` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `created_by` char(36) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `fk_prices_tenant` (`tenant_id`),
  KEY `fk_prices_variant` (`item_variant_id`),
  KEY `fk_prices_created_by` (`created_by`),
  CONSTRAINT `fk_prices_created_by` FOREIGN KEY (`created_by`) REFERENCES `users` (`id`),
  CONSTRAINT `fk_prices_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`),
  CONSTRAINT `fk_prices_variant` FOREIGN KEY (`item_variant_id`) REFERENCES `item_variants` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Data exporting was unselected.

-- Dumping structure for table umkm_odod.roles
CREATE TABLE IF NOT EXISTS `roles` (
  `id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `tenant_id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_roles_name_per_tenant` (`tenant_id`,`name`),
  CONSTRAINT `fk_roles_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Data exporting was unselected.

-- Dumping structure for table umkm_odod.sales
CREATE TABLE IF NOT EXISTS `sales` (
  `id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `tenant_id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `invoice_number` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `customer_name` varchar(150) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `cashier_id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `subtotal` decimal(18,2) NOT NULL DEFAULT '0.00',
  `discount_amount` decimal(18,2) NOT NULL DEFAULT '0.00',
  `tax_amount` decimal(18,2) NOT NULL DEFAULT '0.00',
  `grand_total` decimal(18,2) NOT NULL DEFAULT '0.00',
  `payment_method` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL,
  `payment_status` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `notes` text COLLATE utf8mb4_unicode_ci,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_sales_invoice_per_tenant` (`tenant_id`,`invoice_number`),
  KEY `fk_sales_cashier` (`cashier_id`),
  CONSTRAINT `fk_sales_cashier` FOREIGN KEY (`cashier_id`) REFERENCES `users` (`id`),
  CONSTRAINT `fk_sales_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Data exporting was unselected.

-- Dumping structure for table umkm_odod.sale_items
CREATE TABLE IF NOT EXISTS `sale_items` (
  `id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `tenant_id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `sale_id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `item_variant_id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `item_name_snapshot` varchar(200) COLLATE utf8mb4_unicode_ci NOT NULL,
  `variant_name_snapshot` varchar(150) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `sku_snapshot` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
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

-- Data exporting was unselected.

-- Dumping structure for table umkm_odod.stock_movements
CREATE TABLE IF NOT EXISTS `stock_movements` (
  `id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `tenant_id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `item_variant_id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `movement_type` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `qty` decimal(18,2) NOT NULL,
  `reference_type` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `reference_id` char(36) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `notes` text COLLATE utf8mb4_unicode_ci,
  `created_by` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `fk_stock_tenant` (`tenant_id`),
  KEY `fk_stock_variant` (`item_variant_id`),
  KEY `fk_stock_created_by` (`created_by`),
  CONSTRAINT `fk_stock_created_by` FOREIGN KEY (`created_by`) REFERENCES `users` (`id`),
  CONSTRAINT `fk_stock_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`),
  CONSTRAINT `fk_stock_variant` FOREIGN KEY (`item_variant_id`) REFERENCES `item_variants` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Data exporting was unselected.

-- Dumping structure for table umkm_odod.tenants
CREATE TABLE IF NOT EXISTS `tenants` (
  `id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` varchar(150) COLLATE utf8mb4_unicode_ci NOT NULL,
  `phone` varchar(30) COLLATE utf8mb4_unicode_ci NOT NULL,
  `address` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Data exporting was unselected.

-- Dumping structure for table umkm_odod.users
CREATE TABLE IF NOT EXISTS `users` (
  `id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `tenant_id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `role_id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `full_name` varchar(150) COLLATE utf8mb4_unicode_ci NOT NULL,
  `username` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `password` char(60) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '0',
  `phone` varchar(30) COLLATE utf8mb4_unicode_ci NOT NULL,
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

-- Data exporting was unselected.

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
