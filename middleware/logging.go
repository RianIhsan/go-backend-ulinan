package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Logging() fiber.Handler {
	return logger.New(logger.Config{
		Format: "[${time}] [${status}] ${latency} ${method} ${path} - ${ip} ${response}\n",
	})
}
