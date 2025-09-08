package routes

import (
	"backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App, userController *controllers.UserController) {

	app.Post("/api/signup", userController.CreateUser)
	app.Post("/api/add-permission", userController.AddPermission)
	app.Post("/api/login", userController.Login)

	//product routes
	app.Post("/api/products", controllers.CreateProduct)
	app.Get("/api/product/:id", controllers.ViewProduct)
	app.Get("/api/products/", controllers.ViewAllProducts)
	app.Put("/api/products/:id", controllers.UpdateProduct)
	app.Delete("/api/products/:id", controllers.DeleteProduct)

}
