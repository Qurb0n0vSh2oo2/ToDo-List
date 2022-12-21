package handler

import (
	iCreateTask "topro/internal/task/create"
	iDeleteTask "topro/internal/task/delete"
	iMiddleware "topro/internal/task/middleware"
	iReadTask "topro/internal/task/read"
	iUpdateTask "topro/internal/task/update"

	iAuthLogin "topro/internal/auth/login"
	iAuthRegister "topro/internal/auth/register"
	iAuthDelete "topro/internal/auth/delete"


	"go.uber.org/fx"
)

var Module = fx.Provide(NewHandlerProvider)

type Handler struct {
	ICreateTask *iCreateTask.ITaskProvider
	IDeleteTask *iDeleteTask.ITaskProvider
	IUpdateTask *iUpdateTask.ITaskProvider
	IReadTask   *iReadTask.ITaskProvider
	IMW         *iMiddleware.IMwProvider

	IAuthRegister *iAuthRegister.IAuthRegProvider
	IAuthLogin    *iAuthLogin.IAuthLoginProvider
	IAuthDelete    *iAuthDelete.IAuthDeleteProvider

}

func NewHandlerProvider(
	createTask *iCreateTask.ITaskProvider,
	deleteTask *iDeleteTask.ITaskProvider,
	updateTask *iUpdateTask.ITaskProvider,
	readTask *iReadTask.ITaskProvider,
	iMw *iMiddleware.IMwProvider,

	authRegist *iAuthRegister.IAuthRegProvider,
	authLogin *iAuthLogin.IAuthLoginProvider,
	authDelete *iAuthDelete.IAuthDeleteProvider,


) *Handler {
	return &Handler{
		ICreateTask: createTask,
		IDeleteTask: deleteTask,
		IUpdateTask: updateTask,
		IReadTask:   readTask,
		IMW:         iMw,

		IAuthRegister: authRegist,
		IAuthLogin:    authLogin,
		IAuthDelete: authDelete,

	}
}
