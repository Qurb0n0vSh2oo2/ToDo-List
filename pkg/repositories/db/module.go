package db

import (
	ctx "context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

var ModuleConn = fx.Provide(DbConn)

type RDbDependencies struct {
	fx.In

	Logger *logrus.Logger
}

type RDbProvider struct {
	Logger *logrus.Logger
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "3302"
	dbname   = "topro"
)

func DbConn(conn RDbDependencies) (db *pgxpool.Pool, err error) {
	dbUrl := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err = pgxpool.Connect(ctx.Background(), dbUrl)
	if err != nil {
		panic(err)
	}
	err = db.Ping(ctx.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfuly connected!")
	return
}
