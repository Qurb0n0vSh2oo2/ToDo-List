package rreadtasks

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

var Module = fx.Provide(NewRTaskProvider)

type RTaskReadDependencies struct {
	fx.In

	Postgres *pgxpool.Pool
	Logger   *logrus.Logger
}

type RTaskReadProvider struct {
	Postgres *pgxpool.Pool
	Logger   *logrus.Logger
}

func NewRTaskProvider(params RTaskReadDependencies) *RTaskReadProvider {
	return &RTaskReadProvider{
		Postgres: params.Postgres,
		Logger:   params.Logger,
	}
}
