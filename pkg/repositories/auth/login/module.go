package rlogin

import (

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"

	"go.uber.org/fx"
)

var Module = fx.Provide(AuthLoginProvider)

type RAuthLoginDependencies struct {
	fx.In
	Logger    *logrus.Logger
	Postgres  *pgxpool.Pool
}

type RAuthLoginProvider struct {
	Logger    *logrus.Logger
	Postgres  *pgxpool.Pool
}

func AuthLoginProvider(param RAuthLoginDependencies) *RAuthLoginProvider {
	return &RAuthLoginProvider{
		Logger:    param.Logger,
		Postgres:  param.Postgres,
	}
}
