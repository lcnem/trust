package cli

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/lcnem/trust/x/trust/internal/types"
)

// GetTxCmd is GetTxCmd
func GetTxCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	coinTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Trust transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	coinTxCmd.AddCommand(client.PostCommands(
		getCmdEvaluate(cdc),
		getCmdDistributeTokenByScore(cdc),
		getCmdDistributeTokenByEvaluation(cdc),
	)...)

	return coinTxCmd
}

func getCmdEvaluate(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "evaluate [topic_id] [to_address] [weight1000]",
		Short: "evaluate",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			toAddress, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			weight1000, ok := sdk.NewIntFromString(args[2])
			if !ok {
				return sdk.ErrUnknownRequest(args[2])
			}

			msg := types.NewMsgEvaluate(args[0], cliCtx.GetFromAddress(), toAddress, weight1000)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

func getCmdDistributeTokenByScore(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "distribute-by-score [topic_id] [amount]",
		Short: "distribute your token by score",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			fromAddress := cliCtx.GetFromAddress()

			coin, err := sdk.ParseCoin(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgDistributeTokenByScore(args[0], fromAddress, coin)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

func getCmdDistributeTokenByEvaluation(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "distribute-by-evaluation [topic_id] [address] [amount]",
		Short: "distribute your token by evaluation",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			address, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			fromAddress := cliCtx.GetFromAddress()

			coin, err := sdk.ParseCoin(args[2])
			if err != nil {
				return err
			}

			msg := types.NewMsgDistributeTokenByEvaluation(args[0], address, fromAddress, coin)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
