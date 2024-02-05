package vmbridge

import (
	"github.com/fibonacci-chain/fbc-social/libs/cosmos-sdk/codec"
	"github.com/fibonacci-chain/fbc-social/libs/cosmos-sdk/types/module"
	"github.com/fibonacci-chain/fbc-social/x/vmbridge/keeper"
	"github.com/fibonacci-chain/fbc-social/x/wasm"
)

func RegisterServices(cfg module.Configurator, keeper keeper.Keeper) {
	RegisterMsgServer(cfg.MsgServer(), NewMsgServerImpl(keeper))
}

func GetWasmOpts(cdc *codec.ProtoCodec) wasm.Option {
	return wasm.WithMessageEncoders(RegisterSendToEvmEncoder(cdc))
}
