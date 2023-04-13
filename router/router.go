package router

import (
	"github.com/MCPutro/golang-todo/controller"
	"github.com/gofiber/fiber/v2"
)

func NewRouter(activityController controller.ActivityController, todoController controller.TodoController) *fiber.App {
	app := fiber.New()

	app.Get("/", activityController.GetAll)

	app.Get("/activity-groups", activityController.GetAll)
	app.Get("/activity-groups/:ActId", activityController.GetOne)
	app.Post("/activity-groups", activityController.Create)
	app.Patch("/activity-groups/:ActId", activityController.Update)
	app.Delete("/activity-groups/:ActId", activityController.Delete)

	app.Get("/todo-items", todoController.GetAll)
	app.Post("/todo-items", todoController.Create)
	app.Get("/todo-items/:TodoId", todoController.GetOne)
	app.Delete("/todo-items/:TodoId", todoController.Delete)
	app.Patch("/todo-items/:TodoId", todoController.Update)

	return app
}
