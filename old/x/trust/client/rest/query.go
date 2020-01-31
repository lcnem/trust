package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/lcnem/trust/x/trust/internal/types"

	"github.com/cosmos/cosmos-sdk/types/rest"

	"github.com/gorilla/mux"
)

func getAccountScoresHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bz, _ := cliCtx.Codec.MarshalJSON(types.QueryAccountScoresParam{Address: mux.Vars(r)["address"], TopicIDs: mux.Vars(r)["topic-ids"]})

		res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", storeName, types.QueryAccountScores), bz)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}
		var scores map[string]float64
		json.Unmarshal(res, &scores)

		rest.PostProcessResponse(w, cliCtx, scores)
	}
}
