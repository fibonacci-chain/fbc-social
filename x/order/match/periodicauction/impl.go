package periodicauction

import (
	sdk "github.com/fibonacci-chain/fbc-social/libs/cosmos-sdk/types"

	"github.com/fibonacci-chain/fbc-social/x/order/keeper"
)

// PaEngine is the periodic auction match engine
type PaEngine struct {
}

// nolint
func (e *PaEngine) Run(ctx sdk.Context, keeper keeper.Keeper) {
	cleanupExpiredOrders(ctx, keeper)
	cleanupOrdersWhoseTokenPairHaveBeenDelisted(ctx, keeper)
	matchOrders(ctx, keeper)
}
