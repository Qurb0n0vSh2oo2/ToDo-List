package updateTask

import (
	rUpdateTask "topro/pkg/repositories/task/update"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"

	"go.uber.org/fx"
)

var Module = fx.Provide(NewITaskProvider)

type ITaskDependencies struct {
	fx.In
	Logger      *logrus.Logger
	RTaskUpdate *rUpdateTask.RTaskUpdateProvider
	Postgres    *pgxpool.Pool
}

type ITaskProvider struct {
	Logger      *logrus.Logger
	RTaskUpdate *rUpdateTask.RTaskUpdateProvider
	Postgres    *pgxpool.Pool
}

func NewITaskProvider(param ITaskDependencies) *ITaskProvider {
	return &ITaskProvider{
		Logger:      param.Logger,
		RTaskUpdate: param.RTaskUpdate,
		Postgres:    param.Postgres,
	}
}
