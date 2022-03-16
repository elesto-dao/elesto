package cli

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/elesto-dao/elesto/x/did"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(_ string) *cobra.Command {
	// Group did queries under a subcommand
	cmd := &cobra.Command{
		Use:                        did.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", did.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	// this line is used by starport scaffolding # 1
	cmd.AddCommand(
		GetCmdQueryIdentifer(),
	)

	return cmd
}

// GetCmdQueryIdentifer querys a did on the chain using a GRPC client
func GetCmdQueryIdentifer() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "did [id]",
		Short:   "Query for an did",
		Example: "elestod query did did did:cosmos:net:elesto:bob",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := did.NewQueryClient(clientCtx)

			result, err := queryClient.DidDocument(
				context.Background(),
				&did.QueryDidDocumentRequest{
					Id: args[0],
				},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(result)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
