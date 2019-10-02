package rest

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"

	"github.com/gorilla/mux"
)

// RegisterRoutes - Central function to define routes that get registered by the main application
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, storeName string) {
	r.HandleFunc(fmt.Sprintf("/%s/evaluate", storeName), evaluateHandler(cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/%s/distribute-by-score", storeName), distributeByScoreHandler(cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/%s/distribute-by-evaluation", storeName), distributeByEvaluationHandler(cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/%s/score/{address}", storeName), getAccountScoresHandler(cliCtx, storeName)).Queries("topic-ids", "{topic-ids}").Methods("GET")
}
