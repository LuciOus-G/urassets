package main

import (
	"flag"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	Servers "github.com/lucious/urassets/Locals"
	DB "github.com/lucious/urassets/Locals/Infra"
	"os"
)

var Help = `Usage: upay -service <service> [COMMANDS] [ARGS]

Commands:
	runserver		Run the server with the default IP and Port
					- Usage [PROGRAM] runserver -service <service> -host <host> -port <port>

	migrate			Migrate the migrations file at db/migrations
					- Usage [PROGRAM] migrate -op <op>

	makemigrations	Make file migrations at db/migrations
					- Usage [PROGRAM] makemigrations -name <eg: add_table>

`

// @title My API
// @version 1.0
// @description This is a sample API.
// @host localhost:3000
// @BasePath /
func main() {
	var (
		choices = []string{"UserServices", "WalletService"}
		flags   = flag.NewFlagSet("upay", flag.ExitOnError)
		Service = flags.String("service", "", "This is mandatory "+
			fmt.Sprintf("\n choice is %s\n", choices))
		Host = flags.String("host", "127.0.0.1", "Set host default to localhost\n")
		Port = flags.String("port", "5000", "Set port default to 5000\n")
		Op   = flags.String("op", "status", "Operation migrate"+
			"\n choice is (up, status, down)\n")
		Name = flags.String("name", "new_migrate", "Name of Migration file\n")
	)
	flags.Usage = func() {
		fmt.Println(Help)
		flags.PrintDefaults()
	}
	if len(os.Args) == 1 {
		flags.Usage()
		os.Exit(0)
	} else if err := flags.Parse(os.Args[2:]); err != nil {
		log.Fatalf("failed to parse args: %v", err)
		return
	}
	args := os.Args[1:]

	LoadEnvironment, _ := Servers.Config()
	ServersInit := Servers.Migration{
		Config:  LoadEnvironment,
		Service: *Service,
		Host:    *Host,
		Port:    *Port,
	}

	switch args[0] {
	case "runserver":
		ServersInit.RunServer(DB.Connect(LoadEnvironment))
	case "migrate":
		ServersInit.MigrateDatabase(*Op)
	case "makemigrations":
		ServersInit.MakeMigrateDatabase(*Name)
	default:
		fmt.Printf("error: Command %s not found\n", args[0])
		flags.Usage()
		os.Exit(0)
	}
}
