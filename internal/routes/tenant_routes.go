package routes

import (
	"umkm-odod/internal/constants"
	"umkm-odod/internal/handler"
	"umkm-odod/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterTenantRoutes(rg *gin.RouterGroup, h *handler.TenantHandler) {
	// endpoint tenant
	rg.GET("/tenants", middleware.AuthRole(constants.RoleSuperAdmin), h.GetTenants) // Tuliskan 1 / lebih jenis role yang boleh akses endpoint ini (pakai middleware AuthRole)
	rg.GET("/tenants/:id", middleware.AuthRole(constants.RoleSuperAdmin), h.GetTenantByID)
	rg.GET("/my-tenant", middleware.AuthRole(constants.RoleSuperAdmin, constants.RoleAdmin, constants.RoleOwner), h.GetMyTenant) // untuk akses data tenant sendiri berdasarkan tenantID yang dibawa jwt setelah user berhasil login
	rg.POST("/tenants", middleware.AuthRole(constants.RoleSuperAdmin), h.CreateTenant)
	rg.PUT("/tenants/:id", middleware.AuthRole(constants.RoleSuperAdmin), h.UpdateTenant)
	rg.DELETE("/tenants/:id", middleware.AuthRole(constants.RoleSuperAdmin), h.DeleteTenant)
}
