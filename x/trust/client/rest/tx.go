package rest

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/lcnem/lcnem-trust/x/trust/internal/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
)

type evaluateReq struct {
	BaseReq     rest.BaseReq `json:"base_req"`
	TopicID     string       `json:"topic_id"`
	FromAddress string       `json:"from_address"`
	ToAddress   string       `json:"to_address"`
	Weight1000  sdk.Int      `json:"weight1000"`
}

func evaluateHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req evaluateReq

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		toAddress, err := sdk.AccAddressFromBech32(req.ToAddress)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		fromAddress, err := sdk.AccAddressFromBech32(req.FromAddress)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// create the message
		msg := types.NewMsgEvaluate(req.TopicID, fromAddress, toAddress, req.Weight1000)
		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

type distributeReq struct {
	BaseReq     rest.BaseReq `json:"base_req"`
	TopicID     string       `json:"topic_id"`
	FromAddress string       `json:"from_address"`
	Amount      sdk.Coin     `json:"amount"`
}

func distributeHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req distributeReq
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		fromAddress, err := sdk.AccAddressFromBech32(req.FromAddress)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// create the message
		msg := types.NewMsgDistributeTokenByScore(req.TopicID, fromAddress, req.Amount)
		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}
