package feesplit

import (
	"github.com/fibonacci-chain/fbc-social/x/feesplit/keeper"
	"github.com/fibonacci-chain/fbc-social/x/feesplit/types"
)

const (
	ModuleName = types.ModuleName
	StoreKey   = types.StoreKey
	RouterKey  = types.RouterKey
)

var (
	NewKeeper           = keeper.NewKeeper
	SetParamsNeedUpdate = types.SetParamsNeedUpdate
)

type (
	Keeper = keeper.Keeper
)
