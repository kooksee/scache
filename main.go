package main

import (
	"github.com/kooksee/scache/config"
	"github.com/kooksee/scache/cmd"
)

func main() {
	cfg := config.DefaultConfig()
	cfg.InitLog()
	cfg.InitDb()

	cmd.Init()
	cmd.RunCmd()
}
