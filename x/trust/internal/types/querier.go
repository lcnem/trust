package types

import (
	"github.com/yukimura45z/pagerank-go"
)

// QueryResAccountScores Queries Result Payload for a master-address query
type QueryResAccountScores pagerank.Vector

type QueryAccountScoresParam struct {
	Address  string `json:"address"`
	TopicIDs string `json:"topic_ids"`
}

func NewQueryAccountScoresParam(address string, topicIDs string) QueryAccountScoresParam {
	return QueryAccountScoresParam{
		Address: address,
		TopicIDs: topicIDs,
	}
}