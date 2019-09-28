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
	scoreVector := make(pagerank.Vector)

	for _, topicID := range topicIDs {
		s, err := getScoreVector(store, topicID)
		if err != nil {
			continue
		}
		scoreVector[topicID] = s[account]
	}

	return scoreVector
}

// SetEvaluation sets the evaluation
func (k Keeper) SetEvaluation(ctx sdk.Context, topicID string, fromAddress sdk.AccAddress, toAddress sdk.AccAddress, weight1000 sdk.Int) error {
	store := ctx.KVStore(k.storeKey)
	from := fromAddress.String()
	to := toAddress.String()

	linkMatrix, err := getLinkMatrix(store, topicID)

	if err != nil {
		return err
	}

	if linkMatrix[from] == nil {
		linkMatrix[from] = make(pagerank.Vector)
	}

	linkMatrix[from][to] = float64(weight1000.Int64()) / float64(1000)

	stochasticMatrix, err := getStochasticMatrix(store, topicID)
	stochasticMatrix[from] = pagerank.GetStochastixMatrix(linkMatrix)[from]

	setStochasticMatrix(store, topicID, stochasticMatrix)

	score, err := getScoreVector(store, topicID)
	if err != nil {
		score[from] = 0.5
		score[to] = 0.5
	}

	score = pagerank.TransitionScore(score, stochasticMatrix)
	setScoreVector(store, topicID, score)

	return nil
}

// DistributeTokenByScore distributes token by score
func (k Keeper) DistributeTokenByScore(ctx sdk.Context, topicID string, fromAddress sdk.AccAddress, amount sdk.Coin) error {
	store := ctx.KVStore(k.storeKey)
	score, err := getScoreVector(store, topicID)
	if err != nil {
		return err
	}

	amountvector := make(map[string]sdk.Int)
	sum := sdk.NewInt(0)
	for acc, s := range score {
		val := sdk.NewInt(int64(s * float64(amount.Amount.Int64())))
		amountvector[acc] = val
		sum = sum.Add(val)
	}

	distribute(k, ctx, fromAddress, sdk.NewCoin(amount.Denom, sum), amountvector)

	return nil
}

// DistributeTokenByEvaluation distributes token by evaluation
func (k Keeper) DistributeTokenByEvaluation(ctx sdk.Context, topicID string, address sdk.AccAddress, fromAddress sdk.AccAddress, amount sdk.Coin) error {
	store := ctx.KVStore(k.storeKey)
	score, err := getScoreVector(store, topicID)
	if err != nil {
		return err
	}
	stochasticMatrix, err := getStochasticMatrix(store, topicID)
	if err != nil {
		return err
	}

	amountvector := make(map[string]sdk.Int)
	sum := sdk.NewInt(0)
	for from, vec := range stochasticMatrix {
		stochastic, ok := vec[address.String()]
		if !ok {
			continue
		}
		val := sdk.NewInt(int64(stochastic * score[from] / score[address.String()] * float64(amount.Amount.Int64())))
		amountvector[from] = val
		sum = sum.Add(val)
	}

	distribute(k, ctx, fromAddress, sdk.NewCoin(amount.Denom, sum), amountvector)

	return nil
}

func getScoreVector(store sdk.KVStore, topicID string) (pagerank.Vector, error) {
	key := fmt.Sprintf("%s/score", topicID)
	score := make(pagerank.Vector)

	err := json.Unmarshal(store.Get([]byte(key)), &score)

	return score, err
}

func setScoreVector(store sdk.KVStore, topicID string, scoreVector pagerank.Vector) error {
	key := fmt.Sprintf("%s/score", topicID)

	binary, err := json.Marshal(scoreVector)
	if err != nil {
		return err
	}
	store.Set([]byte(key), binary)

	return nil
}

func getLinkMatrix(store sdk.KVStore, topicID string) (pagerank.Matrix, error) {
	key := fmt.Sprintf("%s/link-matrix", topicID)
	stochasticMatrix := make(pagerank.Matrix)

	err := json.Unmarshal(store.Get([]byte(key)), &stochasticMatrix)

	return stochasticMatrix, err
}

func setLinkMatrix(store sdk.KVStore, topicID string, linkMatrix pagerank.Matrix) error {
	key := fmt.Sprintf("%s/link-matrix", topicID)

	binary, err := json.Marshal(linkMatrix)
	if err != nil {
		return err
	}
	store.Set([]byte(key), binary)

	return nil
}

func getStochasticMatrix(store sdk.KVStore, topicID string) (pagerank.Matrix, error) {
	key := fmt.Sprintf("%s/stochastic-matrix", topicID)
	stochasticMatrix := make(pagerank.Matrix)

	err := json.Unmarshal(store.Get([]byte(key)), &stochasticMatrix)

	return stochasticMatrix, err
}

func setStochasticMatrix(store sdk.KVStore, topicID string, stochasticMatrix pagerank.Matrix) error {
	key := fmt.Sprintf("%s/stochastic-matrix", topicID)

	binary, err := json.Marshal(stochasticMatrix)
	if err != nil {
		return err
	}
	store.Set([]byte(key), binary)

	return nil
}

func distribute(k Keeper, ctx sdk.Context, fromAddress sdk.AccAddress, amount sdk.Coin, amountVector map[string]sdk.Int) error {
	_, err := k.BankKeeper.SubtractCoins(ctx, fromAddress, sdk.NewCoins(amount))
	if err != nil {
		return err
	}

	for acc, val := range amountVector {
		address, _ := sdk.AccAddressFromBech32(acc)
		k.BankKeeper.AddCoins(ctx, address, sdk.NewCoins(sdk.NewCoin(amount.Denom, val)))
	}

	return nil
}
