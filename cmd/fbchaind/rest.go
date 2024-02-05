package main

import (
	"fmt"

	"github.com/spf13/viper"

	"github.com/fibonacci-chain/fbc-social/app"
	"github.com/fibonacci-chain/fbc-social/app/rpc"
	"github.com/fibonacci-chain/fbc-social/app/types"
	"github.com/fibonacci-chain/fbc-social/libs/cosmos-sdk/client"
	"github.com/fibonacci-chain/fbc-social/libs/cosmos-sdk/client/lcd"
	"github.com/fibonacci-chain/fbc-social/libs/cosmos-sdk/server"
	"github.com/fibonacci-chain/fbc-social/libs/cosmos-sdk/types/tx"
	"github.com/fibonacci-chain/fbc-social/libs/cosmos-sdk/x/auth"
	authrest "github.com/fibonacci-chain/fbc-social/libs/cosmos-sdk/x/auth/client/rest"
	bankrest "github.com/fibonacci-chain/fbc-social/libs/cosmos-sdk/x/bank/client/rest"
	mintclient "github.com/fibonacci-chain/fbc-social/libs/cosmos-sdk/x/mint/client"
	mintrest "github.com/fibonacci-chain/fbc-social/libs/cosmos-sdk/x/mint/client/rest"
	supplyrest "github.com/fibonacci-chain/fbc-social/libs/cosmos-sdk/x/supply/client/rest"
	ibctransferrest "github.com/fibonacci-chain/fbc-social/libs/ibc-go/modules/apps/transfer/client/rest"
	ammswaprest "github.com/fibonacci-chain/fbc-social/x/ammswap/client/rest"
	dexclient "github.com/fibonacci-chain/fbc-social/x/dex/client"
	dexrest "github.com/fibonacci-chain/fbc-social/x/dex/client/rest"
	dist "github.com/fibonacci-chain/fbc-social/x/distribution"
	distr "github.com/fibonacci-chain/fbc-social/x/distribution"
	distrest "github.com/fibonacci-chain/fbc-social/x/distribution/client/rest"
	erc20client "github.com/fibonacci-chain/fbc-social/x/erc20/client"
	erc20rest "github.com/fibonacci-chain/fbc-social/x/erc20/client/rest"
	evmclient "github.com/fibonacci-chain/fbc-social/x/evm/client"
	evmrest "github.com/fibonacci-chain/fbc-social/x/evm/client/rest"
	farmclient "github.com/fibonacci-chain/fbc-social/x/farm/client"
	farmrest "github.com/fibonacci-chain/fbc-social/x/farm/client/rest"
	fsrest "github.com/fibonacci-chain/fbc-social/x/feesplit/client/rest"
	govrest "github.com/fibonacci-chain/fbc-social/x/gov/client/rest"
	orderrest "github.com/fibonacci-chain/fbc-social/x/order/client/rest"
	paramsclient "github.com/fibonacci-chain/fbc-social/x/params/client"
	slashingrest "github.com/fibonacci-chain/fbc-social/x/slashing/client/rest"
	stakingrest "github.com/fibonacci-chain/fbc-social/x/staking/client/rest"
	"github.com/fibonacci-chain/fbc-social/x/token"
	tokensrest "github.com/fibonacci-chain/fbc-social/x/token/client/rest"
	wasmrest "github.com/fibonacci-chain/fbc-social/x/wasm/client/rest"
	"github.com/fibonacci-chain/fbc-social/x/wasm/proxy"
)

// registerRoutes registers the routes from the different modules for the LCD.
// NOTE: details on the routes added for each module are in the module documentation
// NOTE: If making updates here you also need to update the test helper in client/lcd/test_helper.go
func registerRoutes(rs *lcd.RestServer) {
	registerGrpc(rs)
	rpc.RegisterRoutes(rs)
	pathPrefix := viper.GetString(server.FlagRestPathPrefix)
	if pathPrefix == "" {
		pathPrefix = types.EthBech32Prefix
	}
	registerRoutesV1(rs, pathPrefix)
	registerRoutesV2(rs, pathPrefix)
	proxy.SetCliContext(rs.CliCtx)
}

func registerGrpc(rs *lcd.RestServer) {
	app.ModuleBasics.RegisterGRPCGatewayRoutes(rs.CliCtx, rs.GRPCGatewayRouter)
	app.ModuleBasics.RegisterRPCRouterForGRPC(rs.CliCtx, rs.Mux)
	tx.RegisterGRPCGatewayRoutes(rs.CliCtx, rs.GRPCGatewayRouter)
}

func registerRoutesV1(rs *lcd.RestServer, pathPrefix string) {
	v1Router := rs.Mux.PathPrefix(fmt.Sprintf("/%s/v1", pathPrefix)).Name("v1").Subrouter()
	client.RegisterRoutes(rs.CliCtx, v1Router)
	authrest.RegisterRoutes(rs.CliCtx, v1Router, auth.StoreKey)
	bankrest.RegisterRoutes(rs.CliCtx, v1Router)
	stakingrest.RegisterRoutes(rs.CliCtx, v1Router)
	slashingrest.RegisterRoutes(rs.CliCtx, v1Router)
	distrest.RegisterRoutes(rs.CliCtx, v1Router, dist.StoreKey)

	orderrest.RegisterRoutes(rs.CliCtx, v1Router)
	tokensrest.RegisterRoutes(rs.CliCtx, v1Router, token.StoreKey)
	dexrest.RegisterRoutes(rs.CliCtx, v1Router)
	ammswaprest.RegisterRoutes(rs.CliCtx, v1Router)
	supplyrest.RegisterRoutes(rs.CliCtx, v1Router)
	farmrest.RegisterRoutes(rs.CliCtx, v1Router)
	evmrest.RegisterRoutes(rs.CliCtx, v1Router)
	erc20rest.RegisterRoutes(rs.CliCtx, v1Router)
	wasmrest.RegisterRoutes(rs.CliCtx, v1Router)
	fsrest.RegisterRoutes(rs.CliCtx, v1Router)
	govrest.RegisterRoutes(rs.CliCtx, v1Router,
		[]govrest.ProposalRESTHandler{
			paramsclient.ProposalHandler.RESTHandler(rs.CliCtx),
			distr.CommunityPoolSpendProposalHandler.RESTHandler(rs.CliCtx),
			distr.ChangeDistributionTypeProposalHandler.RESTHandler(rs.CliCtx),
			distr.WithdrawRewardEnabledProposalHandler.RESTHandler(rs.CliCtx),
			distr.RewardTruncatePrecisionProposalHandler.RESTHandler(rs.CliCtx),
			dexclient.DelistProposalHandler.RESTHandler(rs.CliCtx),
			farmclient.ManageWhiteListProposalHandler.RESTHandler(rs.CliCtx),
			evmclient.ManageContractDeploymentWhitelistProposalHandler.RESTHandler(rs.CliCtx),
			evmclient.ManageSysContractAddressProposalHandler.RESTHandler(rs.CliCtx),
			evmclient.ManageContractByteCodeProposalHandler.RESTHandler(rs.CliCtx),
			mintclient.ManageTreasuresProposalHandler.RESTHandler(rs.CliCtx),
			mintclient.ModifyNextBlockUpdateProposalHandler.RESTHandler(rs.CliCtx),
			erc20client.TokenMappingProposalHandler.RESTHandler(rs.CliCtx),
		},
	)
	mintrest.RegisterRoutes(rs.CliCtx, v1Router)
	ibctransferrest.RegisterOriginRPCRoutersForGRPC(rs.CliCtx, v1Router)
}

func registerRoutesV2(rs *lcd.RestServer, pathPrefix string) {
	v2Router := rs.Mux.PathPrefix(fmt.Sprintf("/%s/v2", pathPrefix)).Name("v1").Subrouter()
	client.RegisterRoutes(rs.CliCtx, v2Router)
	authrest.RegisterRoutes(rs.CliCtx, v2Router, auth.StoreKey)
	bankrest.RegisterRoutes(rs.CliCtx, v2Router)
	stakingrest.RegisterRoutes(rs.CliCtx, v2Router)
	distrest.RegisterRoutes(rs.CliCtx, v2Router, dist.StoreKey)
	orderrest.RegisterRoutesV2(rs.CliCtx, v2Router)
	tokensrest.RegisterRoutesV2(rs.CliCtx, v2Router, token.StoreKey)
	fsrest.RegisterRoutesV2(rs.CliCtx, v2Router)
}
