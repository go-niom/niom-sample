
package main

import (
	// load API Docs files (Swagger)

	"github.com/go-niom/niom-sample/server"
	_ "github.com/go-niom/niom-sample/docs"
	"github.com/go-niom/niom-sample/pkg/config"
)

// @title Niom-Sample
// @version 1.0
// @description Niom-Sample Backend REST API
// @in header
// @name Authorization
// @host 127.0.0.1:7000
// @BasePath /api
func main() {

	// setup various configuration for app
	config.LoadAllConfigs(".env")
	server.Serve()
}


