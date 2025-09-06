package main

import (
	"backend/models"
	"backend/routes"
	"backend/utils"
	"fmt"
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
	app.Use(cors.New())

	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", db)
		return c.Next()
	})

	fmt.Println("About to call routes.Setup()")
	routes.Setup(app)
	fmt.Println("routes.Setup() completed")

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Println("Server starting on :" + port)
	app.Listen(":3000")
}
