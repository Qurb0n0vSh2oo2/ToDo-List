package deleteTask

import (
	rDeleteTask "topro/pkg/repositories/task/delete"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"

	"go.uber.org/fx"
)

var Module = fx.Provide(NewITaskProvider)

type ITaskDependencies struct {
	fx.In
	Logger      *logrus.Logger
	RTaskDelete *rDeleteTask.RTaskDeleteProvider
	Postgres    *pgxpool.Pool
}

type ITaskProvider struct {
	Logger      *logrus.Logger
	RTaskDelete *rDeleteTask.RTaskDeleteProvider
	Postgres    *pgxpool.Pool
}

func NewITaskProvider(param ITaskDependencies) *ITaskProvider {
	return &ITaskProvider{
		Logger:      param.Logger,
		RTaskDelete: param.RTaskDelete,
		Postgres:    param.Postgres,
	}
}
