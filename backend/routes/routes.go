package routes

import (
	"backend/controllers"

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

	app.Post("/api/login", controllers.Login)
	app.Post("/api/register", controllers.Register)

	//product routes
	app.Post("/api/products", controllers.CreateProduct)
	app.Get("/api/product/:id", controllers.ViewProduct)
	app.Get("/api/products/", controllers.ViewAllProducts)
	app.Put("/api/products/:id", controllers.UpdateProduct)
	app.Delete("/api/products/:id", controllers.DeleteProduct)

}
