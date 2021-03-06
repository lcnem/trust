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

// SetEvaluation sets the evaluation
func (k Keeper) SetEvaluation(ctx sdk.Context, topicID string, fromAddress sdk.AccAddress, toAddress sdk.AccAddress, weight1000 sdk.Int) {
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
