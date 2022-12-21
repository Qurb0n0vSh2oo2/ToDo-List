package rupdatetasks

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

var Module = fx.Provide(NewRTaskProvider)

type RTaskUpdateDependencies struct {
	fx.In

	Postgres *pgxpool.Pool
	Logger   *logrus.Logger
}

type RTaskUpdateProvider struct {
	Postgres *pgxpool.Pool
	Logger   *logrus.Logger
}

func NewRTaskProvider(params RTaskUpdateDependencies) *RTaskUpdateProvider {
	return &RTaskUpdateProvider{
		Postgres: params.Postgres,
		Logger:   params.Logger,
	}
}
