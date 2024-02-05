package continuousauction

import (
	sdk "github.com/fibonacci-chain/fbc-social/libs/cosmos-sdk/types"

	"github.com/fibonacci-chain/fbc-social/x/order/keeper"
)

// nolint
type CaEngine struct {
}

// nolint
func (e *CaEngine) Run(ctx sdk.Context, keeper keeper.Keeper) {
	// TODO
}
