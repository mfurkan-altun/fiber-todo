package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mfurkan-altun/fiber-todo/controller" // replace
)

func TodoRoute(route fiber.Router) {
	route.Get("", controllers.GetTodos)
	route.Post("", controllers.CreateTodo) // new
	route.Put("/:id", controllers.UpdateTodo) // new
	route.Delete("/:id", controllers.DeleteTodo) // new
	route.Get("/:id", controllers.GetTodo) // new
}