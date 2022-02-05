package routes

import (
	"github.com/baijupadmanabhan/golang-jwt-authentication/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/register", controllers.Register)
	app.Post("/login", controllers.Login)
	app.Post("/user", controllers.User)
	app.Post("/logout", controllers.Logout)
}
