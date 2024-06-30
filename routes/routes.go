package routes

import (
	"github.com/gofiber/fiber/v2"
	"api_fiber/controllers"
)

func Setup(app *fiber.App) {
	app.Post("/cashiers/:cashierId/login", controllers.Login)
	app.Get("/cashiers/:cashierId/logout", controllers.Logout)
	app.Post("/cashiers/:cashierId/password", controllers.Passcode)

	//Cashier routes
	app.Get("/cashiers", controllers.CashiersList)
	app.Get("/cashiers/:cashierId", controllers.GetCashierDetails)
	app.Post("/cashiers", controllers.CreateCashier)
	app.Delete("/cashiers/:cashierId", controllers.DeleteCashier)
	app.Put("/cashiers/:cashierId", controllers.UpdateCashier)
}