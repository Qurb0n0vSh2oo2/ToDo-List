package rcreatetasks

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

var Module = fx.Provide(NewRTaskProvider)

type RTaskCreateDependencies struct {
	fx.In

	Postgres *pgxpool.Pool
	Logger   *logrus.Logger
}

type RTaskCreateProvider struct {
	Postgres *pgxpool.Pool
	Logger   *logrus.Logger
}

func NewRTaskProvider(params RTaskCreateDependencies) *RTaskCreateProvider {
	return &RTaskCreateProvider{
		Postgres: params.Postgres,
		Logger:   params.Logger,
	}
}
