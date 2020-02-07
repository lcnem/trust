package keeper

import (
	"fmt"

	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/lcnem/trust/x/trust/internal/types"
)

// NewQuerier creates a new querier for trust clients.
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case types.QueryAccountScores:
			return queryAccountScores(ctx, k, path[1:], req)
			// TODO: Put the modules query routes
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown trust query endpoint")
		}
	}
}

func queryAccountScores(ctx sdk.Context, k Keeper, path []string, req abci.RequestQuery) ([]byte, error) {
	var param types.QueryAccountScoresParam
	codec.Cdc.MustUnmarshalJSON(req.Data, &param)

	address, err := sdk.AccAddressFromBech32(param.Address)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, err.Error())
	}
	topicIDs := strings.Split(param.TopicIDs, ",")

	vector := keeper.GetAccountScores(ctx, topicIDs, address)
	res, _ := json.Marshal(vector)

	return res, nil
}

// TODO: Add the modules query functions
// They will be similar to the above one: queryParams()
