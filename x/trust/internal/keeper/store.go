package keeper

import (
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/yukimura45z/pagerank-go"
)

func (k Keeper) getScoreVectorUnmarshaled(ctx sdk.Context, topicID string) pagerank.Vector {
	store := ctx.KVStore(k.storeKey)
	key := fmt.Sprintf("%s/score-vector", topicID)
	vector := pagerank.Vector{}
	json.Unmarshal(store.Get([]byte(key)), &vector)

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
	matrix := pagerank.Matrix{}
	json.Unmarshal(store.Get([]byte(key)), &matrix)

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
	matrix := pagerank.Matrix{}
	json.Unmarshal(store.Get([]byte(key)), &matrix)

	return matrix
}

func (k Keeper) setStochasticMatrixMarshaled(ctx sdk.Context, topicID string, m pagerank.Matrix) {
	store := ctx.KVStore(k.storeKey)
	key := fmt.Sprintf("%s/stochastic-matrix", topicID)
	bz, _ := json.Marshal(m)
	store.Set([]byte(key), bz)
}
