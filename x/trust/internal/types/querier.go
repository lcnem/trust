package types

// QueryResAccountScores Queries Result Payload for a master-address query
type QueryResAccountScores struct {
	Scores string `json:"scores" yaml:"scores"`
}

// implement fmt.Stringer
func (r QueryResAccountScores) String() string {
	return r.Scores
}

type QueryAccountScoresParam struct {
	Address  string `json:"address"`
	TopicIDs string `json:"topic_ids"`
}
