package repositories

import (
	"topro/pkg/repositories/db"
	createTask "topro/pkg/repositories/task/create"
	deleteTask "topro/pkg/repositories/task/delete"
	middleware "topro/pkg/repositories/task/middleware"
	readTask "topro/pkg/repositories/task/read"
	updateTask "topro/pkg/repositories/task/update"

	authRegist "topro/pkg/repositories/auth/register"
	authLogin "topro/pkg/repositories/auth/login"
	authDelete "topro/pkg/repositories/auth/delete"


	"go.uber.org/fx"
)

var Module = fx.Options(
	createTask.Module,
	deleteTask.Module,
	updateTask.Module,
	readTask.Module,
	middleware.Module,

	authRegist.Module,
	authLogin.Module,
	authDelete.Module,


	db.ModuleConn,
)
