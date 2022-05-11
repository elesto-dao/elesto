package cli

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/elesto-dao/elesto/x/credentials"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(_ string) *cobra.Command {
	// Group did queries under a subcommand
	cmd := &cobra.Command{
		Use:                        credentials.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", credentials.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	// this line is used by starport scaffolding # 1
	cmd.AddCommand(
		NewQueryCredentialDefinitionCmd(),
		NewQueryPublicCredentialCmd(),
		NewQueryPublicCredentialsCmd(),
	)

	return cmd
}

func NewQueryCredentialDefinitionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "credential-definition [did]",
		Short:   "query a credential definition by its id",
		Example: "elestod query credentials credential-definition did:cosmos:elesto:cd-1",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := credentials.NewQueryClient(clientCtx)

			result, err := queryClient.CredentialDefinition(
				context.Background(),
				&credentials.QueryCredentialDefinitionRequest{
					Did: args[0],
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
			//clientCtx, err := client.GetClientQueryContext(cmd)
			//if err != nil {
			//	return err
			//}
			//queryClient := credentials.NewQueryClient(clientCtx)

			//// did
			//did := did.NewChainDID(clientCtx.ChainID, args[0])
			//
			//
			//if err != nil {
			//	return err
			//}
			//
			//return clientCtx.PrintProto(result)
			return fmt.Errorf("not implemented")
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func NewQueryPublicCredentialCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "public-credential ID",
		Short:   "fetch a public credential by id",
		Example: "elestod credentials query public-credential example-credential-id",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := credentials.NewQueryClient(clientCtx)

			result, err := queryClient.PublicCredential(
				context.Background(),
				&credentials.QueryPublicCredentialRequest{
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

func NewQueryPublicCredentialsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "public-credentials",
		Short:   "list public credentials",
		Example: "elestod credentials query public-credentials",
		Args:    cobra.ExactArgs(0),
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
