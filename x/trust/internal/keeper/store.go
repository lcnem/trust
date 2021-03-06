package keeper

import (
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/KimuraYu45z/pagerank-go"
)

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
