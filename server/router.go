package server

import (
	"github.com/gofiber/fiber/v2"

	"github.com/go-niom/niom-sample/pkg/common"
	"github.com/go-niom/niom-sample/src/app"
)

func registerRouters(a fiber.Router) {

	common.GeneralRoute(a)
	common.SwaggerRoute(a)
	
	//api router path
	api := a.Group("/api")
	v1 := api.Group("/v1")

	//Add Custom routers
	app.AppRouter(v1)
}
