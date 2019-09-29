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
		Use:   "account-scores [address] [topic-ids[,]]",
		Short: "get account scores",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/account-scores/%s", queryRoute, args[0]), []byte(args[1]))
			if err != nil {
				fmt.Printf("could not get account scores\n")
				return nil
			}

			var out types.QueryResAccountScores
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
