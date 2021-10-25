package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mfurkan-altun/fiber-todo/controller" // replace
)

func TodoRoute(route fiber.Router) {
	route.Get("", controllers.GetTodos)
}