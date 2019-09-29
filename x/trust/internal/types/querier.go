package types

import (
	"github.com/yukimura45z/pagerank-go"
)

// QueryResAccountScores Queries Result Payload for a master-address query
type QueryResAccountScores pagerank.Vector

// QueryAccountScoresParam QueryAccountScoresParam
type QueryAccountScoresParam struct {
	Address  string `json:"address"`
	TopicIDs string `json:"topic_ids"`
}

// NewQueryAccountScoresParam QueryAccountScoresParam
func NewQueryAccountScoresParam(address string, topicIDs string) QueryAccountScoresParam {
	return QueryAccountScoresParam{
		Address:  address,
		TopicIDs: topicIDs,
	}
}
