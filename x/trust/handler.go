package trust

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lcnem/trust/x/trust/internal/types"
)

// NewHandler creates an sdk.Handler for all the trust type messages
func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case MsgEvaluate:
			return handleMsgEvaluate(ctx, k, msg)
		case MsgDistributeTokenByScore:
			return handleMsgDistributeByScore(ctx, k, msg)
		case MsgDistributeTokenByEvaluation:
			return handleMsgDistributeByEvaluation(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName,  msg)
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

// handdeEvaluate does x
func handleMsgEvaluate(ctx sdk.Context, msg MsgType, k Keeper) sdk.Result {

	err := k.Evaluate(ctx, msg.ValidatorAddr)
	if err != nil {
		return err.Result()
	}

	// TODO: Define your msg events
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.ValidatorAddr.String()),
		),
	)

	return sdk.Result{Events: ctx.EventManager().Events()}
}


// handdeDistributeByScore does x
func handleMsgDistributeByScore(ctx sdk.Context, msg MsgType, k Keeper) sdk.Result {

	err := k.DistributeByScore(ctx, msg.ValidatorAddr)
	if err != nil {
		return err.Result()
	}

	// TODO: Define your msg events
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.ValidatorAddr.String()),
		),
	)

	return sdk.Result{Events: ctx.EventManager().Events()}
}


// handdeDistributeByEvaluation does x
func handleMsgDistributeByEvaluation(ctx sdk.Context, msg MsgType, k Keeper) sdk.Result {

	err := k.DistributeByEvaluation(ctx, msg.ValidatorAddr)
	if err != nil {
		return err.Result()
	}

	// TODO: Define your msg events
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.ValidatorAddr.String()),
		),
	)

	return sdk.Result{Events: ctx.EventManager().Events()}
}
