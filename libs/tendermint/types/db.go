package types

import dbm "github.com/fibonacci-chain/fbc-social/libs/tm-db"

// DBBackend This is set at compile time.
var DBBackend = string(dbm.GoLevelDBBackend)
