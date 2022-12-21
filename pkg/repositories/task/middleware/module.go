package middleware

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

var Module = fx.Provide(NewRTaskProvider)

type RMwDependencies struct {
	fx.In

	Postgres *pgxpool.Pool
	Logger   *logrus.Logger
}

type RMwProvider struct {
	Postgres *pgxpool.Pool
	Logger   *logrus.Logger
}

func NewRTaskProvider(params RMwDependencies) *RMwProvider {
	return &RMwProvider{
		Postgres: params.Postgres,
		Logger:   params.Logger,
	}
}
