package main

import (
	"fmt"
	handler "topro/cmd/handler"
	"topro/cmd/route"
	"topro/internal"
	"topro/pkg/config"
	"topro/pkg/logger"
	"topro/pkg/repositories"
	"topro/pkg/server"

	"go.uber.org/fx"
)

func main() {
	fmt.Println("Listening...")
	fmt.Println(config.LoadConfig().Host)
	fmt.Println(config.LoadConfig().Port)

	s := fx.New(
		server.Module,
		handler.Module,
		repositories.Module,
		internal.Module,
		logger.Module,
		route.Lifecycle,

		config.Module,
		route.Module,
	)
	
	s.Run()
	

}
