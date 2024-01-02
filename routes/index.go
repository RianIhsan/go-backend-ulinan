package routes

import (
	"github.com/gofiber/fiber/v2"
	"ulinan/domain/auth"
)

func BootAuthRoute(app *fiber.App, handler auth.AuthHandlerInterface) {
	authGroup := app.Group("api/auth")
	authGroup.Post("/register", handler.Register)
	authGroup.Post("/login", handler.Login)
}
