package app

import (
	"time"

	"github.com/spf13/viper"

	"github.com/fibonacci-chain/fbc-social/libs/cosmos-sdk/codec"
	sdk "github.com/fibonacci-chain/fbc-social/libs/cosmos-sdk/types"
	authtypes "github.com/fibonacci-chain/fbc-social/libs/cosmos-sdk/x/auth"
	authexported "github.com/fibonacci-chain/fbc-social/libs/cosmos-sdk/x/auth/exported"
	abci "github.com/fibonacci-chain/fbc-social/libs/tendermint/abci/types"
	abcitypes "github.com/fibonacci-chain/fbc-social/libs/tendermint/abci/types"
	"github.com/fibonacci-chain/fbc-social/libs/tendermint/libs/log"
	"github.com/fibonacci-chain/fbc-social/libs/tendermint/types"
	dbm "github.com/fibonacci-chain/fbc-social/libs/tm-db"
)

type Option func(option *SetupOption)

type SetupOption struct {
	chainId string
}

func WithChainId(chainId string) Option {
	return func(option *SetupOption) {
		option.chainId = chainId
	}
}

// Setup initializes a new FBChainApp. A Nop logger is set in FBChainApp.
func Setup(isCheckTx bool, options ...Option) *FBChainApp {
	viper.Set(sdk.FlagDBBackend, string(dbm.MemDBBackend))
	types.DBBackend = string(dbm.MemDBBackend)
	db := dbm.NewMemDB()
	app := NewFBChainApp(log.NewNopLogger(), db, nil, true, map[int64]bool{}, 0)

	if !isCheckTx {
		setupOption := &SetupOption{chainId: ""}
		for _, opt := range options {
			opt(setupOption)
		}
		// init chain must be called to stop deliverState from being nil
		genesisState := NewDefaultGenesisState()
		stateBytes, err := codec.MarshalJSONIndent(app.Codec(), genesisState)
		if err != nil {
			panic(err)
		}

		// Initialize the chain
		app.InitChain(
			abci.RequestInitChain{
				Validators:    []abci.ValidatorUpdate{},
				AppStateBytes: stateBytes,
				ChainId:       setupOption.chainId,
			},
		)
	}

	return app
}

func SetupWithGenesisAccounts(isCheckTx bool, genAccs []authexported.GenesisAccount, options ...Option) *FBChainApp {
	viper.Set(sdk.FlagDBBackend, string(dbm.MemDBBackend))
	types.DBBackend = string(dbm.MemDBBackend)
	db := dbm.NewMemDB()
	app := NewFBChainApp(log.NewNopLogger(), db, nil, true, map[int64]bool{}, 0)

	if !isCheckTx {
		setupOption := &SetupOption{chainId: ""}
		for _, opt := range options {
			opt(setupOption)
		}
		// init chain must be called to stop deliverState from being nil
		genesisState := NewDefaultGenesisState()
		authGenesis := authtypes.NewGenesisState(authtypes.DefaultParams(), genAccs)
		genesisState[authtypes.ModuleName] = app.Codec().MustMarshalJSON(authGenesis)
		stateBytes, err := codec.MarshalJSONIndent(app.Codec(), genesisState)
		if err != nil {
			panic(err)
		}

		// Initialize the chain
		testTime, _ := time.Parse("2006-01-02 15:04:05", "2017-04-11 13:33:37")
		app.InitChain(
			abci.RequestInitChain{
				Validators:    []abci.ValidatorUpdate{},
				AppStateBytes: stateBytes,
				ChainId:       setupOption.chainId,
				Time:          testTime,
			},
		)

		app.Commit(abcitypes.RequestCommit{})
		app.BeginBlock(abci.RequestBeginBlock{Header: abcitypes.Header{Height: app.LastBlockHeight() + 1}})
	}

	return app
}
