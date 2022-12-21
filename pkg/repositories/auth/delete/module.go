package rAuthDelete

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

var Module = fx.Provide(NewRAuthProvider)

type RAuthDeleteDependencies struct {
	fx.In

	Postgres *pgxpool.Pool
	Logger   *logrus.Logger
}

type RAuthDeleteProvider struct {
	Postgres *pgxpool.Pool
	Logger   *logrus.Logger
}

func NewRAuthProvider(params RAuthDeleteDependencies) *RAuthDeleteProvider {
	return &RAuthDeleteProvider{
		Postgres: params.Postgres,
		Logger:   params.Logger,
	}
}
