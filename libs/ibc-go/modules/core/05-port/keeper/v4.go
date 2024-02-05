package keeper

import sdk "github.com/fibonacci-chain/fbc-social/libs/cosmos-sdk/types"

func (k Keeper) IsBound(ctx sdk.Context, portID string) bool {
	return k.isBound(ctx, portID)
}
