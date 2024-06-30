package main

import (
	db "api_fiber/config"
	routes "api_fiber/routes"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	fmt.Printf("Start API")
	db.Connect()
	app := fiber.New()
	app.Use(cors.New())

	routes.Setup(app)
	app.Listen(":3001")
}
