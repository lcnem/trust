package trust

import (
	"github.com/lcnem/trust/x/trust/internal/keeper"
	"github.com/lcnem/trust/x/trust/internal/types"
)

const (
	// TODO: define constants that you would like exposed from the internal package

	ModuleName        = types.ModuleName
	RouterKey         = types.RouterKey
	StoreKey          = types.StoreKey
	DefaultParamspace = types.DefaultParamspace
	QueryParams       = types.QueryParams
	QuerierRoute      = types.QuerierRoute
)

var (
	// functions aliases
	NewKeeper           = keeper.NewKeeper
	NewQuerier          = keeper.NewQuerier
	RegisterCodec       = types.RegisterCodec
	NewGenesisState     = types.NewGenesisState
	DefaultGenesisState = types.DefaultGenesisState
	ValidateGenesis     = types.ValidateGenesis
	// TODO: Fill out function aliases
	NewMsgEvaluate                    = types.NewMsgEvaluate
	NewMsgDistributeTokenByScore      = types.NewMsgDistributeTokenByScore
	NewMsgDistributeTokenByEvaluation = types.NewMsgDistributeTokenByEvaluation

	// variable aliases
	ModuleCdc     = types.ModuleCdc
	// TODO: Fill out variable aliases
)

type (
	Keeper       = keeper.Keeper
	GenesisState = types.GenesisState
	Params       = types.Params

	// TODO: Fill out module types
	MsgEvaluate = types.MsgEvaluate
	MsgDistributeTokenByScore = types.MsgDistributeTokenByScore
	MsgDistributeTokenByEvaluation = types.MsgDistributeTokenByEvaluation
)
