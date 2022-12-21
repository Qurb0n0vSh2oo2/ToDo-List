package middleware

import (
	rMW "topro/pkg/repositories/task/middleware"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"

	"go.uber.org/fx"
)

var Module = fx.Provide(NewITaskProvider)

type IMwDependencies struct {
	fx.In
	Logger      *logrus.Logger
	RMiddleware *rMW.RMwProvider
	Postgres    *pgxpool.Pool
}

type IMwProvider struct {
	Logger      *logrus.Logger
	RMiddleware *rMW.RMwProvider
	Postgres    *pgxpool.Pool
}

func NewITaskProvider(param IMwDependencies) *IMwProvider {
	return &IMwProvider{
		Logger:      param.Logger,
		RMiddleware: param.RMiddleware,
		Postgres:    param.Postgres,
	}
}
