package cli

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/lcnem/trust/x/trust/internal/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string, cdc *codec.Codec) *cobra.Command {
	// Group trust queries under a subcommand
	trustQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	trustQueryCmd.AddCommand(
		flags.GetCommands(
			getCmdAccountScores(queryRoute, cdc),
		)...,
	)

	return trustQueryCmd

}

// TODO: Add Query Commands

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
