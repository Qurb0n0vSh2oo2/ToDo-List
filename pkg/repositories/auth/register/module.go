package rregister

import (

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"

	"go.uber.org/fx"
)

var Module = fx.Provide(AuthRegProvider)

type RAuthRegDependencies struct {
	fx.In
	Logger    *logrus.Logger
	Postgres  *pgxpool.Pool
}

type RAuthRegProvider struct {
	Logger    *logrus.Logger
	Postgres  *pgxpool.Pool
}

func AuthRegProvider(param RAuthRegDependencies) *RAuthRegProvider {
	return &RAuthRegProvider{
		Logger:    param.Logger,
		Postgres:  param.Postgres,
	}
}
