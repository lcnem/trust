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
	account := accAddress.String()
	scoreVector := make(pagerank.Vector)

	for _, topicID := range topicIDs {
		s, err := k.getVectorUnmarshaled(ctx, getScoreVectorKey(topicID))
		if err != nil {
			continue
		}
		scoreVector[topicID] = s[account]
	}

	return scoreVector
}

// SetEvaluation sets the evaluation
func (k Keeper) SetEvaluation(ctx sdk.Context, topicID string, fromAddress sdk.AccAddress, toAddress sdk.AccAddress, weight1000 sdk.Int) error {
	from := fromAddress.String()
	to := toAddress.String()
	linkKey := getLinkMatrixKey(topicID)
	stochasticKey := getStochasticMatrixKey(topicID)
	scoreKey := getScoreVectorKey(topicID)

	linkMatrix := k.getMatrixUnmarshaled(ctx, linkKey)
	stochasticMatrix := k.getMatrixUnmarshaled(ctx, stochasticKey)
	scoreVector, _ := k.getVectorUnmarshaled(ctx, scoreKey)

	setEvaluationAndTransition(from, to, weight1000.Int64(), &linkMatrix, &stochasticMatrix, &scoreVector)

	k.setMatrixMarshaled(ctx, linkKey, linkMatrix)
	k.setMatrixMarshaled(ctx, stochasticKey, stochasticMatrix)
	k.setVectorMarshaled(ctx, scoreKey, scoreVector)

	return nil
}

func setEvaluationAndTransition(from string, to string, weight1000 int64, linkMatrix *pagerank.Matrix, stochasticMatrix *pagerank.Matrix, scoreVector *pagerank.Vector) {
	linkMatrix.Set(from, to, float64(weight1000)/float64(1000))
	(*stochasticMatrix)[from] = pagerank.GetStochastixMatrix(*linkMatrix)[from]
	*scoreVector, _ = pagerank.TransitionScore(*scoreVector, *stochasticMatrix)
}

// DistributeTokenByScore distributes token by score
func (k Keeper) DistributeTokenByScore(ctx sdk.Context, topicID string, fromAddress sdk.AccAddress, amount sdk.Coin) error {
	scoreVector, err := k.getVectorUnmarshaled(ctx, getScoreVectorKey(topicID))
	if err != nil {
		return err
	}

	amountVector, sum := getAmountVectorAndSumByScore(amount.Amount, scoreVector)

	distribute(k, ctx, fromAddress, amount.Denom, sum, amountVector)

	return nil
}

func getAmountVectorAndSumByScore(amount sdk.Int, scoreVector pagerank.Vector) (map[string]sdk.Int, sdk.Int) {
	amountVector := make(map[string]sdk.Int)
	sum := sdk.NewInt(0)
	for acc, s := range scoreVector {
		val := sdk.NewInt(int64(s * float64(amount.Int64())))
		amountVector[acc] = val
		sum = sum.Add(val)
	}

	return amountVector, sum
}

// DistributeTokenByEvaluation distributes token by evaluation
func (k Keeper) DistributeTokenByEvaluation(ctx sdk.Context, topicID string, address sdk.AccAddress, fromAddress sdk.AccAddress, amount sdk.Coin) error {
	scoreVector, err := k.getVectorUnmarshaled(ctx, getScoreVectorKey(topicID))
	if err != nil {
		return err
	}
	stochasticMatrix := k.getMatrixUnmarshaled(ctx, getStochasticMatrixKey(topicID))

	amountVector, sum := getAmountVectorAndSumByEvaluation(address.String(), amount.Amount, scoreVector, stochasticMatrix)

	distribute(k, ctx, fromAddress, amount.Denom, sum, amountVector)

	return nil
}

func getAmountVectorAndSumByEvaluation(address string, amount sdk.Int, scoreVector pagerank.Vector, stochasticMatrix pagerank.Matrix) (map[string]sdk.Int, sdk.Int) {
	amountVector := map[string]sdk.Int{}
	sum := sdk.NewInt(0)
	for from, vec := range stochasticMatrix {
		stochastic, ok := vec[address]
		if !ok {
			continue
		}
		val := sdk.NewInt(int64(stochastic * scoreVector[from] / scoreVector[address] * float64(amount.Int64())))
		amountVector[from] = val
		sum = sum.Add(val)
	}

	return amountVector, sum
}

func getScoreVectorKey(topicID string) string {
	return fmt.Sprintf("%s/score-vector", topicID)
}

func getLinkMatrixKey(topicID string) string {
	return fmt.Sprintf("%s/link-matrix", topicID)
}

func getStochasticMatrixKey(topicID string) string {
	return fmt.Sprintf("%s/stochastic-matrix", topicID)
}

func (k Keeper) getVectorUnmarshaled(ctx sdk.Context, key string) (pagerank.Vector, error) {
	store := ctx.KVStore(k.storeKey)
	vector := pagerank.Vector{}
	err := json.Unmarshal(store.Get([]byte(key)), &vector)

	return vector, err
}

func (k Keeper) getMatrixUnmarshaled(ctx sdk.Context, key string) pagerank.Matrix {
	store := ctx.KVStore(k.storeKey)
	rawMatrix := map[string]json.RawMessage{}
	json.Unmarshal(store.Get([]byte(key)), &rawMatrix)

	matrix := pagerank.Matrix{}
	for src, raw := range rawMatrix {
		vec := pagerank.Vector{}
		json.Unmarshal(raw, &vec)
		matrix[src] = vec
	}

	return matrix
}

func (k Keeper) setVectorMarshaled(ctx sdk.Context, key string, v pagerank.Vector) error {
	store := ctx.KVStore(k.storeKey)
	binary, err := json.Marshal(v)
	if err != nil {
		return err
	}
	store.Set([]byte(key), binary)

	return nil
}

func (k Keeper) setMatrixMarshaled(ctx sdk.Context, key string, m pagerank.Matrix) error {
	store := ctx.KVStore(k.storeKey)
	rawMatrix := map[string]json.RawMessage{}

	for src, vec := range m {
		raw, _ := json.Marshal(vec)
		rawMatrix[src] = raw
	}

	binary, err := json.Marshal(rawMatrix)
	if err != nil {
		return err
	}
	store.Set([]byte(key), binary)

	return nil
}

func distribute(k Keeper, ctx sdk.Context, fromAddress sdk.AccAddress, denom string, sum sdk.Int, amountVector map[string]sdk.Int) error {
	_, err := k.BankKeeper.SubtractCoins(ctx, fromAddress, sdk.NewCoins(sdk.NewCoin(denom, sum)))
	if err != nil {
		return err
	}

	for acc, val := range amountVector {
		address, _ := sdk.AccAddressFromBech32(acc)
		k.BankKeeper.AddCoins(ctx, address, sdk.NewCoins(sdk.NewCoin(denom, val)))
	}

	return nil
}
