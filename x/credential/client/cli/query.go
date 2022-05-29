package cli

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/noandrea/rl2020"
	"github.com/spf13/cobra"

	"github.com/elesto-dao/elesto/x/credential"
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
		NewQueryCredentialDefinitionCmd(),
		NewQueryPublicCredentialCmd(),
		NewQueryPublicCredentialsByIssuerCmd(),
		NewQueryCredentialStatusCmd(),
		NewQueryPublicCredentialStatusCmd(),
		NewMakeCredentialFromSchemaCmd(),
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
			queryClient := credential.NewQueryClient(clientCtx)

			result, err := queryClient.CredentialDefinition(
				context.Background(),
				&credential.QueryCredentialDefinitionRequest{
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

func NewQueryPublicCredentialCmd() *cobra.Command {
	var printNative bool

	cmd := &cobra.Command{
		Use:     "public-credential [ID]",
		Short:   "fetch a public credential by id",
		Example: "elestod query credentials public-credential example-credential-id",
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
			return clientCtx.PrintBytes(wc.GetBytes())
		},
	}
	cmd.Flags().BoolVar(&printNative, "native", false, "if set the credential will be printed in the raw format, that is how it is stored on chain")
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func NewQueryPublicCredentialsByIssuerCmd() *cobra.Command {
	var printNative bool

	cmd := &cobra.Command{
		Use:     "public-credentials-by-issuer issuerDID",
		Short:   "list public credentials issued by a issuer ",
		Example: "elestod credentials query public-credentials-by-issuer did:cosmos:key:cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
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
			// query credentials
			if result, err = queryClient.PublicCredentialsByIssuer(
				context.Background(),
				&credential.QueryPublicCredentialsByIssuerRequest{
					Did: args[0],
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
	return cmd
}

func NewQueryCredentialStatusCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "credential-status",
		Short:   "verify the credential status of a credential",
		Example: "elestod query credential credential-status credentialFile",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			var (
				queryClient    = credential.NewQueryClient(clientCtx)
				credentialFile = args[0]
				wc             *credential.WrappedCredential
				rs             revocationStatus
			)

			// read the credential from file
			if wc, err = credential.NewWrappedPublicCredentialFromFile(credentialFile); err != nil {
				println("error building credential definition", err)
				return err
			}
			// check for revocation
			if rs, err = checkRevocation(queryClient, wc); err != nil {
				println("error processing credential revocation:", err)
				return err
			}
			// print results
			return clientCtx.PrintBytes(rs.GetBytes())
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func NewQueryPublicCredentialStatusCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "public-credential-status",
		Short:   "verify the credential status of a credential",
		Example: "elestod query credential public-credential-status credentialID",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			//
			var (
				queryClient  = credential.NewQueryClient(clientCtx)
				credentialID = args[0]
				wc           *credential.WrappedCredential
				rs           revocationStatus
			)
			// retrieve the public credential
			if wc, err = queryPublicCredential(queryClient, credentialID); err != nil {
				fmt.Println(err)
				return err
			}
			// check for revocation
			if rs, err = checkRevocation(queryClient, wc); err != nil {
				println("error processing credential revocation:", err)
				return err
			}
			// print results
			return clientCtx.PrintBytes(rs.GetBytes())
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

//
func queryPublicCredential(qc credential.QueryClient, credentialID string) (wc *credential.WrappedCredential, err error) {
	// query the public credential
	result, err := qc.PublicCredential(
		context.Background(),
		&credential.QueryPublicCredentialRequest{
			Id: credentialID,
		},
	)
	if err != nil {
		err = fmt.Errorf("public credential %s not found: %w", credentialID, err)
		return
	}
	wc, err = credential.NewWrappedCredential(result.Credential)
	if err != nil {
		err = fmt.Errorf("error processing credential %s: %w", credentialID, err)
	}
	return
}

type revocationStatus struct {
	Revoked    bool
	Credential *credential.WrappedCredential
}

func (rs revocationStatus) GetBytes() []byte {
	bs, err := json.Marshal(rs)
	if err != nil {
		panic(err)
	}
	return bs
}

// checkRevocation perform a revocation list check for a credential
func checkRevocation(qc credential.QueryClient, wc *credential.WrappedCredential) (rs revocationStatus, err error) {

	// is there a credential status to use?
	if wc.CredentialStatus == nil {
		err = fmt.Errorf("missing credentialStatus definition from the credential, revocation cannot be checked")
		return
	}
	// retrieve the revocation list
	res, err := qc.PublicCredential(
		context.Background(),
		&credential.QueryPublicCredentialRequest{
			Id: wc.CredentialStatus.RevocationListCredential,
		},
	)
	if err != nil {
		err = fmt.Errorf("revocation list credential %s not found: %w", wc.CredentialStatus.RevocationListCredential, err)
		return
	}
	// check issuer
	if res.Credential.Issuer != wc.Issuer {
		err = fmt.Errorf("credential issuer mismatch, expected: %v, got %v", wc.Issuer, res.Credential.Issuer)
		return
	}
	// load the revocation list
	rl, err := rl2020.NewRevocationListFromJSON(res.Credential.CredentialSubject)
	if err != nil {
		err = fmt.Errorf("error parsing the credential revocation list: %w", err)
		return
	}
	rs.Credential = wc
	rs.Revoked, err = rl.IsRevoked(*wc.CredentialStatus)
	return
}
