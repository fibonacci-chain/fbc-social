package proxy

import (
	"github.com/fibonacci-chain/fbc-social/libs/tendermint/crypto/merkle"
)

func defaultProofRuntime() *merkle.ProofRuntime {
	prt := merkle.NewProofRuntime()
	prt.RegisterOpDecoder(
		merkle.ProofOpSimpleValue,
		merkle.SimpleValueOpDecoder,
	)
	return prt
}
