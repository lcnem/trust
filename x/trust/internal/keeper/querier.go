package keeper

import (
	"strings"

	"github.com/cosmos/cosmos-sdk/codec"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// query endpoints supported by the trust Querier
const (
	QueryAccountScores = "account-scores"
)

// NewQuerier is the module level router for state queries
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryAccountScores:
			return queryAccountScores(ctx, path[1:], req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown trust query endpoint")
		}
	}
}

func queryAccountScores(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	address, err := sdk.AccAddressFromBech32(path[0])
	if err != nil {
		return nil, sdk.ErrInvalidAddress(address.String())
	}
	topicIDs := strings.Split(path[1], ",")

	vector := keeper.GetAccountScores(ctx, topicIDs, address)

	res, err := codec.MarshalJSONIndent(keeper.cdc, vector)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return res, nil
}
