package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"ulinan/config"
	hUser "ulinan/domain/user/handler"
	rUser "ulinan/domain/user/repository"
	sUser "ulinan/domain/user/service"
	"ulinan/helper/database"
	"ulinan/helper/hashing"
	jwt2 "ulinan/helper/jwt"
	"ulinan/middleware"
	"ulinan/routes"

	// hUser "ulinan/domain/user/handler"
	hAuth "ulinan/domain/auth/handler"
	rAuth "ulinan/domain/auth/repository"
	sAuth "ulinan/domain/auth/service"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:       "Welcome to API Ulinan",
		CaseSensitive: false,
	})

	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	var bootConfig = config.BootConfig()

	db := database.BootDatabase(*bootConfig)
	database.MigrateTable(db)
	hash := hashing.NewHash()
	jwt := jwt2.NewJWT(bootConfig.Secret)

	userRepo := rUser.NewUserRepository(db)
	userService := sUser.NewUserService(userRepo, hash)
	userHandler := hUser.NewUserHandler(userService)

	authRepo := rAuth.NewAuthRepository(db)
	authService := sAuth.NewAuthService(authRepo, userService, hash, jwt)
	authHandler := hAuth.NewAuthHandler(authService)

	app.Use(middleware.Logging())
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Hello API Ulinan 🚀",
		})
	})

	routes.BootAuthRoute(app, authHandler)
	routes.BootUserRoute(app, userHandler, jwt, userService)
	addr := fmt.Sprintf(":%d", bootConfig.AppPort)
	app.Listen(addr)
}
