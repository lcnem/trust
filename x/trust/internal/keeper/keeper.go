package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"

	"github.com/KimuraYu45z/pagerank-go"
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
	scoreVector := pagerank.Vector{}

	for _, topicID := range topicIDs {
		s := k.getScoreVectorUnmarshaled(ctx, topicID)
		scoreVector[topicID] = s[account]
	}

	return scoreVector
}

// Evaluate sets the evaluation
func (k Keeper) Evalutate(ctx sdk.Context, topicID string, fromAddress sdk.AccAddress, toAddress sdk.AccAddress, weight1000 sdk.Int) {
	from := fromAddress.String()
	to := toAddress.String()

	linkMatrix := k.getLinkMatrixUnmarshaled(ctx, topicID)
	stochasticMatrix := pagerank.GetStochastixMatrix(linkMatrix)
	scoreVector := k.getScoreVectorUnmarshaled(ctx, topicID)

	setEvaluationAndTransition(from, to, weight1000.Int64(), &linkMatrix, &stochasticMatrix, &scoreVector)

	k.setLinkMatrixMarshaled(ctx, topicID, linkMatrix)
	k.setStochasticMatrixMarshaled(ctx, topicID, stochasticMatrix)
	k.setScoreVectorMarshaled(ctx, topicID, scoreVector)
}

func setEvaluationAndTransition(from string, to string, weight1000 int64, linkMatrix *pagerank.Matrix, stochasticMatrix *pagerank.Matrix, scoreVector *pagerank.Vector) {
	linkMatrix.Set(from, to, float64(weight1000)/float64(1000))
	*stochasticMatrix = pagerank.GetStochastixMatrix(*linkMatrix)
	*scoreVector = pagerank.TransitionScore(*scoreVector, *stochasticMatrix)
}


// DistributeByScore distributes token by score
func (k Keeper) DistributeByScore(ctx sdk.Context, topicID string, fromAddress sdk.AccAddress, amount sdk.Coin) error {
	scoreVector := k.getScoreVectorUnmarshaled(ctx, topicID)

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

// DistributeByEvaluation distributes token by evaluation
func (k Keeper) DistributeByEvaluation(ctx sdk.Context, topicID string, address sdk.AccAddress, fromAddress sdk.AccAddress, amount sdk.Coin) error {
	scoreVector := k.getScoreVectorUnmarshaled(ctx, topicID)

	stochasticMatrix := k.getStochasticMatrixUnmarshaled(ctx, topicID)

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


func (k Keeper) getScoreVectorUnmarshaled(ctx sdk.Context, topicID string) pagerank.Vector {
	store := ctx.KVStore(k.storeKey)
	key := fmt.Sprintf("%s/score-vector", topicID)
	var vector pagerank.Vector
	err := json.Unmarshal(store.Get([]byte(key)), &vector)
	if err != nil {
		vector = pagerank.Vector{}
	}

	return vector
}

func (k Keeper) setScoreVectorMarshaled(ctx sdk.Context, topicID string, v pagerank.Vector) {
	store := ctx.KVStore(k.storeKey)
	key := fmt.Sprintf("%s/score-vector", topicID)
	bz, _ := json.Marshal(v)
	store.Set([]byte(key), bz)
}

func (k Keeper) getLinkMatrixUnmarshaled(ctx sdk.Context, topicID string) pagerank.Matrix {
	store := ctx.KVStore(k.storeKey)
	key := fmt.Sprintf("%s/link-matrix", topicID)
	var matrix pagerank.Matrix
	err := json.Unmarshal(store.Get([]byte(key)), &matrix)
	if err != nil {
		matrix = pagerank.Matrix{}
	}

	return matrix
}

func (k Keeper) setLinkMatrixMarshaled(ctx sdk.Context, topicID string, m pagerank.Matrix) {
	store := ctx.KVStore(k.storeKey)
	key := fmt.Sprintf("%s/link-matrix", topicID)
	bz, _ := json.Marshal(m)
	store.Set([]byte(key), bz)
}

func (k Keeper) getStochasticMatrixUnmarshaled(ctx sdk.Context, topicID string) pagerank.Matrix {
	store := ctx.KVStore(k.storeKey)
	key := fmt.Sprintf("%s/stochastic-matrix", topicID)
	var matrix pagerank.Matrix
	err := json.Unmarshal(store.Get([]byte(key)), &matrix)
	if err != nil {
		matrix = pagerank.Matrix{}
	}

	return matrix
}

func (k Keeper) setStochasticMatrixMarshaled(ctx sdk.Context, topicID string, m pagerank.Matrix) {
	store := ctx.KVStore(k.storeKey)
	key := fmt.Sprintf("%s/stochastic-matrix", topicID)
	bz, _ := json.Marshal(m)
	store.Set([]byte(key), bz)
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
