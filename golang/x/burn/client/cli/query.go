package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
	"github.com/summa-tx/burning-bitcoin/golang/x/burn/types"
)

// GetQueryCmd sets up query CLI commands
func GetQueryCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	burnQueryCommand := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the burn module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	burnQueryCommand.AddCommand(client.GetCommands(
		GetCommandValidated(storeKey, cdc),
	)...)
	return burnQueryCommand
}

// GetCommandValidated queries information about a name
func GetCommandValidated(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "validated [txid]",
		Short: "check if a tx is previously validated",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			digest := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/validated/%s", queryRoute, digest), nil)
			if err != nil {
				fmt.Printf("could not get validation - %s \n", digest)
				return nil
			}

			var out types.QueryResValidated
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
