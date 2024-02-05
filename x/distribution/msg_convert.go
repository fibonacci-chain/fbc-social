package distribution

import (
	"encoding/json"
	"errors"

	"github.com/fibonacci-chain/fbc-social/libs/cosmos-sdk/baseapp"
	sdk "github.com/fibonacci-chain/fbc-social/libs/cosmos-sdk/types"
	tmtypes "github.com/fibonacci-chain/fbc-social/libs/tendermint/types"
	"github.com/fibonacci-chain/fbc-social/x/common"
	"github.com/fibonacci-chain/fbc-social/x/distribution/types"
)

var (
	ErrCheckSignerFail = errors.New("check signer fail")
)

func init() {
	RegisterConvert()
}

func RegisterConvert() {
	enableHeight := tmtypes.GetVenus3Height()
	baseapp.RegisterCmHandle("fbchain/distribution/MsgWithdrawDelegatorAllRewards", baseapp.NewCMHandle(ConvertWithdrawDelegatorAllRewardsMsg, enableHeight))
}

func ConvertWithdrawDelegatorAllRewardsMsg(data []byte, signers []sdk.AccAddress) (sdk.Msg, error) {
	newMsg := types.MsgWithdrawDelegatorAllRewards{}
	err := json.Unmarshal(data, &newMsg)
	if err != nil {
		return nil, err
	}
	err = newMsg.ValidateBasic()
	if err != nil {
		return nil, err
	}
	if ok := common.CheckSignerAddress(signers, newMsg.GetSigners()); !ok {
		return nil, ErrCheckSignerFail
	}
	return newMsg, nil
}
