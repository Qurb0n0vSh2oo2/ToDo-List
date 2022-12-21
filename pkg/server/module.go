package server

import (
	// "fmt"
	"fmt"
	http_ "net/http"
	"topro/pkg/config"
	"topro/pkg/http"

	"go.uber.org/fx"
)

var Module = fx.Provide(Server)

type ServerDependencies struct {
	fx.In

	Router *http.HttpRouter
	Config *config.Config
}

type ServerProvide struct {
	Router *http.HttpRouter
	Config *config.Config
}

func Server(params ServerDependencies) *http_.Server {
	data := &http_.Server{
		Addr:    fmt.Sprintf("%s:%s", params.Config.Host, params.Config.Port),
		Handler: params.Router,
	}
	return data
}
