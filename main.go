package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"ulinan/config"
	hCart "ulinan/domain/cart/handler"
	hUser "ulinan/domain/user/handler"
	rUser "ulinan/domain/user/repository"
	sUser "ulinan/domain/user/service"
	"ulinan/helper/database"
	"ulinan/helper/generator"
	"ulinan/helper/hashing"
	jwt2 "ulinan/helper/jwt"
	"ulinan/helper/payment"
	"ulinan/middleware"
	"ulinan/routes"

	hAuth "ulinan/domain/auth/handler"
	rAuth "ulinan/domain/auth/repository"
	sAuth "ulinan/domain/auth/service"

	hCategory "ulinan/domain/category/handler"
	rCategory "ulinan/domain/category/repository"
	sCategory "ulinan/domain/category/service"

	hProduct "ulinan/domain/product/handler"
	rProduct "ulinan/domain/product/repository"
	sProduct "ulinan/domain/product/service"

	rCart "ulinan/domain/cart/repository"
	sCart "ulinan/domain/cart/service"
	hOrder "ulinan/domain/order/handler"
	rOrder "ulinan/domain/order/repository"
	sOrder "ulinan/domain/order/service"

	hTransaction "ulinan/domain/transaction/handler"
	rTransaction "ulinan/domain/transaction/repository"
	sTransaction "ulinan/domain/transaction/service"

	hDashboard "ulinan/domain/dashboard/handler"
	rDashboard "ulinan/domain/dashboard/repository"
	sDashboard "ulinan/domain/dashboard/service"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:       "Welcome to API Ulinan",
		CaseSensitive: false,
	})

	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin, Content-Type, Accept, Content-Length, Accept-Language, Accept-Encoding, Connection, Authorization",
		AllowOrigins:     "*",
		AllowCredentials: false,
		AllowMethods:     "GET, POST, HEAD, PUT, DELETE, PATCH, OPTIONS",
	}))

	var bootConfig = config.BootConfig()

	db := database.BootDatabase(*bootConfig)
	database.MigrateTable(db)
	hash := hashing.NewHash()
	jwt := jwt2.NewJWT(bootConfig.Secret)
	coreApi := payment.InitSnapMidtrans(*bootConfig)
	generatorID := generator.NewGeneratorUUID(db)

	userRepo := rUser.NewUserRepository(db)
	userService := sUser.NewUserService(userRepo, hash)
	userHandler := hUser.NewUserHandler(userService)

	authRepo := rAuth.NewAuthRepository(db)
	authService := sAuth.NewAuthService(authRepo, userService, hash, jwt)
	authHandler := hAuth.NewAuthHandler(authService)

	categoryRepo := rCategory.NewCategoryRepository(db)
	categoryService := sCategory.NewCategoryService(categoryRepo)
	categoryHandler := hCategory.NewCategoryHandler(categoryService)

	productRepo := rProduct.NewProductRepository(db)
	productService := sProduct.NewProductService(productRepo, categoryService)
	productHandler := hProduct.NewProductHandler(productService)

	cartRepo := rCart.NewCartRepository(db)
	cartService := sCart.NewCartService(cartRepo, productService)
	cartHandler := hCart.NewCartHandler(cartService)

	orderRepo := rOrder.NewOrderRepository(db, coreApi)
	orderService := sOrder.NewOrderService(orderRepo, generatorID, productService, userService, cartService)
	orderHandler := hOrder.NewOrderHandler(orderService)

	transactionRepo := rTransaction.NewTransactionRepository(db)
	transactionService := sTransaction.NewTransactionService(transactionRepo)
	transactionHandler := hTransaction.NewTransactionHandler(transactionService)

	dashboardRepo := rDashboard.NewDashboardRepository(db)
	dashboardService := sDashboard.NewDashboardService(dashboardRepo)
	dashboardHandler := hDashboard.NewDashboardHandler(dashboardService)

	app.Use(middleware.Logging())
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Hello API Ulinan 🚀",
		})
	})

	routes.BootAuthRoute(app, authHandler)
	routes.BootUserRoute(app, userHandler, jwt, userService)
	routes.BootCategoryRoute(app, categoryHandler, jwt, userService)
	routes.BootProductRouter(app, productHandler, jwt, userService)
	routes.BootCartRouter(app, cartHandler, jwt, userService)
	routes.BootOrderRouter(app, orderHandler, jwt, userService)
	routes.BootTransactionRouter(app, transactionHandler, jwt, userService)
	routes.BootDashboardRouter(app, dashboardHandler, jwt, userService)

	addr := fmt.Sprintf(":%d", bootConfig.AppPort)
	if err := app.Listen(addr).Error(); err != addr {
		panic("Appilaction failed to start")
	}
}
