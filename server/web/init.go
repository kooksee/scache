package web

import (
	"github.com/json-iterator/go"
	"github.com/kooksee/kfs/config"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
	cfg  *config.Config
)

func Init() {
	cfg = config.GetCfg()
}
