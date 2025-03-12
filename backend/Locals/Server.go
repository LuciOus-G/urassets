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
	if c.Service == "UserServices" {
		US.RunUserServiceApp(c.Host, c.Port, DB)
	}
}

func (c *Migration) MigrateDatabase(Op string) {
	cmd := exec.Command("goose", "--table", "public.goose_migrations", "-dir",
		fmt.Sprintf("%s/infra/db/migrations", c.Service), "postgres", c.Config.ThisDbConfig, Op)

	Out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("fail %s", err)
	}
	fmt.Print(string(Out))
}

func (c *Migration) MakeMigrateDatabase(Name string) {
	cmd := exec.Command("goose", "-dir", fmt.Sprintf("%s/infra/db/migrations", c.Service),
		"create", Name, "sql")
	Out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("fail %s", err)
	}
	fmt.Print(string(Out))
}
