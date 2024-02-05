package store

import (
	dbm "github.com/fibonacci-chain/fbc-social/libs/tm-db"

	"github.com/fibonacci-chain/fbc-social/libs/cosmos-sdk/store/cache"
	"github.com/fibonacci-chain/fbc-social/libs/cosmos-sdk/store/rootmulti"
	"github.com/fibonacci-chain/fbc-social/libs/cosmos-sdk/store/types"
)

func NewCommitMultiStore(db dbm.DB) types.CommitMultiStore {
	return rootmulti.NewStore(db)
}

func NewCommitKVStoreCacheManager() types.MultiStorePersistentCache {
	return cache.NewCommitKVStoreCacheManager(cache.DefaultCommitKVStoreCacheSize)
}
