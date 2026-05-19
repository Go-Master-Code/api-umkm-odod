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
func ConvertToDTOCatalogCategorySingle(catalogCategory model.CatalogCategory) dto.CatalogCategoryResponse {
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
func ConvertToDTOCatalogItemSingle(catalogItem model.CatalogItem) dto.CatalogItemResponse {
	var catalogItemDTO dto.CatalogItemResponse
	catalogItemDTO.ID = catalogItem.ID
	catalogItemDTO.TenantID = catalogItem.TenantID
	catalogItemDTO.CategoryID = catalogItem.CategoryID
	catalogItemDTO.Name = catalogItem.Name
	catalogItemDTO.Description = catalogItem.Description
	catalogItemDTO.IsActive = catalogItem.IsActive
	return catalogItemDTO
}

func ConvertToDTOCatalogItemPlural(catalogItem []model.CatalogItem) []dto.CatalogItemResponse {
	var catalogItemDTO []dto.CatalogItemResponse
	for _, c := range catalogItem {
		catalogItemDTO = append(catalogItemDTO, dto.CatalogItemResponse{
			ID:          c.ID,
			TenantID:    c.TenantID,
			CategoryID:  c.CategoryID,
			Name:        c.Name,
			Description: c.Description,
			IsActive:    c.IsActive,
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

func ConvertToDTOUserSingle(user model.User) dto.UserResponse {
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
			Phone:       u.Phone,
			IsActive:    u.IsActive,
			LastLoginAt: u.LastLoginAt,
		})
	}
	return userDTO
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

func ConvertToDTOSaleSingle(sale model.Sale) dto.SaleResponse {
	var saleDTO dto.SaleResponse
	saleDTO.ID = sale.ID
	saleDTO.TenantID = sale.TenantID
	saleDTO.TenantName = sale.Tenant.Name
	saleDTO.InvoiceNumber = sale.InvoiceNumber
	saleDTO.CustomerName = sale.CustomerName
	saleDTO.CashierID = sale.CashierID
	saleDTO.CashierName = sale.Cashier.FullName
	saleDTO.Subtotal = sale.Subtotal
	saleDTO.DiscountAmount = sale.DiscountAmount
	saleDTO.TaxAmount = sale.DiscountAmount
	saleDTO.GrandTotal = sale.GrandTotal
	saleDTO.PaymentMethod = sale.PaymentMethod
	saleDTO.PaymentStatus = sale.PaymentStatus
	saleDTO.Notes = sale.Notes
	return saleDTO
}

func ConvertToDTOSalePlural(sales []model.Sale) []dto.SaleResponse {
	var salesDTO []dto.SaleResponse
	for _, s := range sales {
		salesDTO = append(salesDTO, dto.SaleResponse{
			ID:             s.ID,
			TenantID:       s.TenantID,
			TenantName:     s.Tenant.Name,
			InvoiceNumber:  s.InvoiceNumber,
			CustomerName:   s.CustomerName,
			CashierID:      s.CashierID,
			CashierName:    s.Cashier.FullName,
			Subtotal:       s.Subtotal,
			DiscountAmount: s.DiscountAmount,
			TaxAmount:      s.TaxAmount,
			GrandTotal:     s.GrandTotal,
			PaymentMethod:  s.PaymentMethod,
			PaymentStatus:  s.PaymentStatus,
			Notes:          s.Notes,
		})
	}
	return salesDTO
}
