package keeper

import (
	"github.com/fibonacci-chain/fbc-social/x/staking/exported"

	sdk "github.com/fibonacci-chain/fbc-social/libs/cosmos-sdk/types"
)

func (k Keeper) Delegation(ctx sdk.Context, delAddr sdk.AccAddress, address2 sdk.ValAddress) exported.DelegatorI {
	delegator, found := k.GetDelegator(ctx, delAddr)
	if !found {
		return nil
	}

	return delegator
}
