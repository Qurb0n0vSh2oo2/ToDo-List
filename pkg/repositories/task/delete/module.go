package rdeletetasks

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

var Module = fx.Provide(NewRTaskProvider)

type RTaskDeleteDependencies struct {
	fx.In

	Postgres *pgxpool.Pool
	Logger   *logrus.Logger
}

type RTaskDeleteProvider struct {
	Postgres *pgxpool.Pool
	Logger   *logrus.Logger
}

func NewRTaskProvider(params RTaskDeleteDependencies) *RTaskDeleteProvider {
	return &RTaskDeleteProvider{
		Postgres: params.Postgres,
		Logger:   params.Logger,
	}
}
