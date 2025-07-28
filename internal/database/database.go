package database

import (
	"context"
	"database/sql"

	"github.com/doug-martin/goqu"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/rms-diego/rinha-backend-2025/internal/config"
)

var Db *goqu.Database

func Init(ctx context.Context) error {
	db, err := sql.Open("pgx", config.Env.DATABASE_URL)
	if err != nil {
		return err
	}

	if err := db.PingContext(ctx); err != nil {
		return err
	}

	Db = goqu.New("postgres", db)

	return nil
}
