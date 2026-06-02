CREATE TABLE `purchases` (
	`id` CHAR(36) NOT NULL COLLATE 'utf8mb4_unicode_ci',
	`tenant_id` CHAR(36) NOT NULL COLLATE 'utf8mb4_unicode_ci',
	`purchase_number` VARCHAR(100) NOT NULL COLLATE 'utf8mb4_unicode_ci',
	`supplier_id` CHAR(36) NOT NULL COLLATE 'utf8mb4_unicode_ci',
	`invoice_number` VARCHAR(100) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
	`subtotal` DECIMAL(18,2) NOT NULL DEFAULT '0.00',
	`discount_amount` DECIMAL(18,2) NOT NULL DEFAULT '0.00',
	`tax_amount` DECIMAL(18,2) NOT NULL DEFAULT '0.00',
	`grand_total` DECIMAL(18,2) NOT NULL DEFAULT '0.00',
	`notes` TEXT NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
	`created_by` CHAR(36) NOT NULL COLLATE 'utf8mb4_unicode_ci',
	`created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`) USING BTREE,
	UNIQUE INDEX `uk_purchase_number` (`purchase_number`) USING BTREE,
	INDEX `idx_purchase_tenant` (`tenant_id`) USING BTREE,
	INDEX `idx_purchase_supplier` (`supplier_id`) USING BTREE,
	INDEX `idx_purchase_created_by` (`created_by`) USING BTREE,
	INDEX `idx_purchase_created_at` (`created_at`) USING BTREE,
	CONSTRAINT `fk_purchase_created_by` FOREIGN KEY (`created_by`) REFERENCES `users` (`id`) ON UPDATE RESTRICT ON DELETE RESTRICT,
	CONSTRAINT `fk_purchase_supplier` FOREIGN KEY (`supplier_id`) REFERENCES `suppliers` (`id`) ON UPDATE RESTRICT ON DELETE RESTRICT,
	CONSTRAINT `fk_purchase_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON UPDATE RESTRICT ON DELETE RESTRICT
)
COLLATE='utf8mb4_unicode_ci'
ENGINE=InnoDB
;

CREATE TABLE `purchase_items` (
	`id` CHAR(36) NOT NULL COLLATE 'utf8mb4_unicode_ci',
	`tenant_id` CHAR(36) NOT NULL COLLATE 'utf8mb4_unicode_ci',
	`purchase_id` CHAR(36) NOT NULL COLLATE 'utf8mb4_unicode_ci',
	`item_variant_id` CHAR(36) NOT NULL COLLATE 'utf8mb4_unicode_ci',
	`item_name_snapshot` VARCHAR(200) NOT NULL COLLATE 'utf8mb4_unicode_ci',
	`variant_name_snapshot` VARCHAR(150) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
	`sku_snapshot` VARCHAR(100) NOT NULL COLLATE 'utf8mb4_unicode_ci',
	`qty` DECIMAL(18,2) NOT NULL,
	`cost_price` DECIMAL(18,2) NOT NULL,
	`discount_amount` DECIMAL(18,2) NOT NULL DEFAULT '0.00',
	`subtotal` DECIMAL(18,2) NOT NULL,
	`created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`) USING BTREE,
	INDEX `idx_purchase_item_tenant` (`tenant_id`) USING BTREE,
	INDEX `idx_purchase_item_purchase` (`purchase_id`) USING BTREE,
	INDEX `idx_purchase_item_variant` (`item_variant_id`) USING BTREE,
	CONSTRAINT `fk_purchase_item_purchase` FOREIGN KEY (`purchase_id`) REFERENCES `purchases` (`id`) ON UPDATE RESTRICT ON DELETE RESTRICT,
	CONSTRAINT `fk_purchase_item_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON UPDATE RESTRICT ON DELETE RESTRICT,
	CONSTRAINT `fk_purchase_item_variant` FOREIGN KEY (`item_variant_id`) REFERENCES `item_variants` (`id`) ON UPDATE RESTRICT ON DELETE RESTRICT
)
COLLATE='utf8mb4_unicode_ci'
ENGINE=InnoDB
;
