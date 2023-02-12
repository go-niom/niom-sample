package app

import (
	"github.com/gofiber/fiber/v2"
)

func AppRouter(router fiber.Router) {

	route := router.Group("/app")
	route.Get("/:id", AppController.RetrieveApp)
	route.Get("/", AppController.RetrieveAllApp)
	route.Post("/", AppController.CreateApp)
	route.Patch("/", AppController.UpdateApp)
	route.Delete("/:id", AppController.DeleteApp)
}
