package cmd

import (
	"github.com/inconshreveable/log15"
	"github.com/urfave/cli"
	"github.com/kooksee/scache/config"
)

var (
	logger log15.Logger
	cfg    *config.Config
)

func Init() {
	cfg = config.GetCfg()
	logger = cfg.GetLog().New("package", "cmd")
}

func isDevflag() cli.BoolFlag { return cli.BoolFlag{Name: "debug", Destination: &cfg.IsDev, Usage: "debug mode"} }
