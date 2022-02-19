package cli

import (
	"context"
	"fmt"
	"github.com/elesto-dao/elesto/x/credentials"

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
		NewQueryCredentialIssuerCmd(),
		NewQueryRevocationListCmd(),
		NewQueryPublicCredentialsCmd(),
	)

	return cmd
}

func NewQueryCredentialIssuerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "issuer [id]",
		Short: "get a credential issuer",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := credentials.NewQueryClient(clientCtx)

			did := did.NewChainDID(clientCtx.ChainID, args[0])

			result, err := queryClient.CredentialIssuer(
				context.Background(),
				&credentials.QueryCredentialIssuerRequest{
					Id: did.String(),
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

func NewQueryRevocationListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "revocation-list [id]",
		Short: "get the revocation list for a credential issuer",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := credentials.NewQueryClient(clientCtx)

			// did
			did := did.NewChainDID(clientCtx.ChainID, args[0])

			result, err := queryClient.RevocationList(
				context.Background(),
				&credentials.QueryRevocationListRequest{
					Id: did.String(),
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


func NewQueryPublicCredentialsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "public-credentials",
		Short: "list public credentials",
		Example: "elestod credentials query public-credentials",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := credentials.NewQueryClient(clientCtx)

			result, err := queryClient.PublicCredentials(
				context.Background(),
				&credentials.QueryPublicCredentialsRequest{},
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
