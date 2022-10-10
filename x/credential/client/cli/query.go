package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"

	"github.com/elesto-dao/elesto/v3/x/credential"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(_ string) *cobra.Command {
	// Group did queries under a subcommand
	cmd := &cobra.Command{
		Use:                        credential.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", credential.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	// this line is used by starport scaffolding # 1
	cmd.AddCommand(
		NewQueryCredentialDefinitionsCmd(),
		NewQueryCredentialDefinitionCmd(),
		NewQueryPublicCredentialCmd(),
		NewQueryPublicCredentialsByIssuerCmd(),
		NewQueryCredentialStatusCmd(),
		NewQueryPublicCredentialStatusCmd(),
		NewMakeCredentialFromSchemaCmd(),
		NewQueryAllowedCredentialDefinitionsCmd(),
	)

	return cmd
}

func exQuery(cmd ...string) string {
	return fmt.Sprintln(version.AppName, "query", credential.ModuleName, strings.Join(cmd, " "))
}
func use(cmd string, params ...string) string {
	return fmt.Sprintln(cmd, strings.Join(params, " "))
}

func NewQueryCredentialDefinitionsCmd() *cobra.Command {

	var (
		command = "credential-definitions"
	)

	cmd := &cobra.Command{
		Use:     use(command),
		Short:   "query a credential definitions",
		Example: exQuery(command, "credential-definitions"),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := credential.NewQueryClient(clientCtx)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			params := &credential.QueryCredentialDefinitionsRequest{
				Pagination: pageReq,
			}

			result, err := queryClient.CredentialDefinitions(
				context.Background(),
				params,
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(result)
		},
	}
	flags.AddPaginationFlagsToCmd(cmd, command)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func NewQueryCredentialDefinitionCmd() *cobra.Command {
	var (
		command = "credential-definition"
	)
	cmd := &cobra.Command{
		Use:     use(command, "credentialDefinitionID"),
		Short:   "query a credential definition by its id",
		Example: exQuery(command, "did:cosmos:elesto:cd-1"),
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := credential.NewQueryClient(clientCtx)

			result, err := queryClient.CredentialDefinition(
				context.Background(),
				&credential.QueryCredentialDefinitionRequest{
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

func NewQueryPublicCredentialCmd() *cobra.Command {
	var (
		command     = "public-credential"
		printNative bool
	)

	cmd := &cobra.Command{
		Use:     use(command, "credentialID"),
		Short:   "fetch a public credential by id",
		Example: exQuery(command, "example-credential-id"),
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			var (
				queryClient  = credential.NewQueryClient(clientCtx)
				credentialID = args[0]
				wc           *credential.WrappedCredential
			)
			// query the credential
			if wc, err = queryPublicCredential(queryClient, credentialID); err != nil {
				return err
			}
			if printNative {
				return clientCtx.PrintProto(wc.PublicVerifiableCredential)
			}
			wcB, err := wc.GetBytes()
			if err != nil {
				return err
			}

			return clientCtx.PrintBytes(wcB)
		},
	}
	cmd.Flags().BoolVar(&printNative, "native", false, "if set the credential will be printed in the raw format, that is how it is stored on chain")
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func NewQueryPublicCredentialsByIssuerCmd() *cobra.Command {
	var (
		command     = "public-credentials-by-issuer"
		printNative bool
	)

	cmd := &cobra.Command{
		Use:     use(command, "issuerDID"),
		Short:   "list public credentials issued by a issuer ",
		Example: exQuery(command, "did:cosmos:key:cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8"),
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			var (
				queryClient = credential.NewQueryClient(clientCtx)
				result      *credential.QueryPublicCredentialsByIssuerResponse
				pwcs        []*credential.WrappedCredential
				pwcsJSON    []byte
			)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			// query credentials
			if result, err = queryClient.PublicCredentialsByIssuer(
				context.Background(),
				&credential.QueryPublicCredentialsByIssuerRequest{
					Did:        args[0],
					Pagination: pageReq,
				},
			); err != nil {
				return err
			}
			//
			if printNative {
				return clientCtx.PrintProto(result)
			}
			// process credentials
			for _, pvc := range result.Credential {
				pwc, errWC := credential.NewWrappedCredential(pvc)
				if errWC != nil {
					fmt.Printf("warning, cannot process credential %v: %v, for further inspection run the 'public-credential' command with the '--native' flag", pvc.Id, err)
					continue
				}
				pwcs = append(pwcs, pwc)
			}
			if pwcsJSON, err = json.Marshal(pwcs); err != nil {
				return err
			}

			return clientCtx.PrintBytes(pwcsJSON)
		},
	}
	cmd.Flags().BoolVar(&printNative, "native", false, "if set the credential will be printed in the raw format, that is how it is stored on chain")
	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, command)

	return cmd
}

func NewQueryAllowedCredentialDefinitionsCmd() *cobra.Command {

	var (
		command = "allowed-credential-definitions"
	)

	cmd := &cobra.Command{
		Use:     use(command),
		Short:   "query all allowed credential definitions",
		Example: exQuery(command, "allowed-credential-definitions"),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := credential.NewQueryClient(clientCtx)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			params := &credential.QueryAllowedPublicCredentialsRequest{
				Pagination: pageReq,
			}

			result, err := queryClient.AllowedPublicCredentials(
				context.Background(),
				params,
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(result)
		},
	}
	flags.AddPaginationFlagsToCmd(cmd, command)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
