package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	recoverer "github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/joho/godotenv"
	"github.com/mcctrix/ctrix-social-go-backend/db"
	"github.com/mcctrix/ctrix-social-go-backend/utils"
)

func main() {

	loadEnvironment()

	port := os.Getenv("PORT")

	if port == "" {
		port = "4000"
	}
	mainRouter := makeRouter()

	mainRouter.Get("/", func(c fiber.Ctx) error {
		return c.SendString("This is backend of Ctrix Social App!")
	})

	err := mainRouter.Listen(":" + port)
	if err != nil {
		fmt.Println(err)
	}
}

func loadEnvironment() {
	// Load the .env file in the current directory
	godotenv.Load()

	if _, err := os.Stat("./ecdsa_private_key.pem"); err == nil {
		// Pem File exist so do nothing
	} else {
		utils.GenerateEcdsaPrivateKey()
	}
	// db.ResetDB()
	db.CreateInitialDBStructure()
}

func makeRouter() *fiber.App {
	router := fiber.New(fiber.Config{
		TrustProxy: true,
		TrustProxyConfig: fiber.TrustProxyConfig{
			Proxies:  []string{"127.0.0.1", "0.0.0.0"},
			Loopback: true,
		},
	})
	router.Use(recoverer.New())

	router.Use(logger.New(logger.Config{
		Format: "[${ip}]: ${port} ${status} - ${method} ${path}\n",
	}))
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

	router.Use(corsMiddleware)
	return router
}
