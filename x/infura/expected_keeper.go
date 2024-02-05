package infura

import evm "github.com/fibonacci-chain/fbc-social/x/evm/watcher"

type EvmKeeper interface {
	SetObserverKeeper(keeper evm.InfuraKeeper)
}
