package types

// QueryAccountScoresParam QueryAccountScoresParam
type QueryAccountScoresParam struct {
	Address  string `json:"address"`
	TopicIDs string `json:"topic_ids"`
}

// QueryResAccountScores Queries Result Payload
type QueryResAccountScores struct {
	Scores string `json:"scores" yaml:"scores"`
}
