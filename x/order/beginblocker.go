package order

import (
	sdk "github.com/fibonacci-chain/fbc-social/libs/cosmos-sdk/types"

	"github.com/fibonacci-chain/fbc-social/x/common/perf"
	"github.com/fibonacci-chain/fbc-social/x/order/keeper"
	"github.com/fibonacci-chain/fbc-social/x/order/types"
	//"github.com/fibonacci-chain/fbc-social/x/common/version"
)

// BeginBlocker runs the logic of BeginBlocker with version 0.
// BeginBlocker resets keeper cache.
func BeginBlocker(ctx sdk.Context, keeper keeper.Keeper) {
	seq := perf.GetPerf().OnBeginBlockEnter(ctx, types.ModuleName)
	defer perf.GetPerf().OnBeginBlockExit(ctx, types.ModuleName, seq)

	keeper.ResetCache(ctx)
}
