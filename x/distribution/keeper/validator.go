package keeper

import (
	sdk "github.com/fibonacci-chain/fbc-social/libs/cosmos-sdk/types"

	"github.com/fibonacci-chain/fbc-social/x/distribution/types"
	"github.com/fibonacci-chain/fbc-social/x/staking/exported"
)

// initialize rewards for a new validator
func (k Keeper) initializeValidator(ctx sdk.Context, val exported.ValidatorI) {
	if k.CheckDistributionProposalValid(ctx) {
		k.initializeValidatorDistrProposal(ctx, val)
		return
	}

	// set accumulated commissions
	k.SetValidatorAccumulatedCommission(ctx, val.GetOperator(), types.InitialValidatorAccumulatedCommission())
}
