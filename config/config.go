package config

import (
	"os"
	"github.com/inconshreveable/log15"
	"github.com/kooksee/kdb"
	"github.com/kooksee/scache/types"
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

func DefaultConfig() *Config {
	once.Do(func() {
		cfg = &Config{}
		cfg.Home = homeDir("kdata")
		cfg.LogLevel = "debug"
		types.MustNotErr(types.EnsureDir(cfg.Home, os.FileMode(0755)))
	})
	return cfg
}
