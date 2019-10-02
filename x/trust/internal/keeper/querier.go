package keeper

import (
	"encoding/json"
	"strings"

	"github.com/cosmos/cosmos-sdk/codec"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lcnem/lcnem-trust/x/trust/internal/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// NewQuerier is the module level router for state queries
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case types.QueryAccountScores:
			return queryAccountScores(ctx, path[1:], req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown trust query endpoint")
		}
	}
}

func queryAccountScores(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	var param types.QueryAccountScoresParam
	codec.Cdc.MustUnmarshalJSON(req.Data, &param)

	address, err := sdk.AccAddressFromBech32(param.Address)
	if err != nil {
		return nil, sdk.ErrInvalidAddress(address.String())
	}
	topicIDs := strings.Split(param.TopicIDs, ",")

	vector := keeper.GetAccountScores(ctx, topicIDs, address)
	res, _ := json.Marshal(vector)

	return res, nil
}
