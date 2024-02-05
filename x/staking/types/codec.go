package types

import (
	"github.com/fibonacci-chain/fbc-social/libs/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types for codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgCreateValidator{}, "fbchain/staking/MsgCreateValidator", nil)
	cdc.RegisterConcrete(MsgEditValidator{}, "fbchain/staking/MsgEditValidator", nil)
	cdc.RegisterConcrete(MsgEditValidatorCommissionRate{}, "fbchain/staking/MsgEditValidatorCommissionRate", nil)
	cdc.RegisterConcrete(MsgDestroyValidator{}, "fbchain/staking/MsgDestroyValidator", nil)
	cdc.RegisterConcrete(MsgDeposit{}, "fbchain/staking/MsgDeposit", nil)
	cdc.RegisterConcrete(MsgWithdraw{}, "fbchain/staking/MsgWithdraw", nil)
	cdc.RegisterConcrete(MsgAddShares{}, "fbchain/staking/MsgAddShares", nil)
	cdc.RegisterConcrete(MsgRegProxy{}, "fbchain/staking/MsgRegProxy", nil)
	cdc.RegisterConcrete(MsgBindProxy{}, "fbchain/staking/MsgBindProxy", nil)
	cdc.RegisterConcrete(MsgUnbindProxy{}, "fbchain/staking/MsgUnbindProxy", nil)
	cdc.RegisterConcrete(CM45Validator{}, "cosmos-sdk/staking/validator", nil)
}

// ModuleCdc is generic sealed codec to be used throughout this module
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
