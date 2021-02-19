package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
	// this line is used by starport scaffolding # 1
	cdc.RegisterConcrete(MsgCreateVote{}, "cosmos/CreateVote", nil)
	cdc.RegisterConcrete(MsgSetVote{}, "cosmos/SetVote", nil)
	cdc.RegisterConcrete(MsgDeleteVote{}, "cosmos/DeleteVote", nil)
	cdc.RegisterConcrete(MsgCreatePoll{}, "cosmos/CreatePoll", nil)
	cdc.RegisterConcrete(MsgSetPoll{}, "cosmos/SetPoll", nil)
	cdc.RegisterConcrete(MsgDeletePoll{}, "cosmos/DeletePoll", nil)
}

// ModuleCdc defines the module codec
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
