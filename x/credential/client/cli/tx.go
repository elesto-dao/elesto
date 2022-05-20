package cli

import (
	"encoding/base64"
	"fmt"
	"os"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"

	"github.com/elesto-dao/elesto/x/credential"
	"github.com/elesto-dao/elesto/x/did"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        credential.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", credential.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	// this line is used by starport scaffolding # 1
	cmd.AddCommand(
		NewPublishCredentialDefinition(),
		NewIssuePublicCredential(),
	)

	return cmd
}

// NewPublishCredentialDefinition defines the command to publish credential definitions
func NewPublishCredentialDefinition() *cobra.Command {

	var credentialFileOut string

	cmd := &cobra.Command{
		Use:     "issue-public-credential credential-definition-id credential_file",
		Short:   "issue a public, on-chain, credential",
		Example: "elestod tx credentials issue-public-credential example-definition-id credential.json",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {

			cID, credentialFile := args[0], args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			// did
			definitionDID := did.NewChainDID(clientCtx.ChainID, cID)
			// verification
			signer := clientCtx.GetFromAddress()

			// initialize the definition
			wc, err := credential.NewWrappedPublicCredentialFromFile(credentialFile)
			if err != nil {
				println("error building credential definition", err)
				return err
			}
			// get the issuer did
			vmID := wc.GetIssuerDID().NewVerificationMethodID(signer.String())
			if err = sign(wc, clientCtx.Keyring, signer, vmID); err != nil {
				println("error signing the credential:", err)
				return err
			}
			// write to the output file
			if !credential.IsEmpty(credentialFileOut) {
				if err = os.WriteFile(credentialFileOut, wc.GetBytes(), 0600); err != nil {
					fmt.Printf("error writing the credential to %v: %v", credentialFileOut, err)
					return err
				}
			}

			if err = sign(wc, clientCtx.Keyring, signer, vmID); err != nil {
				println("error signing the credential:", err)
				return err
			}

			pvc, err := wc.GetCredential()
			if err != nil {
				println("error extracting the credential:", err)
				return err
			}

			// create the message
			msg := credential.NewMsgIssuePublicVerifiableCredentialRequest(
				pvc,
				definitionDID,
				signer,
			)
			// execute
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	// add flags
	cmd.Flags().StringVar(&credentialFileOut, "export", "", "export the signed credential to a json file")

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// NewIssuePublicCredential defines the command to publish credential definitions
func NewIssuePublicCredential() *cobra.Command {

	var (
		isPublic       bool
		inactive       bool
		publisherID    string
		descr          string
		expirationDays int
	)

	cmd := &cobra.Command{
		Use:     "publish-credential-definition id name schemaFile vocabFile",
		Short:   "publish a credential definition",
		Example: "elestod tx credentials publish-credential-definition example-definition-id example-credential schema.json vocab.json",
		Args:    cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {

			cID, name, schemaFile, vocabFIle := args[0], args[1], args[2], args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			// did
			definitionDID := did.NewChainDID(clientCtx.ChainID, cID)
			// verification
			signer := clientCtx.GetFromAddress()

			publisherDID := did.NewKeyDID(signer.String())
			if !credential.IsEmpty(publisherID) {
				publisherDID = did.DID(publisherID)
			}

			// initialize the definition
			def, err := credential.NewCredentialDefinitionFromFile(definitionDID, publisherDID, name, descr, isPublic, !inactive, schemaFile, vocabFIle)
			if err != nil {
				println("error building credential definition", err)
				return err
			}
			// create the message
			msg := credential.NewMsgPublishCredentialDefinitionRequest(
				def,
				signer.String(),
			)
			// execute
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	// add flags to set did relationships
	cmd.Flags().StringVar(&descr, "description", "", "a human-readable description about the credential usage")
	cmd.Flags().StringVarP(&publisherID, "publisher", "p", "", "the publisher DID. If not set, the DID key of the signer account will be used instead")
	cmd.Flags().BoolVar(&isPublic, "public", false, "if is set, the credential is a public one and can be issued on chain")
	cmd.Flags().BoolVar(&inactive, "inactive", false, "if is set, the credential definition will be flagged as inactive, client may refuse to issue credentials based on an inactive definition")
	cmd.Flags().IntVar(&expirationDays, "expiration", 365, "number of days that the definition can be ")
	// add flags to set did relationships
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// Sign signs a credential with a provided private key
func sign(
	wc *credential.WrappedCredential,
	keyring keyring.Keyring,
	address sdk.Address,
	verificationMethodID string,
) error {
	tm := time.Now()
	// reset the proof
	wc.Proof = nil
	// TODO: this could be expensive review this signing method
	// TODO: we can hash this an make this less expensive
	signature, pubKey, err := keyring.SignByAddress(address, wc.GetBytes())
	if err != nil {
		return err
	}

	p := credential.NewProof(
		pubKey.Type(),
		tm.Format(time.RFC3339),
		// TODO: define proof purposes
		did.AssertionMethod,
		verificationMethodID,
		base64.StdEncoding.EncodeToString(signature),
	)
	wc.Proof = &p
	return nil
}
