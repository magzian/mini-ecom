package routes

import (
	"backend/controllers"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

/* func RegisterRoutes(r *mux.Router, db *gorm.DB) {
	r.HandleFunc("/api/register", controllers.Register(db)).Methods("POST")
	r.HandleFunc("/api/login", controllers.Login(db)).Methods("POST")

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/products", controllers.ListProducts(db)).Methods("GET")
	api.HandleFunc("/orders", controllers.CreateOrder(db)).Methods("POST")
	api.Use(middleware.JWTAuth)
} */

func Setup(app *fiber.App) {
	fmt.Println("Inside Setup function - registering routes...") // ADD THIS

	app.Post("/api/test-register", func(c *fiber.Ctx) error {
		fmt.Println("Test register route hit!") // ADD THIS
		return c.JSON(fiber.Map{"message": "Test register route works"})
	})

	app.Post("/api/login", controllers.Login)
	app.Post("/api/register", controllers.Register)

	fmt.Println("Routes registered:")
	fmt.Println("- POST /api/login")
	fmt.Println("- POST /api/register")
	fmt.Println("- POST /api/test-register") // ADD THIS
}
