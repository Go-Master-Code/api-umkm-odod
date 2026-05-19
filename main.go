package main

import (
	"umkm-odod/internal/database"
	"umkm-odod/internal/handler"
	"umkm-odod/internal/repository"
	"umkm-odod/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// inisialisasi db
	database.InitDB()

	// instance gin engine
	r := gin.New()

	// dependency injection
	tenantRepo := repository.NewTenantRepository(database.DB)
	tenantService := service.NewTenantService(tenantRepo)
	tenantHandler := handler.NewTenantHandler(tenantService)

	// endpoint tenant
	r.GET("/api/tenants", tenantHandler.GetTenants)
	r.GET("/api/tenants/:id", tenantHandler.GetTenantByID)
	r.GET("/api/tenants?name", tenantHandler.GetTenants)
	r.POST("/api/tenants", tenantHandler.CreateTenant)
	r.PUT("/api/tenants/:id", tenantHandler.UpdateTenant)
	r.DELETE("/api/tenants/:id", tenantHandler.DeleteTenant)

	// dependency injection roles
	roleRepo := repository.NewRoleRepository(database.DB)
	roleService := service.NewRoleService(roleRepo)
	roleHandler := handler.NewRoleHandler(roleService)

	// endpoint roles
	r.GET("/api/roles", roleHandler.GetRoles)
	r.GET("/api/roles/:id", roleHandler.GetRoleByID)
	r.POST("/api/roles/", roleHandler.CreateRole)
	r.PUT("/api/roles/:id", roleHandler.UpdateRole)

	// run server
	r.Run("localhost:8080")
}
