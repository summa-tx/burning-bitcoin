package cli

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/summa-tx/burning-bitcoin/golang/x/burn/types"
)

// GetTxCmd sets up transaction CLI commands
func GetTxCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	burnTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Burn transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	burnTxCmd.AddCommand(client.PostCommands(
		GetCmdDoThing(cdc),
	)...)

	return burnTxCmd
}

// GetCmdDoThing is the CLI command for sending a DoThing transaction
func GetCmdDoThing(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "doThing",
		Short: "do a thing",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			msg := types.NewMsgDoThing(cliCtx.GetFromAddress())
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
