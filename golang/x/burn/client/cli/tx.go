package cli

import (
	"encoding/json"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/summa-tx/burning-bitcoin/golang/x/burn/types"

	"github.com/summa-tx/bitcoin-spv/golang/btcspv"
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
		GetCmdBurnProof(cdc),
	)...)

	return burnTxCmd
}

type burnProofCall struct {
	Proof       btcspv.SPVProof        `json:"proof"`
	HeaderChain []btcspv.BitcoinHeader `json:"headers"`
	Address     string                 `json:"signer"`
}

// GetCmdBurnProof is the CLI command for sending a DoThing transaction
func GetCmdBurnProof(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "burnproof [lots of json]",
		Short: "prove a burn",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			var req burnProofCall
			json.Unmarshal([]byte(args[0]), &req)

			cliCtx := context.NewCLIContext().WithCodec(cdc)

			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			msg := types.NewMsgBurnProof(
				req.Proof,
				req.HeaderChain,
				req.Address,
				cliCtx.GetFromAddress())
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
