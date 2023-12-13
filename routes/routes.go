package routes

import (
	"boilerplate/backend/controllers"

	"github.com/VDiPaola/api-network-server/api_fiber"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	//app.Post("/api/user/signup", controllers.SignUp)
	app.Post("/api/healthcheck", api_fiber.NodeHealthCheck)

	app.Get("/api/test", controllers.NodeTest)
	app.Get("/ping", controllers.Ping)
}
