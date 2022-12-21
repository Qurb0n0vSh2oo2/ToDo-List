package internal

import (
	createTask "topro/internal/task/create"
	deleteTask "topro/internal/task/delete"
	middleware "topro/internal/task/middleware"
	readTask "topro/internal/task/read"
	updateTask "topro/internal/task/update"

	
	authRegister "topro/internal/auth/register"
	authLogin "topro/internal/auth/login"
	authDelete "topro/internal/auth/delete"



	"go.uber.org/fx"
)

var Module = fx.Options(
	createTask.Module,
	updateTask.Module,
	deleteTask.Module,
	readTask.Module,

	authRegister.Module,
	authLogin.Module,
	authDelete.Module,
	

	middleware.Module,

)
