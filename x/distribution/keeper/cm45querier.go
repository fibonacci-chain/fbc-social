package keeper

import (
	"github.com/fibonacci-chain/fbc-social/libs/cosmos-sdk/codec"
	sdk "github.com/fibonacci-chain/fbc-social/libs/cosmos-sdk/types"
	abci "github.com/fibonacci-chain/fbc-social/libs/tendermint/abci/types"
	comm "github.com/fibonacci-chain/fbc-social/x/common"
	"github.com/fibonacci-chain/fbc-social/x/distribution/types"
)

func cm45QueryValidatorCommission(ctx sdk.Context, _ []string, req abci.RequestQuery, k Keeper) ([]byte, error) {
	var params types.QueryValidatorCommissionRequest
	err := k.cdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, comm.ErrUnMarshalJSONFailed(err.Error())
	}

	res, err := k.ValidatorCommission(sdk.WrapSDKContext(ctx), &params)
	if err != nil {
		return nil, err
	}
	bz, err := codec.MarshalJSONIndent(k.cdc, res)
	if err != nil {
		return nil, comm.ErrMarshalJSONFailed(err.Error())
	}

	return bz, nil
}
