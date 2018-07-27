package cmd

import (
	"github.com/urfave/cli"
	"github.com/kooksee/sp2p"
	"net"
	"fmt"
)

func NodeCmd() cli.Command {
	return cli.Command{
		Name:    "create-node",
		Aliases: []string{"cn"},
		Usage:   "create node",
		Flags:   []cli.Flag{},
		Action: func(c *cli.Context) error {
			_, err := fmt.Println(sp2p.CreateNode(sp2p.GenNodeID(), net.ParseIP("127.0.0.1"), uint16(1234)))
			return err
		},
	}
}
