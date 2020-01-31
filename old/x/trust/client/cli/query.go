package cli

import (
	"encoding/json"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/lcnem/trust/x/trust/internal/types"
	"github.com/spf13/cobra"
)

// GetQueryCmd returns query commands
func GetQueryCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	coinQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the trust module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	coinQueryCmd.AddCommand(client.GetCommands(
		getCmdAccountScores(storeKey, cdc),
	)...)
	return coinQueryCmd
}

func getCmdAccountScores(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "account-scores [address] [topic_ids[,]]",
		Short: "get account scores",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			bz, _ := cdc.MarshalJSON(types.QueryAccountScoresParam{Address: args[0], TopicIDs: args[1]})

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", queryRoute, types.QueryAccountScores), bz)
			if err != nil {
				fmt.Printf(err.Error())
				return nil
			}
			var scores map[string]float64
			json.Unmarshal(res, &scores)

			return cliCtx.PrintOutput(scores)
		},
	}
}
