package config

import (
	"github.com/inconshreveable/log15"
	"os"
	"github.com/kooksee/kdb"
	"path/filepath"
	"github.com/kooksee/scache/types"
	"sync"
	"github.com/kooksee/sp2p"
)

var (
	cfg  *Config
	once sync.Once
)

type Config struct {
	Home     string   `json:"home"`
	IsDev    bool     `json:"is_dev"`
	LogLevel string   `json:"log_level"`
	Seeds    []string `json:"seeds"`
	NodeAddr  string   `json:"node_addr"`
	WebAddr  string   `json:"web_addr"`

	l   log15.Logger
	db  kdb.IKDB
	p2p sp2p.ISP2P
}

func (t *Config) InitLog() {
	t.l = log15.New()
	ll, err := log15.LvlFromString(t.LogLevel)
	types.MustNotErr(types.ErrWithMsg("config InitLog error", err))
	t.l.SetHandler(log15.LvlFilterHandler(ll, log15.StreamHandler(os.Stdout, log15.TerminalFormat())))
}

func (t *Config) InitConfigFile() {
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

func (t *Config) InitSP2P() {
	//kf := sp2p.DefaultConfig()

	//kf.InitDb(t.GetDb()).InitConn().InitLocalNode().InitP2P()
	//t.p2p = kf.GetP2P()
}
