package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// ModuleCdc is the codec for the module
var ModuleCdc = codec.New()

func init() {
	RegisterCodec(ModuleCdc)
}

// RegisterCodec registers concrete types on the Amino codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgEvaluate{}, "trust/MsgEvaluate", nil)
	cdc.RegisterConcrete(MsgDistributeTokenByScore{}, "trust/MsgDistributeTokenByScore", nil)
	cdc.RegisterConcrete(MsgDistributeTokenByEvaluation{}, "trust/MsgDistributeTokenByEvaluation", nil)
}
