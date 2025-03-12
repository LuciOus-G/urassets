package Infra

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/lucious/urassets/Locals"
	"os"
)

func Connect(config Locals.ConfigField) (db *sql.DB) {
	db, err := sql.Open("postgres", config.ThisDbConfig)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return
}
