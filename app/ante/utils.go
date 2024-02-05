package ante

import (
	sdk "github.com/fibonacci-chain/fbc-social/libs/cosmos-sdk/types"
	"github.com/fibonacci-chain/fbc-social/libs/cosmos-sdk/x/auth"
	"github.com/fibonacci-chain/fbc-social/libs/cosmos-sdk/x/auth/exported"
)

func getAccount(ak *auth.AccountKeeper, ctx *sdk.Context, addr sdk.AccAddress, accCache auth.Account) (auth.Account, sdk.Gas) {
	gasMeter := ctx.GasMeter()
	var gasUsed sdk.Gas
	if accCache != nil {
		var ok bool
		if ok, gasUsed = exported.TryAddGetAccountGas(gasMeter, ak, accCache); ok {
			return accCache, gasUsed
		}
	}
	return exported.GetAccountAndGas(ctx, ak, addr)
}
