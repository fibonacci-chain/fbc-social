package ibc_tx

import (
	"fmt"
	ibctx "github.com/fibonacci-chain/fbc-social/libs/cosmos-sdk/types/ibc-adapter"
	"github.com/fibonacci-chain/fbc-social/libs/cosmos-sdk/x/auth/types"
	"github.com/gogo/protobuf/proto"
)

func IbcTxEncoder() ibctx.IBCTxEncoder {
	return func(tx ibctx.Tx) ([]byte, error) {
		txWrapper, ok := tx.(*wrapper)
		if !ok {
			return nil, fmt.Errorf("expected %T, got %T", &wrapper{}, tx)
		}

		raw := &types.TxRaw{
			BodyBytes:     txWrapper.getBodyBytes(),
			AuthInfoBytes: txWrapper.getAuthInfoBytes(),
			Signatures:    txWrapper.tx.Signatures,
		}

		return proto.Marshal(raw)
	}
}
