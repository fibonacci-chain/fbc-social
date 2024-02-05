package simulator

import (
	sdk "github.com/fibonacci-chain/fbc-social/libs/cosmos-sdk/types"
)

type Simulator interface {
	Simulate([]sdk.Msg, sdk.CacheMultiStore) (*sdk.Result, error)
	Context() *sdk.Context
	Release()
}

var NewWasmSimulator func() Simulator
