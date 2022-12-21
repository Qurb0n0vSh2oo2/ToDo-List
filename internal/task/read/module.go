package readTask

import (
	rReadTask "topro/pkg/repositories/task/read"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"

	"go.uber.org/fx"
)

var Module = fx.Provide(NewITaskProvider)

type ITaskDependencies struct {
	fx.In
	Logger    *logrus.Logger
	RTaskRead *rReadTask.RTaskReadProvider
	Postgres  *pgxpool.Pool
}

type ITaskProvider struct {
	Logger    *logrus.Logger
	RTaskRead *rReadTask.RTaskReadProvider
	Postgres  *pgxpool.Pool
}

func NewITaskProvider(param ITaskDependencies) *ITaskProvider {
	return &ITaskProvider{
		Logger:    param.Logger,
		RTaskRead: param.RTaskRead,
		Postgres:  param.Postgres,
	}
}
