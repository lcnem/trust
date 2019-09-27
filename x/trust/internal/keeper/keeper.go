package keeper

import (
	"encoding/json"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"

	"github.com/yukimura45z/pagerank-go"
)

// Keeper maintains the link to storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	cdc *codec.Codec // The wire codec for binary encoding/decoding.

	storeKey sdk.StoreKey // Unexposed key to access store from sdk.Context

	BankKeeper bank.Keeper
}

// NewKeeper creates new instances of the coin Keeper
func NewKeeper(cdc *codec.Codec, storeKey sdk.StoreKey, bankKeeper bank.Keeper) Keeper {
	return Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		BankKeeper: bankKeeper,
	}
}

// GetAccountScores returns the score
func (k Keeper) GetAccountScores(ctx sdk.Context, topicIDs []string, accAddress sdk.AccAddress) pagerank.Vector {
	store := ctx.KVStore(k.storeKey)
	account := accAddress.String()
	vector := make(pagerank.Vector)

	for _, topicID := range topicIDs {
		key := fmt.Sprintf("%s/score", topicID)
		v := make(pagerank.Vector)
		err := json.Unmarshal(store.Get([]byte(key)), &v)
		if err != nil {
			continue
		}
		vector[topicID] = v[account]
	}

	return vector
}

// SetEvaluation sets the evaluation
func (k Keeper) SetEvaluation(ctx sdk.Context, topicID string, fromAddress sdk.AccAddress, toAddress sdk.AccAddress, weight1000 sdk.Int) error {
	store := ctx.KVStore(k.storeKey)
	key := fmt.Sprintf("%s/link-matrix", topicID)
	from := fromAddress.String()
	to := toAddress.String()
	linkMatrix := make(pagerank.Matrix)

	err := json.Unmarshal(store.Get([]byte(key)), &linkMatrix)
	if err != nil {
		return err
	}
	if linkMatrix[from] == nil {
		linkMatrix[from] = make(pagerank.Vector)
	}

	linkMatrix[from][to] = float64(weight1000.Int64()) / float64(1000)
	binary, err := json.Marshal(linkMatrix)
	if err != nil {
		return err
	}
	store.Set([]byte(key), binary)

	key = fmt.Sprintf("%s/stochastic-matrix", topicID)
	stochasticMatrix := make(pagerank.Matrix)

	err = json.Unmarshal(store.Get([]byte(key)), &stochasticMatrix)
	if err != nil {
		return err
	}

	stochasticMatrix[from] = pagerank.GetStochastixMatrix(linkMatrix)[from]
	binary, err = json.Marshal(stochasticMatrix)
	if err != nil {
		return err
	}
	store.Set([]byte(key), binary)

	key = fmt.Sprintf("%s/score", topicID)
	score := make(pagerank.Vector)
	err = json.Unmarshal(store.Get([]byte(key)), &score)
	if err != nil {
		score[from] = 0.5
		score[to] = 0.5
	}

	score = pagerank.TransitionScore(score, stochasticMatrix)
	binary, err = json.Marshal(score)
	if err != nil {
		return err
	}
	store.Set([]byte(key), binary)

	return nil
}

// DistributeTokenByScore distributes token by score
func (k Keeper) DistributeTokenByScore(ctx sdk.Context, topicID string, fromAddress sdk.AccAddress, amount sdk.Coin) error {
	store := ctx.KVStore(k.storeKey)
	key := fmt.Sprintf("%s/score", topicID)
	score := make(pagerank.Vector)

	err := json.Unmarshal(store.Get([]byte(key)), &score)
	if err != nil {
		return err
	}

	distribution := make(map[string]sdk.Int)
	sum := sdk.NewInt(0)
	for acc, s := range score {
		val, _ := sdk.NewIntFromString(fmt.Sprintf("%f", s*float64(amount.Amount.Int64())))
		distribution[acc] = val
		sum = sum.Add(val)
	}

	_, err = k.BankKeeper.SubtractCoins(ctx, fromAddress, sdk.NewCoins(sdk.NewCoin(amount.Denom, sum)))
	if err != nil {
		return err
	}

	for acc, val := range distribution {
		address, _ := sdk.AccAddressFromBech32(acc)
		k.BankKeeper.AddCoins(ctx, address, sdk.NewCoins(sdk.NewCoin(amount.Denom, val)))
	}

	return nil
}
