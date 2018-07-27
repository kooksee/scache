package config

import (
	"os"
	"github.com/inconshreveable/log15"
	"github.com/kooksee/kdb"
	"github.com/kooksee/scache/types"
	"github.com/kooksee/sp2p"
)

func (t *Config) GetDb() kdb.IKDB {
	if t.db == nil {
		panic("please init db")
	}
	return t.db
}

func GetCfg() *Config {
	if cfg == nil {
		panic("please init config")
	}
	return cfg
}

func homeDir(defaultHome string) string {
	if len(os.Args) > 2 && os.Args[len(os.Args)-2] == "--home" {
		defaultHome = os.Args[len(os.Args)-1]
		os.Args = os.Args[:len(os.Args)-2]
	}
	return defaultHome
}

func (t *Config) GetLog() log15.Logger {
	if t.l == nil {
		panic("please init log")
	}
	return t.l
}

func (t *Config) GetSP2P() sp2p.ISP2P {
	if t.p2p == nil {
		panic("please init p2p")
	}
	return t.p2p
}

/*
	Seeds    []string `json:"seeds"`
	NodeUrl  string   `json:"node_url"`
	WebAddr  string   `json:"web_url"`
 */
func DefaultConfig() *Config {
	once.Do(func() {
		cfg = &Config{
			Home:     "kdata",
			IsDev:    true,
			LogLevel: log15.LvlDebug.String(),
			Seeds:    []string{},
			NodeAddr: "sp2p://cc7d13d8573a5a95dea60793c39338e7a98a0be62f2af3c0dd9257b9880d81bb@127.0.0.1:8088",
			WebAddr:  "0.0.0.0:8080",
		}
		cfg.Home = homeDir("kdata")
		cfg.LogLevel = "debug"
		types.MustNotErr(types.EnsureDir(cfg.Home, os.FileMode(0755)))
	})
	return cfg
}
