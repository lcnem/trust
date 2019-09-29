package trust

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewHandler returns a handler for "coin" type messages.
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case MsgEvaluate:
			return handleMsgEvaluate(ctx, keeper, msg)
		case MsgDistributeTokenByScore:
			return handleMsgDistributeTokenByScore(ctx, keeper, msg)
		case MsgDistributeTokenByEvaluation:
			return handleMsgDistributeTokenByEvaluation(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized coin Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleMsgEvaluate(ctx sdk.Context, keeper Keeper, msg MsgEvaluate) sdk.Result {
	err := keeper.SetEvaluation(ctx, msg.TopicID, msg.FromAddress, msg.ToAddress, msg.Weight1000)
	if err != nil {
		return sdk.ResultFromError(err)
	}

	return sdk.Result{}
}

func handleMsgDistributeTokenByScore(ctx sdk.Context, keeper Keeper, msg MsgDistributeTokenByScore) sdk.Result {
	err := keeper.DistributeTokenByScore(ctx, msg.TopicID, msg.FromAddress, msg.Amount)
	if err != nil {
		return sdk.ResultFromError(err)
	}

	return sdk.Result{}
}

func handleMsgDistributeTokenByEvaluation(ctx sdk.Context, keeper Keeper, msg MsgDistributeTokenByEvaluation) sdk.Result {
	err := keeper.DistributeTokenByEvaluation(ctx, msg.TopicID, msg.Address, msg.FromAddress, msg.Amount)
	if err != nil {
		return sdk.ResultFromError(err)
	}

	return sdk.Result{}
}
