package keeper

import (
	"testing"
	"time"

	"github.com/okex/okchain/x/common"

	"github.com/tendermint/go-amino"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/okex/okchain/x/dex/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
)

func TestQuerier_ProductsAndDepositsAndMatchOrder(t *testing.T) {

	testInput := CreateTestInputWithBalance(t, 1, 10000)
	ctx := testInput.Ctx
	addr, err := sdk.AccAddressFromBech32(types.TestTokenPairOwner)
	require.Nil(t, err)
	tokenPair0 := &types.TokenPair{
		BaseAssetSymbol:  "bToken0",
		QuoteAssetSymbol: common.NativeToken,
		Owner:            addr,
		Deposits:         sdk.NewDecCoin(sdk.DefaultBondDenom, sdk.NewInt(50)),
		BlockHeight:      8,
	}

	tokenPair1 := &types.TokenPair{
		BaseAssetSymbol:  "bToken1",
		QuoteAssetSymbol: common.NativeToken,
		Owner:            addr,
		Deposits:         sdk.NewDecCoin(sdk.DefaultBondDenom, sdk.NewInt(100)),
		BlockHeight:      10,
	}

	// SaveTokenPair
	err = testInput.DexKeeper.SaveTokenPair(ctx, tokenPair0)
	require.Nil(t, err)
	err = testInput.DexKeeper.SaveTokenPair(ctx, tokenPair1)
	require.Nil(t, err)
	querier := NewQuerier(testInput.DexKeeper)

	var normalPath = []string{types.QueryProducts, types.QueryDeposits, types.QueryMatchOrder}

	for _, path := range normalPath {
		// successful case
		queryParams, err := types.NewQueryDexInfoParams(types.TestTokenPairOwner, 1, 50)
		require.Nil(t, err)
		bz, err := amino.MarshalJSON(queryParams)
		require.Nil(t, err)
		data, err := querier(ctx, []string{path}, abci.RequestQuery{Data: bz})
		require.Nil(t, err)
		require.True(t, data != nil)

		// error case : failed to query data because  param is nil
		dataUnmarshalJSON, err := querier(ctx, []string{path}, abci.RequestQuery{Data: nil})
		require.Error(t, err)
		require.True(t, dataUnmarshalJSON == nil)

		// successful case : query data while page limit is out range of data amount
		queryParams, err = types.NewQueryDexInfoParams(types.TestTokenPairOwner, 2, 50)
		require.Nil(t, err)
		bz, err = amino.MarshalJSON(queryParams)
		require.Nil(t, err)
		data, err = querier(ctx, []string{path}, abci.RequestQuery{Data: bz})
		require.Nil(t, err)
		require.True(t, data != nil)

		// successful case : query data while  page limit is in range of data amount
		queryParams, err = types.NewQueryDexInfoParams(types.TestTokenPairOwner, 1, 1)
		require.Nil(t, err)
		bz, err = amino.MarshalJSON(queryParams)
		require.Nil(t, err)
		data, err = querier(ctx, []string{path}, abci.RequestQuery{Data: bz})
		require.Nil(t, err)
		require.True(t, data != nil)

	}

	// error case : failed to query because query path is wrong
	queryParams, err := types.NewQueryDexInfoParams(types.TestTokenPairOwner, 1, 50)
	require.Nil(t, err)
	bz, err := amino.MarshalJSON(queryParams)
	require.Nil(t, err)
	dataOther, err := querier(ctx, []string{"others"}, abci.RequestQuery{Data: bz})
	require.NotNil(t, err)
	require.True(t, dataOther == nil)

}

func TestQuerier_QueryParams(t *testing.T) {

	testInput := CreateTestInputWithBalance(t, 1, 10000)
	ctx := testInput.Ctx

	testInput.DexKeeper.SetParams(ctx, types.Params{DelistMaxDepositPeriod: time.Second * 123})
	querier := NewQuerier(testInput.DexKeeper)
	res, err := querier(ctx, []string{"params"}, abci.RequestQuery{})

	require.True(t, len(res) > 0)
	require.True(t, err == nil)

}

func TestQuerier_QueryProductsDelisting(t *testing.T) {

	testInput := CreateTestInputWithBalance(t, 1, 10000)
	ctx := testInput.Ctx

	querier := NewQuerier(testInput.DexKeeper)

	tokenPair := GetBuiltInTokenPair()
	tokenPair.Delisting = true

	err := testInput.DexKeeper.SaveTokenPair(ctx, tokenPair)
	require.Nil(t, err)

	res, err := querier(ctx, []string{types.QueryProductsDelisting}, abci.RequestQuery{})
	require.True(t, len(res) > 0)
	require.Nil(t, err)

}
