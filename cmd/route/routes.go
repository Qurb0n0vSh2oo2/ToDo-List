package route

import (
	"context"
	"net/http"

	handler "topro/cmd/handler"
	"topro/pkg/config"
	mux "topro/pkg/http"

	"go.uber.org/fx"
)

var Module = fx.Provide(Start)

func Start(handl *handler.Handler) (router *mux.HttpRouter) {
	router = mux.NewRouter()
	router.POST("/creating", handl.CreateTaskHandler, handl.PreCheck)
	router.DELETE("/deleting", handl.DeleteTaskHandler, handl.PreCheck)
	router.PUT("/updating", handl.UpdateTaskHandler, handl.PreCheck)
	router.GET("/reading", handl.ReadTaskHandler, handl.PreCheck)

	router.POST("/creatinguser", handl.AccountRegister)
	router.POST("/loginuser", handl.AccountLogin)
	router.DELETE("/deletinguser", handl.DeleteUserHandler)


	return router
}

var Lifecycle = fx.Invoke(RegisterHooks)

func RegisterHooks(lifecycle fx.Lifecycle, server *http.Server, config *config.Config) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go server.ListenAndServe()
			return nil
		},
		// OnStop: func(ctx context.Context) error {
		// 	return mux.Shutdown(ctx)
		// },
	},
	)
}
