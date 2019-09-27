package rest

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"

	"github.com/gorilla/mux"
)

const (
	restName = "coin"
)

// RegisterRoutes - Central function to define routes that get registered by the main application
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, storeName string) {
	r.HandleFunc(fmt.Sprintf("/%s/evaluate", storeName), evaluateHandler(cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/%s/distribute", storeName), distributeHandler(cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/%s/master-address", storeName), getAccountScoresHandler(cliCtx, storeName)).Methods("GET")
}
