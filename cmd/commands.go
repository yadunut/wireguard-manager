package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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
		Action: func(ctx *cli.Context) error {
			c, err := cmd.DB.FindClient(ctx.Args().First())
			if err != nil {
				return err
			}
			reader := bufio.NewReader(os.Stdin)
			fmt.Printf("Are you sure you want to remove this client: %s : %s (y / n) : ", c.Name, c.IP.String())
			text, _ := reader.ReadString('\n')
			text = strings.ToLower(strings.TrimSpace(text))
			if text != "y" && text != "yes" {
				fmt.Println("Not deleting")
				return nil
			}
			fmt.Println("deleting client")
			if err := cmd.DB.DeleteClient(c); err != nil {
				return err
			}
			return nil
		},
	}
}
func (cmd *Command) List() *cli.Command {
	return &cli.Command{
		Name:    "list",
		Aliases: []string{"d"},
		Usage:   "list all clients and their IP address",
		Action: func(ctx *cli.Context) error {
			clients, err := cmd.DB.ListClients()
			if err != nil {
				return err
			}

			for client, _ := range clients {
				fmt.Println(client)
			}

			return nil
		},
	}
}
