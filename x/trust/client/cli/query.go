package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/lcnem/lcnem-trust/x/trust/internal/types"
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

			params := types.NewQueryAccountScoresParam(args[0], args[1])
			bz, err := cdc.MarshalJSON(params)
			if err != nil {
				return err
			}

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/account-scores", queryRoute), bz)
			if err != nil {
				fmt.Printf(err.Error())
				return nil
			}

			var out types.QueryResAccountScores
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
