package cmd

import (
	"github.com/urfave/cli"
)

func DaemonCmd() cli.Command {
	return cli.Command{
		Name:    "daemon",
		Aliases: []string{"d"},
		Usage:   "start kfs daemon",
		Flags:   []cli.Flag{},
		Action: func(c *cli.Context) error {
			return nil
		},
	}
}
