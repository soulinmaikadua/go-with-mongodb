package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/soulinmaikadua/go-with-mongodb/pkg/controllers"
)

func UserRoutes(app *fiber.App) {

	route := app.Group("/users")

	route.Get("/", controllers.GetUsers)
	route.Get("/:id", controllers.GetUser)
}
