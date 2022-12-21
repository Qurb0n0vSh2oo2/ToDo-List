package authDelete

import (
	rAuthDelete "topro/pkg/repositories/auth/delete"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"

	"go.uber.org/fx"
)

var Module = fx.Provide(NewIAuthDeleteProvider)

type IAuthDeleteDependencies struct {
	fx.In
	Logger      *logrus.Logger
	RAuthDelete *rAuthDelete.RAuthDeleteProvider
	Postgres    *pgxpool.Pool
}

type IAuthDeleteProvider struct {
	Logger      *logrus.Logger
	RAuthDelete *rAuthDelete.RAuthDeleteProvider
	Postgres    *pgxpool.Pool
}

func NewIAuthDeleteProvider(param IAuthDeleteDependencies) *IAuthDeleteProvider {
	return &IAuthDeleteProvider{
		Logger:      param.Logger,
		RAuthDelete: param.RAuthDelete,
		Postgres:    param.Postgres,
	}
}
