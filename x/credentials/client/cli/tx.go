package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/elesto-dao/elesto/x/credentials"
	"github.com/elesto-dao/elesto/x/did"
	"github.com/spf13/cobra"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        did.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", did.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	// this line is used by starport scaffolding # 1
	cmd.AddCommand(
		NewRegisterIssuerCmd(),
		NewUpdateRevocationList(),
	)

	return cmd
}


// NewRegisterIssuerCmd defines the command to create a new IBC light client.
func NewRegisterIssuerCmd() *cobra.Command {

	var revocationServiceURL string

	cmd := &cobra.Command{
		Use:     "register-issuer [id]",
		Short:   "register a credential issuer for a did",
		Example: "elestod credentials register-issuer example-issuer",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			// did
			didID := did.NewChainDID(clientCtx.ChainID, args[0])
			// verification
			signer := clientCtx.GetFromAddress()

			// initialize revocation list
			rl, seed, err := InitRevocationList()
			if err != nil {
				return err
			}
			// initialize the issuer
			issuer := &credentials.CredentialIssuer{
				Did:         didID.String(),
				Revocations: rl,
			}
			// create the message
			msg := credentials.NewMsgRegisterCredentialIssuerRequest(
				issuer,
				revocationServiceURL,
				signer.String(),
			)
			// execute
			if err := tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg); err != nil {
				return err
			}
			fmt.Println("Store this information in a safe place:")
			fmt.Println("credential issuer DID:  ", didID.String())
			fmt.Println("revocation list secret: ", seed)
			return nil
		},
	}

	// add flags to set did relationships
	cmd.Flags().StringVarP(&revocationServiceURL, "revocationServiceURL", "r", "http://localhost:2109/revocations", "The revocation service URL to check for revoked credentials")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// NewUpdateRevocationList defines the command to create a new IBC light client.
func NewUpdateRevocationList() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "update-revocation-list [id] [revocation-list-secret] [entry]",
		Short:   "adds a new entry to the issuer revocation list",
		Example: "elestod credentials update-revocation-list example-issuer 1234 elesto-dao:membership/andrea",
		Args:    cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			id, seed, credentialID := args[0], args[1], args[2]
			// did
			didID := did.NewChainDID(clientCtx.ChainID, id)
			// verification
			signer := clientCtx.GetFromAddress()

			// TODO: fetch current revocation list

			rl, err := BuildRevocationList(seed, credentialID)
			if err != nil {
				return err
			}

			msg := credentials.NewMsgUpdateRevocationListRequest(didID.String(), rl, signer.String())
			// execute
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	// add flags to set did relationships
	//cmd.Flags().StringSliceVarP(&relationships, "relationship", "r", []string{}, "the relationships to set for the verification method in the DID")
	//cmd.Flags().BoolVar(&unsafe, "unsafe", false, fmt.Sprint("do not ensure that '", did.Authentication, "' relationship is set"))

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
