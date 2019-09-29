package types

import (
	"encoding/json"

	"github.com/yukimura45z/pagerank-go"
)

// QueryResAccountScores Queries Result Payload for a master-address query
type QueryResAccountScores pagerank.Vector

// implement fmt.Stringer
func (r QueryResAccountScores) String() string {
	bz, _ := json.Marshal(r)
	return string(bz)
}
