package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	"github.com/elesto-dao/elesto/x/credentials"
	"github.com/elesto-dao/elesto/x/did"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        credentials.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", credentials.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	// this line is used by starport scaffolding # 1
	cmd.AddCommand(
		NewPublishCredentialDefinition(),
	)

	return cmd
}

// NewPublishCredentialDefinition defines the command to publish credential definitions
func NewPublishCredentialDefinition() *cobra.Command {

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
			if !credentials.IsEmpty(publisherID) {
				publisherDID = did.DID(publisherID)
			}

			// initialize the definition
			def, err := credentials.NewCredentialDefinitionFromFile(definitionDID, publisherDID, name, descr, isPublic, !inactive, schemaFile, vocabFIle)
			if err != nil {
				println("error building credential definition", err)
				return err
			}
			// create the message
			msg := credentials.NewMsgPublishCredentialDefinitionRequest(
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
