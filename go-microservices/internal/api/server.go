package api

import (
	config "go-react-ecommerce-app/configs"
	"go-react-ecommerce-app/internal/api/rest"
	"go-react-ecommerce-app/internal/api/rest/handlers"
	"go-react-ecommerce-app/internal/domain"
	"go-react-ecommerce-app/internal/helper"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartServer(config config.AppConfig) {
	app := fiber.New()

	db, err := gorm.Open(postgres.Open(config.DBConnection), &gorm.Config{})

	if err != nil {
		log.Fatalf("database connection error: #{err}\n")
	}

	log.Println("database connected!")

	// run migration
	db.AutoMigrate(&domain.User{})

	auth := helper.SetupAuth(config.AppSecret)

	// app.Get("/health", HealthCheck)
	rh := &rest.RestHandler{
		App:  app,
		DB:   db,
		Auth: auth,
	}

	setupRoutes(rh)

	log.Printf("Listening to %v", config.ServerPort)
	app.Listen(config.ServerPort)
}

// func HealthCheck(ctx *fiber.Ctx) error {
// 	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
// 		"message": "Health check successful",
// 	})
// }

func setupRoutes(rh *rest.RestHandler) {
	// user handler
	handlers.SetupUserRoutes(rh)

	// transactions handler

	// catalog handler
}
