package routes

import (
	"github.com/gofiber/fiber/v2"
	"ulinan/domain/auth"
	"ulinan/domain/category"
	"ulinan/domain/order"
	"ulinan/domain/product"
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

func BootCategoryRoute(app *fiber.App, handler category.CategoryHandlerInterface, jwtService jwt.IJwt, userService user.UserServiceInterface) {
	CategoryGroup := app.Group("api/category")
	CategoryGroup.Post("/", middleware.Protected(jwtService, userService), handler.CreateCategory)
	CategoryGroup.Get("/", handler.GetAllCategories)
	CategoryGroup.Get("/:id", handler.GetCategoryById)
	CategoryGroup.Patch("/:id", middleware.Protected(jwtService, userService), handler.UpdateCategory)
	CategoryGroup.Delete("/:id", middleware.Protected(jwtService, userService), handler.DeleteCategory)
}

func BootProductRouter(app *fiber.App, handler product.ProductHandlerInterface, jwtService jwt.IJwt, userService user.UserServiceInterface) {
	productGroup := app.Group("api/product")
	productGroup.Post("/", middleware.Protected(jwtService, userService), handler.CreateProduct)
	productGroup.Post("/image", middleware.Protected(jwtService, userService), handler.CreateProductImage)
	productGroup.Get("/", handler.GetAllProducts)
	productGroup.Get("/:id", handler.GetProductById)
	productGroup.Patch("/:id", middleware.Protected(jwtService, userService), handler.UpdateProduct)
	productGroup.Delete("/:id", middleware.Protected(jwtService, userService), handler.DeleteProduct)
	productGroup.Delete("/:productId/image/:imageId", middleware.Protected(jwtService, userService), handler.DeleteProductImage)
}

func BootOrderRouter(app *fiber.App, handler order.OrderHandlerInterface, jwtService jwt.IJwt, userService user.UserServiceInterface) {
	orderGroup := app.Group("api/order")
	orderGroup.Post("/", middleware.Protected(jwtService, userService), handler.CreateOrder)
	orderGroup.Post("/callback", handler.Callback)
}
