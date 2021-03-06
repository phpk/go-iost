package consensus

import (
	"github.com/iost-official/go-iost/consensus/pob"
	"github.com/iost-official/go-iost/core/blockcache"
	"github.com/iost-official/go-iost/core/global"
	"github.com/iost-official/go-iost/core/txpool"
	"github.com/iost-official/go-iost/p2p"
)

// Type is the type of consensus
type Type uint8

// The types of consensus
const (
	_ Type = iota
	Pob
)

// Consensus is a consensus server.
type Consensus interface {
	Start() error
	Stop()
}

// New returns the different consensus strategy.
func New(cType Type, baseVariable global.BaseVariable, blkcache blockcache.BlockCache, txPool txpool.TxPool, service p2p.Service) Consensus {
	switch cType {
	case Pob:
		return pob.New(baseVariable, blkcache, txPool, service)
	default:
		return pob.New(baseVariable, blkcache, txPool, service)
	}
}
