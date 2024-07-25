package routes

import (
	"boilerplate-api/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
    api := app.Group("/api")

    api.Get("/example", controllers.ExampleController)
}
