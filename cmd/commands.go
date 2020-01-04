package cmd

import (
	"github.com/urfave/cli/v2"
	"github.com/yadunut/wireguard-manager/lib/clients"
	"github.com/yadunut/wireguard-manager/lib/db"
)

type Command struct {
	DB *db.DB
}

func (cmd *Command) Add() *cli.Command {
	return &cli.Command{
		Name:    "add",
		Aliases: []string{"d"},
		Usage:   "add a client",
		Action: func(ctx *cli.Context) error {
			client, err := clients.NewClient(ctx.Args().Get(0), ctx.Args().Get(1), ctx.Args().Get(2))
			if err != nil {
				return err
			}

			if err := cmd.DB.AddClient(client); err != nil {
				return err
			}
			return nil
		},
	}
}
func (cmd *Command) Delete() *cli.Command {
	return &cli.Command{
		Name:      "del",
		Aliases:   []string{"d"},
		Usage:     "delete a client",
		ArgsUsage: "[name]",
		Action:    func(ctx *cli.Context) error { return nil },
	}
}
func (cmd *Command) List() *cli.Command {
	return &cli.Command{
		Name:    "list",
		Aliases: []string{"d"},
		Usage:   "list all clients and their IP address",
		Action:  func(ctx *cli.Context) error { return nil },
	}
}
