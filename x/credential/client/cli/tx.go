package cli

import (
	"encoding/base64"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/noandrea/rl2020"
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
		NewPublishCredentialDefinitionCmd(),
		NewIssuePublicCredentialCmd(),
		NewCreateRevocationListCmd(),
		NewUpdateRevocationListCmd(),
	)

	return cmd
}

func exTx(cmd ...string) string {
	return fmt.Sprintln(version.AppName, "tx", credential.ModuleName, strings.Join(cmd, " "))
}

// NewIssuePublicCredentialCmd defines the command to publish credentials
func NewIssuePublicCredentialCmd() *cobra.Command {

	var (
		command           = "issue-public-credential"
		signOnly          bool
		credentialFileOut string
	)

	cmd := &cobra.Command{
		Use:     fmt.Sprintln(command, "credential-definition-id", "credential_file"),
		Short:   "issue a public, on-chain, credential",
		Example: exTx(command, "example-definition-id", "credential.json"),
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			var (
				cID, credentialFile = args[0], args[1]
				definitionDID       = did.NewChainDID(clientCtx.ChainID, cID)
				signer              = clientCtx.GetFromAddress()
				pwc                 *credential.WrappedCredential
			)

			// initialize the definition
			if pwc, err = credential.NewWrappedPublicCredentialFromFile(credentialFile); err != nil {
				println("error building credential definition", err)
				return err
			}
			// get the issuer did
			vmID := pwc.GetIssuerDID().NewVerificationMethodID(signer.String())
			if err = sign(pwc, clientCtx.Keyring, signer, vmID); err != nil {
				println("error signing the credential:", err)
				return err
			}
			// write to the output file
			if !credential.IsEmpty(credentialFileOut) {
				if err = os.WriteFile(credentialFileOut, pwc.GetBytes(), 0600); err != nil {
					fmt.Printf("error writing the credential to %v: %v", credentialFileOut, err)
					return err
				}
				fmt.Sprintln("credential exported to", credentialFileOut)
				if signOnly {
					return nil
				}
			}
			// create the message
			msg := credential.NewMsgIssuePublicVerifiableCredentialRequest(
				pwc.PublicVerifiableCredential,
				definitionDID,
				signer,
			)
			// execute
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	// add flags
	cmd.Flags().StringVar(&credentialFileOut, "export", "", "export the signed credential to a json file")
	cmd.Flags().BoolVar(&signOnly, "sign-only", false, "only sign the credential, do not broadcast (requires  --export)")
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// NewCreateRevocationListCmd update a revocation list
func NewCreateRevocationListCmd() *cobra.Command {

	var (
		command            = "create-revocation-list"
		issuerDIDstr       string
		definitionID       string
		revocationListSize int
		revocationIndexes  []int
	)

	cmd := &cobra.Command{
		Use:     fmt.Sprintln(command, "revocation-credential-id"),
		Short:   "create a revocation list credential",
		Example: exTx(command, "https://revocations.id/list/001"),
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			var (
				cID           = args[0]
				signer        = clientCtx.GetFromAddress()
				definitionDID = did.NewChainDID(clientCtx.ChainID, definitionID)
				pwc           *credential.WrappedCredential
				rl            rl2020.RevocationList2020
				issuerDID     = did.NewChainDID(clientCtx.ChainID, signer.String())
			)
			// REVOCATION LIST CREATION
			// create the revocation list
			if rl, err = rl2020.NewRevocationList(cID, revocationListSize); err != nil {
				err = fmt.Errorf("revocation list corrupted: %w", err)
				return err
			}
			// update the revocation list
			if err = rl.Revoke(revocationIndexes...); err != nil {
				err = fmt.Errorf("credential revocations failed: %w", err)
				return err
			}
			// check if credential exits
			if _, err = queryPublicCredential(credential.NewQueryClient(clientCtx), cID); err == nil {
				err = fmt.Errorf("revocation list credential %s exists", cID)
				return err
			}
			// if issuer is set use the provided one
			if issuerDIDstr != "" {
				issuerDID = did.DID(issuerDIDstr)
			}
			// PUBLIC CREDENTIAL
			// create the credential
			if pwc, err = credential.NewWrappedCredential(
				credential.NewPublicVerifiableCredential(
					cID,
					credential.WithIssuerDID(issuerDID),
					credential.WithContext("https://w3id.org/vc-revocation-list-2020/v1"),
					credential.WithType(fmt.Sprint(rl2020.TypeRevocationList2020, "Credential")),
					credential.WithIssuanceDate(time.Now()),
				),
			); err != nil {
				err = fmt.Errorf("error composing credential: %w", err)
				return err
			}
			// set the credential subject
			if err = pwc.SetSubject(rl); err != nil {
				err = fmt.Errorf("encoding of credentials failed: %w", err)
				return err
			}
			// sign the updated credential
			vmID := pwc.GetIssuerDID().NewVerificationMethodID(signer.String())
			if err = sign(pwc, clientCtx.Keyring, signer, vmID); err != nil {
				println("error signing the credential:", err)
				return err
			}
			// publish the new credential
			msg := credential.NewMsgIssuePublicVerifiableCredentialRequest(
				pwc.PublicVerifiableCredential,
				definitionDID,
				signer,
			)
			// execute
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	// add flags
	cmd.Flags().IntVar(&revocationListSize, "size", 16, "the size of the revocation list, in KB")
	cmd.Flags().StringVar(&issuerDIDstr, "issuer", "", "the issuer DID. If not set the signer key will be used as issuer")
	cmd.Flags().StringVar(&definitionID, "definition-id", "revocation-list-2020", "the RevocationList2020 definition ID")
	cmd.Flags().IntSliceVarP(&revocationIndexes, "revoke", "r", []int{}, "index of credentials to be revoked")
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// NewUpdateRevocationListCmd update a revocation list
func NewUpdateRevocationListCmd() *cobra.Command {

	var (
		command           = "update-revocation-list"
		definitionID      string
		revocationIndexes []int
		resetIndexes      []int
	)

	cmd := &cobra.Command{
		Use:     fmt.Sprintln(command, "credentialID"),
		Short:   "update a revocation list",
		Example: exTx(command, "https://revocations.id/list/001"),
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			var (
				cID           = args[0]
				signer        = clientCtx.GetFromAddress()
				definitionDID = did.NewChainDID(clientCtx.ChainID, definitionID)
				pwc           *credential.WrappedCredential
				rl            rl2020.RevocationList2020
			)

			// query the credential
			if pwc, err = queryPublicCredential(credential.NewQueryClient(clientCtx), cID); err != nil {
				err = fmt.Errorf("revocation list credential not available: %w", err)
				return err
			}
			// check the credential type, it must be a revocation list
			if !pwc.HasType(rl2020.TypeRevocationList2020) {
				err = fmt.Errorf("expecting credential type %v, found: %v", rl2020.TypeRevocationList2020, pwc.Type)
				return err
			}
			// parse the revocation list
			if rl, err = rl2020.NewRevocationListFromJSON(pwc.PublicVerifiableCredential.CredentialSubject); err != nil {
				err = fmt.Errorf("revocation list corrupted: %w", err)
				return err
			}
			// update the revocation list
			if err = rl.Revoke(revocationIndexes...); err != nil {
				err = fmt.Errorf("credential revocations failed: %w", err)
				return err
			}
			if err = rl.Reset(resetIndexes...); err != nil {
				err = fmt.Errorf("credential resets failed: %w", err)
				return err
			}
			// update the credential
			if err = pwc.SetSubject(rl); err != nil {
				err = fmt.Errorf("encoding of credentials failed: %w", err)
				return err
			}
			// sign the updated credential
			vmID := pwc.GetIssuerDID().NewVerificationMethodID(signer.String())
			if err = sign(pwc, clientCtx.Keyring, signer, vmID); err != nil {
				println("error signing the credential:", err)
				return err
			}
			// publish the new credential
			msg := credential.NewMsgIssuePublicVerifiableCredentialRequest(
				pwc.PublicVerifiableCredential,
				definitionDID,
				signer,
			)
			// execute
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	// add flags
	cmd.Flags().StringVar(&definitionID, "definition-id", "revocation-list-2020", "the RevocationList2020 definition ID")
	cmd.Flags().IntSliceVarP(&revocationIndexes, "revoke", "r", []int{}, "index of credentials to be revoked")
	cmd.Flags().IntSliceVarP(&resetIndexes, "reset", "t", []int{}, "index of credentials to be reset")
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// NewPublishCredentialDefinitionCmd defines the command to publish credential definitions
func NewPublishCredentialDefinitionCmd() *cobra.Command {

	var (
		command        = "publish-credential-definition"
		isPublic       bool
		inactive       bool
		publisherID    string
		descr          string
		expirationDays int
	)

	cmd := &cobra.Command{
		Use:     fmt.Sprintln(command, "id", "name", "schemaFile", "contextFile"),
		Short:   "publish a credential definition",
		Example: exTx(command, "example-definition-id", "example-credential", "schema.json", "vocab.json"),
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
	data := wc.GetBytes()
	signature, pubKey, err := keyring.SignByAddress(address, data)
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
