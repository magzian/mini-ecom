package main

import (
	"backend/controllers"
	"backend/models"
	"backend/routes"
	"backend/utils"
	"context"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Connect to database
	db, err := utils.ConnectDatabase()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Run migrations
	db.AutoMigrate(&models.User{}, &models.Permissions{}, &models.Product{})

	// Create controller with the database instance directly (not the global variable)
	userController := controllers.NewUserController(db, context.Background(), utils.RedisClient)

	app := fiber.New()

	// Configure CORS for Vue.js frontend
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization",
		AllowCredentials: true,
	}))

	// Database middleware
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", db)
		return c.Next()
	})

	routes.Setup(app, userController)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Println("Server starting on :" + port)
	app.Listen(":" + port)
}
