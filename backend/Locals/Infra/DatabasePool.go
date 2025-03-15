package Infra

import (
	"database/sql"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	_ "github.com/lib/pq"
	"github.com/lucious/urassets/Locals"
)

func mustParseConfig(dsn string) *pgx.ConnConfig {
	cfg, err := pgx.ParseConfig(dsn)
	if err != nil {
		panic(err)
	}
	return cfg
}

func Connect(conf Locals.ConfigField) (db *sql.DB) {
	db = stdlib.OpenDB(*mustParseConfig(conf.ThisDbConfig))

	if err := db.Ping(); err != nil {
		panic(fmt.Sprintf("Failed to connect: %v", err))
	}

	fmt.Println("Connected to DB successfully!")
	return
}
