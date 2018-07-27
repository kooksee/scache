package packets

import (
	"github.com/kooksee/sp2p"
	"github.com/kooksee/kdb"
	"github.com/kooksee/scache/config"
)

var (
	cfg    *config.Config
	kvDb   kdb.IKHash
	kvPrefix = []byte("kv")
)

func Init() {
	cfg = config.GetCfg()
	kvDb = cfg.GetDb().KHash(kvPrefix)

	sp2p.RegistryHandlers(
		KVSetReq{},
	)
}
