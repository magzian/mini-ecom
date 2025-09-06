// In your main.go
package main

import (
	"backend/models"
	"backend/routes"
	"backend/utils"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	db, err := utils.ConnectDatabase()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})

	app := fiber.New()

	// Configure CORS for Vue.js frontend
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173", // Vue dev server default port
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization",
		AllowCredentials: true,
	}))

	// Database middleware
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", db)
		return c.Next()
	})

	routes.Setup(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Println("Server starting on :" + port)
	app.Listen(":3000")
}
