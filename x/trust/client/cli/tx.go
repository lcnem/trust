package cli

import (
	"fmt"
	"bufio"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/lcnem/trust/x/trust/internal/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	trustTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%S transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	trustTxCmd.AddCommand(flags.PostCommands(
		// TODO: Add tx based commands
		getCmdEvaluate(cdc),
		getCmdDistributeByScore(cdc),
		getCmdDistributeByEvaluation(cdc),
	)...)

	return trustTxCmd
}

// Example:
//
// GetCmd<Action> is the CLI command for doing <Action>
// func GetCmd<Action>(cdc *codec.Codec) *cobra.Command {
// 	return &cobra.Command{
// 		Use:   "/* Describe your action cmd */",
// 		Short: "/* Provide a short description on the cmd */",
// 		Args:  cobra.ExactArgs(2), // Does your request require arguments
// 		RunE: func(cmd *cobra.Command, args []string) error {
// 			cliCtx := context.NewCLIContext().WithCodec(cdc)
// 			inBuf := bufio.NewReader(cmd.InOrStdin())
// 			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

// 			msg := types.NewMsg<Action>(/* Action params */)
// 			err = msg.ValidateBasic()
// 			if err != nil {
// 				return err
// 			}

// 			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
// 		},
// 	}
// }

func getCmdEvaluate(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "evaluate [topic_id] [to_address] [weight]",
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

func getCmdDistributeByScore(cdc *codec.Codec) *cobra.Command {
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

			msg := types.NewMsgDistributeByScore(args[0], fromAddress, coin)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

func getCmdDistributeByEvaluation(cdc *codec.Codec) *cobra.Command {
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

			msg := types.NewMsgDistributeByEvaluation(args[0], address, fromAddress, coin)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
