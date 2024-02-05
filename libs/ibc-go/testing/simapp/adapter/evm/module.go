package evm

import (
	"encoding/json"
	sdk "github.com/fibonacci-chain/fbc-social/libs/cosmos-sdk/types"
	"github.com/fibonacci-chain/fbc-social/libs/ibc-go/testing/simapp/adapter"
	abci "github.com/fibonacci-chain/fbc-social/libs/tendermint/abci/types"
	"github.com/fibonacci-chain/fbc-social/x/evm"
	"github.com/fibonacci-chain/fbc-social/x/evm/types"
)

type EvmModuleAdapter struct {
	evm.AppModule

	tkeeper *evm.Keeper
	ak      types.AccountKeeper
}

func TNewEvmModuleAdapter(k *evm.Keeper, ak types.AccountKeeper) *EvmModuleAdapter {
	ret := &EvmModuleAdapter{}
	ret.AppModule = evm.NewAppModule(k, ak)
	ret.tkeeper = k
	ret.ak = ak
	return ret
}

func (ea EvmModuleAdapter) DefaultGenesis() json.RawMessage {
	return adapter.ModuleCdc.MustMarshalJSON(types.DefaultGenesisState())
}
func (ea EvmModuleAdapter) InitGenesis(ctx sdk.Context, data json.RawMessage) []abci.ValidatorUpdate {
	return ea.initGenesis(ctx, data)
}

func (ea EvmModuleAdapter) initGenesis(ctx sdk.Context, data json.RawMessage) []abci.ValidatorUpdate {
	var genState types.GenesisState
	// Initialize global index to index in genesis state

	adapter.ModuleCdc.MustUnmarshalJSON(data, &genState)
	genState.Params.EnableCall = true
	genState.Params.MaxGasLimitPerTx = 10000000000000
	evm.InitGenesis(ctx, *ea.tkeeper, ea.ak, genState)

	return []abci.ValidatorUpdate{}
}
