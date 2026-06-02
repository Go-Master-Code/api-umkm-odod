package main

import (
	"time"
	"umkm-odod/internal/database"
	"umkm-odod/internal/handler"
	"umkm-odod/internal/middleware"
	"umkm-odod/internal/repository"
	"umkm-odod/internal/routes"
	"umkm-odod/internal/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// inisialisasi db
	database.InitDB()

	// instance gin engine
	r := gin.New()

	// tambahkan CORS apabila server backend berbeda dengan frontend
	// ===============================
	// 🔥 CORS CONFIG
	// ===============================
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173", // frontend Vue
		},

		// method yang diizinkan
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
			"OPTIONS",
		},

		// 🔥 HEADER YANG DIIZINKAN
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Authorization", // WAJIB untuk JWT
		},

		// optional
		ExposeHeaders: []string{
			"Content-Length",
		},

		AllowCredentials: true,

		MaxAge: 12 * time.Hour,
	}))
	// go get github.com/gin-contrib/cors

	// dependency injection
	tenantRepo := repository.NewTenantRepository(database.DB)
	tenantService := service.NewTenantService(tenantRepo)
	tenantHandler := handler.NewTenantHandler(tenantService)

	// dependency injection roles
	roleRepo := repository.NewRoleRepository(database.DB)
	roleService := service.NewRoleService(roleRepo)
	roleHandler := handler.NewRoleHandler(roleService)

	// dependency injection user
	userRepo := repository.NewUserRepository(database.DB)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// dependency injection catalog category
	catalogCategoryRepo := repository.NewCatalogCategoryRepository(database.DB)
	catalogCategoryService := service.NewCatalogCategoryService(catalogCategoryRepo)
	catalogCategoryHandler := handler.NewCatalogCategoryHandler(catalogCategoryService)

	// dependency injection catalog item
	catalogItemRepo := repository.NewCatalogItemRepository(database.DB)
	catalogItemService := service.NewCatalogItemService(catalogItemRepo)
	catalogItemHandler := handler.NewCatalogItemHandler(catalogItemService)

	// dependency injection item variant
	itemVariantRepo := repository.NewItemVariantRepository(database.DB)
	stockMovementRepo := repository.NewStockMovementRepository(database.DB)

	itemVariantService := service.NewItemVariantService(itemVariantRepo, stockMovementRepo)
	itemVariantHandler := handler.NewItemVariantHandler(itemVariantService)

	// dependency injection stock movement
	stockMovementService := service.NewStockMovementService(database.DB, stockMovementRepo, itemVariantRepo)
	stockMovementHandler := handler.NewStockMovementHandler(stockMovementService)

	// dependency injection sale item
	saleItemRepo := repository.NewSaleItemRepository(database.DB)
	// tidak ada service dan handler untuk sale item repo karena ini bukan standalone business process
	// dia merupakan bagian dari sale transaction

	// dependency injection sale
	saleRepo := repository.NewSaleRepository(database.DB)
	saleService := service.NewSaleService(database.DB, saleRepo, saleItemRepo, itemVariantRepo, stockMovementRepo)
	saleHandler := handler.NewSaleHandler(saleService)

	// router group public tidak perlu pakai middleware AuthRequired
	public := r.Group("/api")
	// endpoint login
	public.POST("/login", userHandler.Login)

	// authorization yang akan dipasang pada tiap endpoint yang dilindungi (harus punya token)
	authorized := r.Group("/api")

	authorized.Use(middleware.AuthRequired()) // file auth_required.go -> handler ini dieksekusi dulu sebeleum eksekusi handler endpoint
	{
		// list handler role
		routes.RegisterRoleRoutes(authorized, roleHandler)
		// list handler user
		routes.RegisterUserRoutes(authorized, userHandler)
		// list handler tenant
		routes.RegisterTenantRoutes(authorized, tenantHandler)
		// list handler catalog category
		routes.RegisterCatalogCategoryRoutes(authorized, catalogCategoryHandler)
		// list handler catalog item
		routes.RegisterCatalogItemRoutes(authorized, catalogItemHandler)
		// list handler item variant
		routes.RegisterItemVariantRoutes(authorized, itemVariantHandler)
		// list handler stock movement
		routes.RegisterStockMovementRoutes(authorized, stockMovementHandler)
		// list handler sales
		routes.RegisterSaleRoutes(authorized, saleHandler)
	}

	// run server
	r.Run("localhost:8080")
}
