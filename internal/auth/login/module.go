package auth

import (
	
	rLogin "topro/pkg/repositories/auth/login"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"

	"go.uber.org/fx"
)

var Module = fx.Provide(AuthLoginProvider)

type IAuthLoginDependencies struct {
	fx.In
	Logger    *logrus.Logger
	RAuthLogin *rLogin.RAuthLoginProvider
	Postgres  *pgxpool.Pool
}

type IAuthLoginProvider struct {
	Logger    *logrus.Logger
	RAuthLogin *rLogin.RAuthLoginProvider
	Postgres  *pgxpool.Pool
}

func AuthLoginProvider(param IAuthLoginDependencies) *IAuthLoginProvider {
	return &IAuthLoginProvider{
		Logger:    param.Logger,
		Postgres:  param.Postgres,
		RAuthLogin: param.RAuthLogin,
	}
}
