package main

import (
	"boilerplate-api/config"
	"boilerplate-api/database"
	"boilerplate-api/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    config.LoadConfig()
	database.Connect2()
    
    routes.SetupRoutes(app)

    log.Fatal(app.Listen(":3000"))
}