package proto

import (
	"os"
	"testing"

	"github.com/fibonacci-chain/fbc-social/libs/cosmos-sdk/store"
	sdk "github.com/fibonacci-chain/fbc-social/libs/cosmos-sdk/types"
	abci "github.com/fibonacci-chain/fbc-social/libs/tendermint/abci/types"
	"github.com/fibonacci-chain/fbc-social/libs/tendermint/libs/log"
	dbm "github.com/fibonacci-chain/fbc-social/libs/tm-db"
	"github.com/stretchr/testify/require"
)

func createTestInput(t *testing.T) (sdk.Context, ProtocolKeeper) {
	keyMain := sdk.NewKVStoreKey("main")

	db := dbm.NewMemDB()
	ms := store.NewCommitMultiStore(db)
	ms.MountStoreWithDB(keyMain, sdk.StoreTypeIAVL, db)

	require.NoError(t, ms.LoadLatestVersion())

	ctx := sdk.NewContext(ms, abci.Header{}, false, log.NewTMLogger(os.Stdout))

	keeper := NewProtocolKeeper(keyMain)

	return ctx, keeper
}
