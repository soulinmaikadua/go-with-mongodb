package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/soulinmaikadua/go-with-mongodb/pkg/controllers"
)

func AuthRoutes(app *fiber.App) {

	route := app.Group("/auth")

	// route.Get("/singup", controllers.GetUsers)
	route.Post("/login", controllers.Login)
}
