package trust

import (
	"github.com/lcnem/lcnem-trust/x/trust/internal/keeper"
	"github.com/lcnem/lcnem-trust/x/trust/internal/types"
)

// nolint
const (
	ModuleName = types.ModuleName
	RouterKey  = types.RouterKey
	StoreKey   = types.StoreKey
)

// nolint
var (
	NewKeeper                         = keeper.NewKeeper
	NewQuerier                        = keeper.NewQuerier
	NewMsgEvaluate                    = types.NewMsgEvaluate
	NewMsgDistributeTokenByScore      = types.NewMsgDistributeTokenByScore
	NewMsgDistributeTokenByEvaluation = types.NewMsgDistributeTokenByEvaluation
	ModuleCdc                         = types.ModuleCdc
	RegisterCodec                     = types.RegisterCodec
)

type (
	// Keeper keeper.Keeper
	Keeper = keeper.Keeper
	// MsgEvaluate types.MsgEvaluate
	MsgEvaluate = types.MsgEvaluate
	// MsgDistributeTokenByScore types.MsgDistributeTokenByScore
	MsgDistributeTokenByScore = types.MsgDistributeTokenByScore
	// MsgDistributeTokenByEvaluation types.MsgDistributeTokenByEvaluation
	MsgDistributeTokenByEvaluation = types.MsgDistributeTokenByEvaluation
	// QueryResAccountScores types.QueryResAccountScores
	QueryResAccountScores = types.QueryResAccountScores
)
