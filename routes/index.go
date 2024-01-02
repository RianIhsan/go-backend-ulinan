package routes

import (
	"github.com/gofiber/fiber/v2"
	"ulinan/domain/auth"
	"ulinan/domain/user"
	"ulinan/helper/jwt"
	"ulinan/middleware"
)

func BootAuthRoute(app *fiber.App, handler auth.AuthHandlerInterface) {
	authGroup := app.Group("api/auth")
	authGroup.Post("/register", handler.Register)
	authGroup.Post("/login", handler.Login)
}

func BootUserRoute(app *fiber.App, handler user.UserHandlerInterface, jwtService jwt.IJwt, userService user.UserServiceInterface) {
	userGroup := app.Group("api/user")
	userGroup.Get("/me", middleware.Protected(jwtService, userService), handler.GetCurrentUser)
	userGroup.Patch("me/avatar", middleware.Protected(jwtService, userService), handler.UpdateAvatar)
	userGroup.Patch("me/profile", middleware.Protected(jwtService, userService), handler.UpdateUser)
	userGroup.Patch("me/change-password", middleware.Protected(jwtService, userService), handler.UpdatePassword)

}
