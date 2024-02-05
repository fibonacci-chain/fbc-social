package gov

import (
	"github.com/fibonacci-chain/fbc-social/libs/cosmos-sdk/client/context"
	"github.com/fibonacci-chain/fbc-social/libs/cosmos-sdk/codec"
	"github.com/fibonacci-chain/fbc-social/libs/cosmos-sdk/types/module"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"

	"github.com/fibonacci-chain/fbc-social/x/gov/types"

	anytypes "github.com/fibonacci-chain/fbc-social/libs/cosmos-sdk/codec/types"
	GovCli "github.com/fibonacci-chain/fbc-social/x/gov/client/cli"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
)

var (
	_ module.AppModuleBasicAdapter = AppModuleBasic{}
)

func (a AppModuleBasic) RegisterInterfaces(registry anytypes.InterfaceRegistry) {
}

func (a AppModuleBasic) RegisterGRPCGatewayRoutes(cliContext context.CLIContext, serveMux *runtime.ServeMux) {
}

func (a AppModuleBasic) GetTxCmdV2(cdc *codec.CodecProxy, reg anytypes.InterfaceRegistry) *cobra.Command {
	proposalCLIHandlers := make([]*cobra.Command, len(a.proposalHandlers))
	for i, proposalHandler := range a.proposalHandlers {
		proposalCLIHandlers[i] = proposalHandler.CLIHandler(cdc, reg)
	}

	return GovCli.GetTxCmd(types.StoreKey, cdc.GetCdc(), proposalCLIHandlers)
}

func (a AppModuleBasic) GetQueryCmdV2(cdc *codec.CodecProxy, reg anytypes.InterfaceRegistry) *cobra.Command {
	return nil
}

func (a AppModuleBasic) RegisterRouterForGRPC(cliCtx context.CLIContext, r *mux.Router) {}
