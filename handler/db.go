package handler

import "github.com/jackc/pgx/v5/pgxpool"

type DatabaseController struct {
	Pool *pgxpool.Pool
}
