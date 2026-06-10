package helper

import (
	"umkm-odod/internal/dto"
	"umkm-odod/internal/model"
)

// tenants
func ConvertToDTOTenantSingle(tenant *model.Tenant) dto.TenantResponse {
	var tenantDTO dto.TenantResponse
	tenantDTO.ID = tenant.ID
	tenantDTO.Name = tenant.Name
	tenantDTO.Address = tenant.Address
	tenantDTO.Phone = tenant.Phone
	return tenantDTO
}

func ConvertToDTOTenantPlural(tenant []model.Tenant) []dto.TenantResponse {
	var tenantsDTO []dto.TenantResponse
	for _, t := range tenant {
		tenantsDTO = append(tenantsDTO, dto.TenantResponse{
			ID:      t.ID,
			Name:    t.Name,
			Address: t.Address,
			Phone:   t.Phone,
		})
	}
	return tenantsDTO
}

// catalog categories
func ConvertToDTOCatalogCategorySingle(catalogCategory *model.CatalogCategory) dto.CatalogCategoryResponse {
	var ccDTO dto.CatalogCategoryResponse
	ccDTO.ID = catalogCategory.ID
	ccDTO.Name = catalogCategory.Name
	ccDTO.TenantID = catalogCategory.TenantID
	ccDTO.TenantName = catalogCategory.Tenant.Name
	return ccDTO
}

func ConvertToDTOCatalogCategoryPlural(catalogCategory []model.CatalogCategory) []dto.CatalogCategoryResponse {
	var ccDTO []dto.CatalogCategoryResponse
	for _, cc := range catalogCategory {
		ccDTO = append(ccDTO, dto.CatalogCategoryResponse{
			ID:         cc.ID,
			Name:       cc.Name,
			TenantID:   cc.TenantID,
			TenantName: cc.Tenant.Name,
		})
	}
	return ccDTO
}

// item attribute
func ConvertToDTOItemAttributeSingle(itemAttribute model.ItemAttribute) dto.ItemAttributeResponse {
	var itemAttributeDTO dto.ItemAttributeResponse
	itemAttributeDTO.ID = itemAttribute.ID
	itemAttributeDTO.Name = itemAttribute.Name
	itemAttributeDTO.TenantID = itemAttribute.TenantID
	itemAttributeDTO.TenantName = itemAttribute.Tenant.Name
	return itemAttributeDTO
}

func ConvertToDTOItemAttributePlural(itemAttribute []model.ItemAttribute) []dto.ItemAttributeResponse {
	var itemAttributeDTO []dto.ItemAttributeResponse
	for _, i := range itemAttribute {
		itemAttributeDTO = append(itemAttributeDTO, dto.ItemAttributeResponse{
			ID:         i.ID,
			Name:       i.Name,
			TenantID:   i.TenantID,
			TenantName: i.Tenant.Name,
		})
	}
	return itemAttributeDTO
}

// catalog categories
func ConvertToDTOCatalogItemSingle(catalogItem *model.CatalogItem) dto.CatalogItemResponse {
	var catalogItemDTO dto.CatalogItemResponse
	catalogItemDTO.ID = catalogItem.ID
	catalogItemDTO.TenantID = catalogItem.TenantID
	catalogItemDTO.TenantName = catalogItem.Tenant.Name
	catalogItemDTO.CategoryID = catalogItem.CategoryID
	catalogItemDTO.CategoryName = catalogItem.CatalogCategory.Name
	catalogItemDTO.Name = catalogItem.Name
	catalogItemDTO.Description = catalogItem.Description
	catalogItemDTO.IsActive = catalogItem.IsActive
	return catalogItemDTO
}

func ConvertToDTOCatalogItemPlural(catalogItem []model.CatalogItem) []dto.CatalogItemResponse {
	var catalogItemDTO []dto.CatalogItemResponse
	for _, c := range catalogItem {
		catalogItemDTO = append(catalogItemDTO, dto.CatalogItemResponse{
			ID:           c.ID,
			TenantID:     c.TenantID,
			TenantName:   c.Tenant.Name,
			CategoryID:   c.CategoryID,
			CategoryName: c.CatalogCategory.Name,
			Name:         c.Name,
			Description:  c.Description,
			IsActive:     c.IsActive,
		})
	}
	return catalogItemDTO
}

func ConvertToDTORoleSingle(role *model.Role) dto.RoleResponse {
	var roleDTO dto.RoleResponse
	roleDTO.ID = role.ID
	roleDTO.Name = role.Name
	roleDTO.TenantID = role.TenantID
	roleDTO.TenantName = role.Tenant.Name
	return roleDTO
}

func ConvertToDTORolePlural(role []model.Role) []dto.RoleResponse {
	var roleDTO []dto.RoleResponse
	for _, r := range role {
		roleDTO = append(roleDTO, dto.RoleResponse{
			ID:         r.ID,
			TenantID:   r.TenantID,
			TenantName: r.Tenant.Name,
			Name:       r.Name,
		})
	}
	return roleDTO
}

func ConvertToDTOUserSingle(user *model.User) dto.UserResponse {
	var userDTO dto.UserResponse
	userDTO.ID = user.ID
	userDTO.TenantID = user.TenantID
	userDTO.TenantName = user.Tenant.Name
	userDTO.RoleID = user.RoleID
	userDTO.RoleName = user.Role.Name
	userDTO.FullName = user.FullName
	userDTO.Username = user.Username
	userDTO.Phone = user.Phone
	userDTO.IsActive = user.IsActive
	userDTO.LastLoginAt = user.LastLoginAt
	return userDTO
}

func ConvertToDTOUserPlural(user []model.User) []dto.UserResponse {
	var userDTO []dto.UserResponse
	for _, u := range user {
		userDTO = append(userDTO, dto.UserResponse{
			ID:          u.ID,
			TenantID:    u.TenantID,
			TenantName:  u.Tenant.Name,
			RoleID:      u.RoleID,
			RoleName:    u.Role.Name,
			Username:    u.Username,
			FullName:    u.FullName,
			Phone:       u.Phone,
			IsActive:    u.IsActive,
			LastLoginAt: u.LastLoginAt,
		})
	}
	return userDTO
}

func ConvertToDTOItemVariantPlural(item []model.ItemVariant) []dto.ItemVariantResponse {
	var ivDTO []dto.ItemVariantResponse
	for _, iv := range item {
		ivDTO = append(ivDTO, dto.ItemVariantResponse{
			ID:           iv.ID,
			TenantID:     iv.TenantID,
			TenantName:   iv.Tenant.Name,
			ItemID:       iv.ItemID,    // catalog item id
			ItemName:     iv.Item.Name, // catalog item name
			SKU:          iv.SKU,
			Barcode:      iv.Barcode,
			VariantName:  iv.VariantName,
			CostPrice:    iv.CostPrice,
			SellingPrice: iv.SellingPrice,
			IsActive:     iv.IsActive,
		})
	}
	return ivDTO
}

func ConvertToDTOItemVariantSingle(item *model.ItemVariant) dto.ItemVariantResponse {
	ivDTO := dto.ItemVariantResponse{
		ID:           item.ItemID,
		TenantID:     item.TenantID,
		TenantName:   item.Tenant.Name,
		ItemID:       item.ItemID,
		ItemName:     item.Item.Name,
		SKU:          item.SKU,
		Barcode:      item.Barcode,
		VariantName:  item.VariantName,
		CostPrice:    item.CostPrice,
		SellingPrice: item.SellingPrice,
		IsActive:     item.IsActive,
	}
	return ivDTO
}

func ConvertToDTOItemVariantAttributeValueSingle(item model.ItemVariantAttributeValue) dto.ItemVariantAttributeValueResponse {
	var itemDTO dto.ItemVariantAttributeValueResponse
	itemDTO.TenantID = item.TenantID
	itemDTO.TenantName = item.Tenant.Name
	itemDTO.AttributeValueID = item.AttributeValueID
	itemDTO.AttributeValue = item.AttributeValue.Value
	itemDTO.VariantID = item.VariantID
	itemDTO.VariantName = item.Variant.VariantName
	return itemDTO
}

func ConvertToDTOItemVariantAttributeValuePlural(items []model.ItemVariantAttributeValue) []dto.ItemVariantAttributeValueResponse {
	var itemsDTO []dto.ItemVariantAttributeValueResponse
	for _, i := range items {
		itemsDTO = append(itemsDTO, dto.ItemVariantAttributeValueResponse{
			TenantID:         i.TenantID,
			TenantName:       i.Tenant.Name,
			VariantID:        i.VariantID,
			VariantName:      i.Variant.VariantName,
			AttributeValueID: i.AttributeValueID,
			AttributeValue:   i.AttributeValue.Value,
		})
	}
	return itemsDTO
}

func ConvertToDTOSaleSingle(sale *model.Sale) dto.SaleResponse {
	// add sale items dulu untuk disisipkan pada master sale
	var saleItemDTO []dto.SaleItemResponse

	for _, item := range sale.SaleItems {
		saleItemDTO = append(saleItemDTO, dto.SaleItemResponse{
			ID:         item.ID,
			TenantID:   item.TenantID,
			TenantName: item.Tenant.Name,
			SaleID:     item.SaleID,
			// InvoiceNumber:       item.Sale.InvoiceNumber, ga usah pake karena sudah muncul di master sale
			ItemVariantID:       item.ItemVariantID,
			ItemVariantName:     item.ItemVariant.VariantName,
			ItemNameSnapshot:    item.ItemNameSnapshot,
			VariantNameSnapshot: item.VariantNameSnapshot,
			SKUSnapshot:         item.SKUSnapshot,
			Qty:                 item.Qty,
			UnitPrice:           item.UnitPrice,
			DiscountAmount:      item.DiscountAmount,
			Subtotal:            item.Subtotal,
			CreatedAt:           item.CreatedAt,
		})
	}

	// masukkan juga saleItemDTO di atas ke dalam field Items
	saleDTO := dto.SaleResponse{
		ID:             sale.ID,
		TenantID:       sale.TenantID,
		TenantName:     sale.Tenant.Name,
		InvoiceNumber:  sale.InvoiceNumber,
		CustomerName:   sale.CustomerName,
		CashierID:      sale.CashierID,
		CashierName:    sale.Cashier.FullName,
		Subtotal:       sale.Subtotal,
		DiscountAmount: sale.DiscountAmount,
		TaxAmount:      sale.TaxAmount,
		GrandTotal:     sale.GrandTotal,
		PaymentMethod:  sale.PaymentMethod,
		PaymentStatus:  sale.PaymentStatus,
		Notes:          sale.Notes,
		CreatedAt:      sale.CreatedAt,
		Items:          saleItemDTO,
	}

	return saleDTO
}

func ConvertToDTOSalePlural(sales []model.Sale) []dto.SaleResponse {
	var salesDTO []dto.SaleResponse
	for _, s := range sales {
		salesDTO = append(salesDTO, ConvertToDTOSaleSingle(&s)) // method yang lebih cepat untuk for range loop plural
	}
	return salesDTO
}

func ConvertToDTOPurchaseSingle(purchase *model.Purchase) dto.PurchaseResponse {
	// add sale items dulu untuk disisipkan pada master purchase
	var purchaseItemDTO []dto.PurchaseItemResponse

	for _, item := range purchase.PurchaseItems {
		purchaseItemDTO = append(purchaseItemDTO, dto.PurchaseItemResponse{
			ID:                  item.ID,
			TenantID:            item.TenantID,
			TenantName:          item.Tenant.Name,
			PurchaseID:          item.PurchaseID,
			ItemVariantID:       item.ItemVariantID,
			ItemVariantName:     item.ItemVariant.VariantName,
			ItemNameSnapshot:    item.ItemNameSnapshot,
			VariantNameSnapshot: item.VariantNameSnapshot,
			SKUSnapshot:         item.SKUSnapshot,
			Qty:                 item.Qty,
			CostPrice:           item.CostPrice,
			DiscountAmount:      item.DiscountAmount,
			Subtotal:            item.Subtotal,
			CreatedAt:           item.CreatedAt,
		})
	}

	// masukkan juga purchaseItemDTO di atas ke dalam field Items
	purchaseDTO := dto.PurchaseResponse{
		ID:             purchase.ID,
		TenantID:       purchase.TenantID,
		TenantName:     purchase.Tenant.Name,
		InvoiceNumber:  purchase.InvoiceNumber,
		SupplierID:     purchase.SupplierID,
		SupplierName:   purchase.Supplier.Name,
		PurchaseNumber: purchase.PurchaseNumber,
		Subtotal:       purchase.Subtotal,
		DiscountAmount: purchase.DiscountAmount,
		TaxAmount:      purchase.TaxAmount,
		GrandTotal:     purchase.GrandTotal,
		CreatedBy:      purchase.Creator.ID,
		CreatedByName:  purchase.Creator.FullName,
		Notes:          purchase.Notes,
		CreatedAt:      purchase.CreatedAt,
		Items:          purchaseItemDTO,
	}

	return purchaseDTO
}

func ConvertToDTOPurchasePlural(purchases []model.Purchase) []dto.PurchaseResponse {
	var purchasesDTO []dto.PurchaseResponse
	for _, p := range purchases {
		purchasesDTO = append(purchasesDTO, ConvertToDTOPurchaseSingle(&p))
	}
	return purchasesDTO
}

func ConvertToDTOPurchaseReturnSingle(purchaseReturn *model.PurchaseReturn) dto.PurchaseReturnResponse {
	// add detil purchase return items dulu untuk disisipkan pada master purchase return
	var returnItemsDTO []dto.PurchaseReturnItemResponse

	for _, item := range purchaseReturn.Items {
		returnItemsDTO = append(returnItemsDTO, dto.PurchaseReturnItemResponse{
			ID:               item.ID,
			TenantID:         item.TenantID,
			PurchaseReturnID: item.PurchaseReturnID,
			ItemVariantID:    item.ItemVariantID,
			ItemVariantName:  item.ItemVariant.VariantName,
			Qty:              item.Qty,
			Notes:            item.Notes,
			CreatedAt:        item.CreatedAt,
		})
	}

	purchaseReturnDTO := dto.PurchaseReturnResponse{
		ID:            purchaseReturn.PurchaseID,
		TenantID:      purchaseReturn.ID,
		TenantName:    purchaseReturn.Tenant.Name,
		PurchaseID:    purchaseReturn.PurchaseID,
		ReturnNumber:  purchaseReturn.ReturnNumber,
		Reason:        purchaseReturn.Reason,
		Notes:         purchaseReturn.Notes,
		CreatedBy:     purchaseReturn.CreatedBy,
		CreatedByName: purchaseReturn.User.FullName,
		CreatedAt:     purchaseReturn.CreatedAt,
		Items:         returnItemsDTO,
	}
	return purchaseReturnDTO
}

func ConvertToDTOPurchaseReturnPlural(purchaseReturns []model.PurchaseReturn) []dto.PurchaseReturnResponse {
	var purchaseReturnsDTO []dto.PurchaseReturnResponse
	for _, pr := range purchaseReturns {
		purchaseReturnsDTO = append(purchaseReturnsDTO, ConvertToDTOPurchaseReturnSingle(&pr))
	}
	return purchaseReturnsDTO
}

func ConvertToDTOCurrentStock(itemVariantID string, stock float64) dto.CurrentStockResponse {
	currentStock := dto.CurrentStockResponse{
		ItemVariantID: itemVariantID,
		CurrentStock:  stock,
	}
	return currentStock
}

func ConvertToDTOStockMovementSingle(sm *model.StockMovement) dto.StockMovementResponse {
	stockMovementDTO := dto.StockMovementResponse{
		ID:              sm.ID,
		TenantID:        sm.TenantID,
		TenantName:      sm.Tenant.Name,
		ItemVariantID:   sm.ItemVariantID,
		ItemVariantName: sm.ItemVariant.VariantName,
		MovementType:    sm.MovementType,
		Qty:             sm.Qty,
		ReferenceType:   sm.ReferenceType,
		ReferenceID:     sm.ReferenceID,
		Notes:           sm.Notes,
		CreatedBy:       sm.CreatedBy,
		CreatedByName:   sm.CreatedByUser.FullName,
	}
	return stockMovementDTO
}

func ConvertToDTOSupplierPlural(suppliers []model.Supplier) []dto.SupplierResponse {
	var suppliersDTO []dto.SupplierResponse
	for _, supplier := range suppliers {
		suppliersDTO = append(suppliersDTO, dto.SupplierResponse{
			ID:       supplier.ID,
			TenantID: supplier.TenantID,
			Name:     supplier.Name,
			Phone:    supplier.Phone,
			Address:  supplier.Address,
			IsActive: supplier.IsActive,
		})
	}
	return suppliersDTO
}

func ConvertToDTOSupplierSingle(supplier *model.Supplier) dto.SupplierResponse {
	supplierDTO := dto.SupplierResponse{
		ID:       supplier.ID,
		TenantID: supplier.TenantID,
		Name:     supplier.Name,
		Phone:    supplier.Phone,
		Address:  supplier.Address,
		IsActive: supplier.IsActive,
	}
	return supplierDTO
}
