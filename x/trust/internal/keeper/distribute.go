package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/yukimura45z/pagerank-go"
)

// DistributeTokenByScore distributes token by score
func (k Keeper) DistributeTokenByScore(ctx sdk.Context, topicID string, fromAddress sdk.AccAddress, amount sdk.Coin) error {
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

// DistributeTokenByEvaluation distributes token by evaluation
func (k Keeper) DistributeTokenByEvaluation(ctx sdk.Context, topicID string, address sdk.AccAddress, fromAddress sdk.AccAddress, amount sdk.Coin) error {
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
