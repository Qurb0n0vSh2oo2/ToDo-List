package auth

import (
	
	rRegist "topro/pkg/repositories/auth/register"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"

	"go.uber.org/fx"
)

var Module = fx.Provide(AuthRegProvider)

type IAuthRegDependencies struct {
	fx.In
	Logger    *logrus.Logger
	RAuthRegist *rRegist.RAuthRegProvider
	Postgres  *pgxpool.Pool
}

type IAuthRegProvider struct {
	Logger    *logrus.Logger
	RAuthRegist *rRegist.RAuthRegProvider
	Postgres  *pgxpool.Pool
}

func AuthRegProvider(param IAuthRegDependencies) *IAuthRegProvider {
	return &IAuthRegProvider{
		Logger:    param.Logger,
		Postgres:  param.Postgres,
		RAuthRegist: param.RAuthRegist,
	}
}
