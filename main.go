package main

import (
	log "github.com/sirupsen/logrus"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/urfave/cli/v2"
	"github.com/yadunut/wireguard-manager/cmd"
	"github.com/yadunut/wireguard-manager/lib/db"
)

const (
	DBPath = "foo.db"
)

func main() {
	log.SetOutput(os.Stderr)
	log.SetLevel(log.InfoLevel)
	command := cmd.Command{}
	var err error

	command.DB, err = db.InitDB(DBPath)
	if err != nil {
		log.Fatal(err)
	}
	defer command.DB.Close()

	app := &cli.App{
		Name:  "Wireguard Client Manager",
		Usage: "Easily Manage wireguard clients",
		Commands: []*cli.Command{
			command.Add(),
			command.List(),
			command.Delete(),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
