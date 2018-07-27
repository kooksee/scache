package core

import (
	"github.com/inconshreveable/log15"
	"github.com/kooksee/kdb"
	"github.com/kooksee/scache/config"
)

var (
	logger log15.Logger
	cfg    *config.Config
	kvDb   kdb.IKHash

	kvPrefix = []byte("kv")
)

func Init() {
	cfg = config.GetCfg()
	logger = cfg.GetLog().New("package", "api.core")
	kvDb = cfg.GetDb().KHash(kvPrefix)
}
