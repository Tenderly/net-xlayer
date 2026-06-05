package eth

import (
	"github.com/tenderly/net-xlayer/core/types"
	"github.com/tenderly/net-xlayer/rpc"
)

func (b *EthAPIBackend) HistoricalRPCService() *rpc.Client {
	return b.eth.historicalRPCService
}

func (b *EthAPIBackend) Genesis() *types.Block {
	return b.eth.blockchain.Genesis()
}
