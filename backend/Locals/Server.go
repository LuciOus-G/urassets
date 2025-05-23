package Locals

import (
	"database/sql"
	"fmt"
	US "github.com/lucious/urassets/UrAssetsCore"
	"log"
	"os/exec"
)

type Migration struct {
	Config  ConfigField
	Service string
	Host    string
	Port    string
	Schema  string
}

func (c *Migration) RunServer(DB *sql.DB) {
	US.UrAssetsCore(c.Host, c.Port, DB)
}

func (c *Migration) MigrateDatabase(Op string) {
	fmt.Println("Running migrations...")
	cmd := exec.Command("goose", "--table", "public.goose_migrations", "-dir",
		fmt.Sprint("UrAssetsCore/migrations"), "postgres", c.Config.ThisDbConfig, Op)

	Out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("fail %s", err)
	}
	fmt.Print(string(Out))

	fmt.Println("Updating the model ... ")
	cmd = exec.Command("sqlboiler", "psql", "--output", "UrAssetsCore/core/models", "--wipe")
	Out, err = cmd.CombinedOutput()
	if err != nil {
		log.Printf("fail %s", err)
	}
	fmt.Println("Complete")
}

func (c *Migration) MakeMigrateDatabase(Name string) {
	cmd := exec.Command("goose", "-dir", fmt.Sprint("UrAssetsCore/migrations"),
		"create", Name, "sql")
	Out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("fail %s", err)
	}
	fmt.Print(string(Out))
}
