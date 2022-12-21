package middleware

import (

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"

	"go.uber.org/fx"
)

var Module = fx.Provide(MwProvider)

type IMwDependencies struct {
	fx.In
	Logger    *logrus.Logger
	// RTaskRead *rReadTask.RTaskReadProvider
	Postgres  *pgxpool.Pool
}

type IMwProvider struct {
	Logger    *logrus.Logger
	// RTaskRead *rReadTask.RTaskReadProvider
	Postgres  *pgxpool.Pool
}

func MwProvider(param IMwDependencies) *IMwProvider {
	return &IMwProvider{
		Logger:    param.Logger,
		// RTaskRead: param.RTaskRead,
		Postgres:  param.Postgres,
	}
}
