package packets

import (
	"github.com/kooksee/sp2p"
	"github.com/kooksee/scache/types"
)

/*
gossip 分发存储
 */

type kv struct {
	K []byte `json:"k,omitempty"`
	V []byte `json:"v,omitempty"`
}

type KVSetReq struct{ kv }

func (t *KVSetReq) T() byte        { return KVSetReqT }
func (t *KVSetReq) String() string { return KVSetReqS }
func (t *KVSetReq) OnHandle(p sp2p.ISP2P, msg *sp2p.KMsg) error {
	// 存储内容
	if err := kvDb.Set(t.K, t.V); err != nil {
		return types.ErrWithMsg("kvset error", err)
	}

	// 分发到其他的节点
	p.Broadcast(&sp2p.KMsg{Data: msg.Data})
	return nil
}
