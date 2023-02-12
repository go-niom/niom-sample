package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/go-niom/niom-sample/pkg/middleware"
)

//register global Middleware
func registerMiddleware(fApp *fiber.App) {
	middleware.FiberMiddleware(fApp)
}
