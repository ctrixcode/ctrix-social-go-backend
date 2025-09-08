package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func CorsMiddleware() fiber.Handler {
	corsMethods := strings.Split(strings.Join([]string{
		fiber.MethodGet,
		fiber.MethodPost,
		fiber.MethodHead,
		fiber.MethodPut,
		fiber.MethodDelete,
		fiber.MethodPatch,
		fiber.MethodOptions,
	}, ","), ",")

	corsMiddleware := cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://127.0.0.1:3000", "https://ctrix-social.vercel.app"}, // Allows all origins
		AllowMethods:     corsMethods,
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Set-Cookie"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	})

	return corsMiddleware
}
