package createTask

import (
	rCreateTask "topro/pkg/repositories/task/create"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"

	"go.uber.org/fx"
)

var Module = fx.Provide(NewITaskProvider)

type ITaskDependencies struct {
	fx.In
	Logger      *logrus.Logger
	RTaskCreate *rCreateTask.RTaskCreateProvider
	Postgres    *pgxpool.Pool
}

type ITaskProvider struct {
	Logger      *logrus.Logger
	RTaskCreate *rCreateTask.RTaskCreateProvider
	Postgres    *pgxpool.Pool
}

func NewITaskProvider(param ITaskDependencies) *ITaskProvider {
	return &ITaskProvider{
		Logger:      param.Logger,
		RTaskCreate: param.RTaskCreate,
		Postgres: param.Postgres,
	}
}
