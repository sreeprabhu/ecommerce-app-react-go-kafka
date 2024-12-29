package api

import (
	config "go-react-ecommerce-app/configs"
	"go-react-ecommerce-app/internal/api/rest"
	"go-react-ecommerce-app/internal/api/rest/handlers"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func StartServer(config config.AppConfig) {
	app := fiber.New()

	app.Get("/health", HealthCheck)

	rh := &rest.RestHandler{
		App: app,
	}

	setupRoutes(rh)

	log.Printf("Listening to %v", config.ServerPort)
	app.Listen(config.ServerPort)
}

func HealthCheck(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Health check successful",
	})
}

func setupRoutes(rh *rest.RestHandler) {
	// user handler
	handlers.SetupUserRoutes(rh)

	// transactions handler

	// catalog handler
}
