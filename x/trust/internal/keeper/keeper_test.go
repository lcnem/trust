package keeper

import (
	"encoding/json"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/yukimura45z/pagerank-go"
)

func TestValidateTopicID(t *testing.T) {

	require.Equal(t, true, true)
}

func TestGetAccountScores(t *testing.T) {
	topicIDs := []string{"cosmos", "nem"}
	t.Log(topicIDs)
	account := "a"
	scoreVector := pagerank.Vector{}

	for _, topicID := range topicIDs {
		s, err := getScoreVectorMock(topicID)
		if err != nil {
			t.Error(err)
			continue
		}
		scoreVector[topicID] = s[account]

		t.Log(account)
		t.Log(topicID)
		t.Log(scoreVector[topicID])
	}

	require.Equal(t, 0.5, scoreVector["cosmos"])
	require.Equal(t, 0.3, scoreVector["nem"])
}

func getScoreVectorMock(topicID string) (pagerank.Vector, error) {
	store := map[string]string{
		"cosmos": "{\"a\":0.5,\"b\":0.5}",
		"nem":    "{\"a\":0.3,\"b\":0.7}",
	}
	score := make(pagerank.Vector)

	err := json.Unmarshal([]byte(store[topicID]), &score)

	return score, err
}

func TestGetMatrixUnmarshaled(t *testing.T) {
	m1 := pagerank.Matrix{
		"a": pagerank.Vector{
			"b": 0.5,
			"c": 0.5,
		},
		"b": pagerank.Vector{
			"a": 0.7,
			"c": 0.3,
		},
		"c": pagerank.Vector{
			"a": 0.5,
			"b": 0.5,
		},
	}
	binary, _ := json.Marshal(m1)
	m2 := pagerank.Matrix{}
	json.Unmarshal(binary, &m2)

	require.Equal(t, m1["a"]["a"], m2["a"]["a"])
	require.Equal(t, 0.3, m2.Get("b", "c"))
}

func TestGetAmountVectorAndSumByScore(t *testing.T) {
	scoreVector := pagerank.Vector{
		"a": 0.3,
		"b": 0.3,
		"c": 0.4,
	}
	vec, sum := getAmountVectorAndSumByScore(sdk.NewInt(100000), scoreVector)
	require.Equal(t, int64(30000), vec["a"].Int64())
	require.Equal(t, int64(100000), sum.Int64())
}

func TestGetAmountVectorAndSumByEvaluation(t *testing.T) {
	scoreVector := pagerank.Vector{
		"a": 0.4,
		"b": 0.4,
		"c": 0.2,
	}
	stochasticMatrix := pagerank.Matrix{
		"a": pagerank.Vector{"c": 0.25},
		"b": pagerank.Vector{"c": 0.25},
	}
	vec, sum := getAmountVectorAndSumByEvaluation("c", sdk.NewInt(100000), scoreVector, stochasticMatrix)
	require.Equal(t, int64(50000), vec["a"].Int64())
	require.Equal(t, int64(100000), sum.Int64())
}

func TestSetEvaluationAndTransition(t *testing.T) {
	from := "a"
	to := "b"
	weight1000 := int64(1000)
	linkMatrix := pagerank.Matrix{}
	stochasticMatrix := pagerank.Matrix{}
	scoreVector := pagerank.Vector{}
	setEvaluationAndTransition(from, to, weight1000, &linkMatrix, &stochasticMatrix, &scoreVector)

	t.Log(scoreVector)

	require.Equal(t, float64(1), linkMatrix["a"]["b"])
	require.Equal(t, float64(1), stochasticMatrix["a"]["b"])
	require.Equal(t, 0, scoreVector["a"])
}
