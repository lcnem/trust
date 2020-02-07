package types


// Query endpoints supported by the trust querier
const (
	//TODO: Describe query parameters, update <action> with your query
	// Query<Action>    = "<action>"
	QueryAccountScores    = "account-scores"
)

/* 
Below you will be able how to set your own queries:


// QueryResList Queries Result Payload for a query
type QueryResList []string

// implement fmt.Stringer
func (n QueryResList) String() string {
	return strings.Join(n[:], "\n")
}

*/


// QueryAccountScoresParam QueryAccountScoresParam
type QueryAccountScoresParam struct {
	Address  string `json:"address"`
	TopicIDs string `json:"topic_ids"`
}

// QueryResAccountScores Queries Result Payload
type QueryResAccountScores struct {
	Scores string `json:"scores" yaml:"scores"`
}
