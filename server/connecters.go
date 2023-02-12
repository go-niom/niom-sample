package server

import (
	"runtime/debug"

	"github.com/go-niom/niom-sample/pkg/logger"
)

func initConnectors() {
	log := logger.Get()
	buildInfo, _ := debug.ReadBuildInfo()
	log.Info().Str("go_version", buildInfo.GoVersion).Msg("Logger initialized")
}
