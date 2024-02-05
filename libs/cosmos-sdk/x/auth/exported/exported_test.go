package exported_test

import (
	"testing"

	"github.com/fibonacci-chain/fbc-social/libs/tendermint/crypto/secp256k1"
	"github.com/stretchr/testify/require"

	sdk "github.com/fibonacci-chain/fbc-social/libs/cosmos-sdk/types"
	"github.com/fibonacci-chain/fbc-social/libs/cosmos-sdk/x/auth/exported"
	authtypes "github.com/fibonacci-chain/fbc-social/libs/cosmos-sdk/x/auth/types"
)

func TestGenesisAccountsContains(t *testing.T) {
	pubkey := secp256k1.GenPrivKey().PubKey()
	addr := sdk.AccAddress(pubkey.Address())
	acc := authtypes.NewBaseAccount(addr, nil, secp256k1.GenPrivKey().PubKey(), 0, 0)

	genAccounts := exported.GenesisAccounts{}
	require.False(t, genAccounts.Contains(acc.GetAddress()))

	genAccounts = append(genAccounts, acc)
	require.True(t, genAccounts.Contains(acc.GetAddress()))
}
