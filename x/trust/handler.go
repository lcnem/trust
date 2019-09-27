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
		default:
			errMsg := fmt.Sprintf("Unrecognized coin Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

// Handle a message to evaluate
func handleMsgEvaluate(ctx sdk.Context, keeper Keeper, msg MsgEvaluate) sdk.Result {
	keeper.SetEvaluation()

	return sdk.Result{}
}

// Handle a message to distribute token by score
func handleMsgDistributeTokenByScore(ctx sdk.Context, keeper Keeper, msg MsgDistributeTokenByScore) sdk.Result {
	keeper.DistributeTokenByScore()

	return sdk.Result{}
}
