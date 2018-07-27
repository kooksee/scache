package config

import (
	"github.com/inconshreveable/log15"
	"os"
	"github.com/kooksee/kdb"
	"path/filepath"
	"github.com/kooksee/scache/types"
	"sync"
)

var (
	cfg *Config
	once *sync.Once
)

type Config struct {
	Home     string
	IsDev    bool
	LogLevel string
	Seeds    []string

	l  log15.Logger
	db kdb.IKDB
}

func (t *Config) InitLog() {
	t.l = log15.New()
	ll, err := log15.LvlFromString(t.LogLevel)
	types.MustNotErr(types.ErrWithMsg("config InitLog error", err))
	t.l.SetHandler(log15.LvlFilterHandler(ll, log15.StreamHandler(os.Stdout, log15.TerminalFormat())))
}

func (t *Config) InitDb() {
	kf := kdb.DefaultConfig()
	kf.InitKdb(filepath.Join(t.Home, "db"))
	t.db = kf.GetDb()
}
