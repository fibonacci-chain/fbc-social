package keeper

import (
	sdk "github.com/fibonacci-chain/fbc-social/libs/cosmos-sdk/types"
	"github.com/fibonacci-chain/fbc-social/x/farm/types"
)

// SetParams sets the farm parameters to the param space.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramSubspace.SetParamSet(ctx, &params)
}

// GetParams returns the total set of farm parameters.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramSubspace.GetParamSet(ctx, &params)
	return
}
