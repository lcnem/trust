package keeper

import (
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/yukimura45z/pagerank-go"
)

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

func (k Keeper) getMatrixUnmarshaled(ctx sdk.Context, key string) (pagerank.Matrix, error) {
	store := ctx.KVStore(k.storeKey)
	matrix := pagerank.Matrix{}
	err := json.Unmarshal(store.Get([]byte(key)), &matrix)

	return matrix, err
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
	matrix := pagerank.Matrix{}

	binary, err := json.Marshal(matrix)
	if err != nil {
		return err
	}
	store.Set([]byte(key), binary)

	return nil
}
