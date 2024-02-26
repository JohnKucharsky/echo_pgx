package db

import (
	"context"
	"github.com/induzo/gocom/database/pginit/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

type ApiConfig struct {
	Pool *pgxpool.Pool
}

func DatabaseConnection(dbAddressString string) *ApiConfig {
	ctx := context.Background()
	pgi, err := pginit.New(dbAddressString)

	if err != nil {
		log.Fatal("Can't connect to db", err.Error())
	}

	pool, err := pgi.ConnPool(ctx)

	if err != nil {
		log.Fatal("Can't connect to db", err.Error())
	}

	return &ApiConfig{Pool: pool}
}
