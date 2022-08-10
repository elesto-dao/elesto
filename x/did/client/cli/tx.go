package cli

import (
	"crypto/sha256"
	"fmt"
	"strings"

	"github.com/coinbase/kryptology/pkg/core/curves/native/bls12381"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptodid "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/multiformats/go-multibase"
	"github.com/spf13/cobra"

	"github.com/elesto-dao/elesto/v2/x/did"
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
		NewCreateDidDocumentCmd(),
		NewAddServiceCmd(),
		NewDeleteServiceCmd(),
		NewAddControllerCmd(),
		NewDeleteControllerCmd(),
		NewAddVerificationCmd(),
		NewSetVerificationRelationshipCmd(),
		NewRevokeVerificationCmd(),
		NewLinkAriesAgentCmd(),
	)

	return cmd
}

// deriveVMType derive the verification method type from a public key
func deriveVMType(pubKey cryptodid.PubKey) (vmType did.VerificationMethodType, err error) {
	switch pubKey.(type) {
	//case *ed25519.PubKey:
	//	vmType = did.Ed25519VerificationKey2018
	case *secp256k1.PubKey:
		vmType = did.EcdsaSecp256k1VerificationKey2019
	default:
		err = did.ErrKeyFormatNotSupported
	}
	return
}

// NewCreateDidDocumentCmd defines the command to create a new DID document for the public key
// that signed the transaction
func NewCreateDidDocumentCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create-did [id]",
		Short:   "create a decentralized did (did) document",
		Example: `tx did create-did regulator --from regulator --chain-id elesto`,
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
			// pubkey
			info, err := clientCtx.Keyring.KeyByAddress(signer)
			if err != nil {
				return err
			}
			pubKey := info.GetPubKey()
			// verification method id
			vmID := didID.NewVerificationMethodID(signer.String())
			// understand the vmType
			vmType, err := deriveVMType(pubKey)
			if err != nil {
				return err
			}
			auth := did.NewVerification(
				did.NewVerificationMethod(
					vmID,
					didID,
					did.NewPublicKeyMultibase(pubKey.Bytes()),
					vmType,
				),
				[]string{did.Authentication},
				nil,
			)
			// create the message
			msg := did.NewMsgCreateDidDocument(
				didID.String(),
				did.Verifications{auth},
				did.Services{},
				signer.String(),
			)
			// validate
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			// execute
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func validateVerificationType(value string) (did.VerificationMethodType, error) {
	valid := map[string]did.VerificationMethodType{
		"Bls12381G1Key2020":                 did.Bls12381G1Key2020,
		"Bls12381G2Key2020":                 did.Bls12381G2Key2020,
		"EcdsaSecp256k1VerificationKey2019": did.EcdsaSecp256k1VerificationKey2019,
	}

	if val, found := valid[value]; found {
		return val, nil
	}
	availableTypes := []string{}

	for k := range valid {
		availableTypes = append(availableTypes, k)
	}

	return "", fmt.Errorf("invalid or missing verification type, allowed types: %s", strings.Join(availableTypes, ", "))
}

// NewAddVerificationCmd defines the command to add a verification method to a given did
func NewAddVerificationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "add-verification-method [id] [type] [pubkey] --verification-type [verification-type]",
		Short:   "add an verification method to a decentralized did (did) document",
		Example: `tx did add-verification-method emti $(elestod keys show emti -p) --from validator --chain-id elesto --verification-type EcdsaSecp256k1VerificationKey2019`,
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			rawVerificationType, _ := cmd.Flags().GetString("verification-method-type")
			verificationType, err := validateVerificationType(rawVerificationType)
			if err != nil {
				return err
			}

			var msg sdk.Msg
			var verification *did.Verification

			id := args[0]
			rawPubKey := args[1]

			// signer address
			signer := clientCtx.GetFromAddress()
			// document did
			didID := did.NewChainDID(clientCtx.ChainID, id)

			blsVerificationBuilder := func(pkBytes []byte) (*did.Verification, error) {
				pkHash := sha256.Sum256(pkBytes)
				var pkHashStr string
				pkHashStr, err = multibase.Encode(multibase.Base58BTC, pkHash[:])
				if err != nil {
					return nil, fmt.Errorf("cannot encode pubkey hash to string, %w", err)
				}

				// verification method id
				vmID := didID.NewVerificationMethodID(pkHashStr)

				return did.NewVerification(
					did.NewVerificationMethod(
						vmID,
						didID,
						did.NewPublicKeyMultibase(pkBytes),
						verificationType,
					),
					[]string{did.Authentication},
					nil,
				), nil
			}

			switch verificationType {
			case did.CosmosAccountAddress, did.Ed25519VerificationKey2018, did.X25519KeyAgreementKey2019:
				return fmt.Errorf("verification type type %s not supported", verificationType)
			case did.EcdsaSecp256k1VerificationKey2019:
				var pk cryptodid.PubKey
				err = clientCtx.Codec.UnmarshalInterfaceJSON([]byte(rawPubKey), &pk)
				if err != nil {
					return err
				}

				// verification method id
				vmID := didID.NewVerificationMethodID(sdk.MustBech32ifyAddressBytes(
					sdk.GetConfig().GetBech32AccountAddrPrefix(),
					pk.Address().Bytes(),
				))

				verification = did.NewVerification(
					did.NewVerificationMethod(
						vmID,
						didID,
						did.NewPublicKeyMultibase(pk.Bytes()),
						verificationType,
					),
					[]string{did.Authentication},
					nil,
				)

			case did.Bls12381G1Key2020:
				// rawPubKey must be multibase
				_, pkBytes, err := multibase.Decode(rawPubKey)
				if err != nil {
					return fmt.Errorf("public key format error, %w", err)
				}

				// verify that passed data is actually a G1 coordinate
				g1 := bls12381.G1{}
				switch len(pkBytes) {
				case bls12381.FieldBytes: // compressed
					pkBytesPtr := (*[bls12381.FieldBytes]byte)(pkBytes)
					if _, err = g1.FromCompressed(pkBytesPtr); err != nil {
						return fmt.Errorf("malformed bls12381 key, %w", err)
					}

				case bls12381.WideFieldBytes:
					pkBytesPtr := (*[bls12381.WideFieldBytes]byte)(pkBytes)
					if _, err = g1.FromUncompressed(pkBytesPtr); err != nil {
						return fmt.Errorf("malformed bls12381 key, %w", err)
					}

				}

				verification, err = blsVerificationBuilder(pkBytes)
				if err != nil {
					return err
				}

			case did.Bls12381G2Key2020:
				// rawPubKey must be multibase
				_, pkBytes, err := multibase.Decode(rawPubKey)
				if err != nil {
					return fmt.Errorf("public key format error, %w", err)
				}

				// verify that passed data is actually a G1 coordinate
				g2 := bls12381.G2{}
				switch len(pkBytes) {
				case bls12381.WideFieldBytes: // compressed
					pkBytesPtr := (*[bls12381.WideFieldBytes]byte)(pkBytes)
					if _, err = g2.FromCompressed(pkBytesPtr); err != nil {
						return fmt.Errorf("malformed bls12381 key, %w", err)
					}

				case bls12381.DoubleWideFieldBytes:
					pkBytesPtr := (*[bls12381.DoubleWideFieldBytes]byte)(pkBytes)
					if _, err = g2.FromUncompressed(pkBytesPtr); err != nil {
						return fmt.Errorf("malformed bls12381 key, %w", err)
					}

				}

				verification, err = blsVerificationBuilder(pkBytes)
				if err != nil {
					return err
				}
			}

			// add verification
			msg = did.NewMsgAddVerification(
				didID.String(),
				verification,
				signer.String(),
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.PersistentFlags().String("verification-method-type", "", "verification method type")

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// NewAddServiceCmd adds a new service to a given did document
func NewAddServiceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "add-service [id] [service_id] [type] [endpoint]",
		Short:   "add a service to a decentralized did (did) document",
		Example: `tx did add-service emti service:emti-agent DIDComm "https://agents.elesto.app.beta.starport.cloud/emti" --from emti --chain-id elesto`,
		Args:    cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// tx signer
			signer := clientCtx.GetFromAddress()
			// service parameters
			serviceID, serviceType, endpoint := args[1], args[2], args[3]
			// document did
			didID := did.NewChainDID(clientCtx.ChainID, args[0])

			service := did.NewService(
				serviceID,
				serviceType,
				endpoint,
			)

			msg := did.NewMsgAddService(
				didID.String(),
				service,
				signer.String(),
			)
			// broadcast
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// NewRevokeVerificationCmd revokes a verification method from a given did document
func NewRevokeVerificationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "revoke-verification-method [did_id] [verification_method_id_fragment]",
		Short:   "revoke a verification method from a decentralized did (did) document",
		Example: `tx did revoke-verification-method 575d062c-d110-42a9-9c04-cb1ff8c01f06 Z46DAL1MrJlVW_WmJ19WY8AeIpGeFOWl49Qwhvsnn2M --from alice --chain-id elesto`,
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			// document did
			didID := did.NewChainDID(clientCtx.ChainID, args[0])
			// signer
			signer := clientCtx.GetFromAddress()
			// verification method id
			vmID := didID.NewVerificationMethodID(args[1])
			// build the message
			msg := did.NewMsgRevokeVerification(
				didID.String(),
				vmID,
				signer.String(),
			)
			// validate
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			// broadcast
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// NewDeleteServiceCmd deletes a service from a DID Document
func NewDeleteServiceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "delete-service [id] [service-id]",
		Short:   "deletes a service from a decentralized did (did) document",
		Example: "tx did delete-service emti service:emti-agent --from emti --chain-id elesto",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			// document did
			didID := did.NewChainDID(clientCtx.ChainID, args[0])
			// signer
			signer := clientCtx.GetFromAddress()
			// service id
			sID := args[1]

			msg := did.NewMsgDeleteService(
				didID.String(),
				sID,
				signer.String(),
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// NewAddControllerCmd adds a controller to a did document
func NewAddControllerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "add-controller [id] [controllerAddress]",
		Short:   "updates a decentralized identifier (did) document to contain a controller",
		Example: "add-controller vasp cosmos1kslgpxklq75aj96cz3qwsczr95vdtrd3p0fslp --from emti --chain-id elesto",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			// document did
			didID := did.NewChainDID(clientCtx.ChainID, args[0])

			// did key to use as the controller
			didKey := did.NewKeyDID(args[1])

			// signer
			signer := clientCtx.GetFromAddress()

			msg := did.NewMsgAddController(
				didID.String(),
				didKey.String(),
				signer.String(),
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// NewDeleteControllerCmd deletes a controller of a did document
func NewDeleteControllerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "delete-controller [id] [controllerAddress]",
		Short:   "updates a decentralized identifier (did) document removing a controller",
		Example: "delete-controller vasp cosmos1kslgpxklq75aj96cz3qwsczr95vdtrd3p0fslp --from emti --chain-id elesto",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			// document did
			didID := did.NewChainDID(clientCtx.ChainID, args[0])

			// did key to use as the controller
			didKey := did.NewKeyDID(args[1])

			// signer
			signer := clientCtx.GetFromAddress()

			msg := did.NewMsgDeleteController(
				didID.String(),
				didKey.String(),
				signer.String(),
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// NewSetVerificationRelationshipCmd sets a verification relationship for a verification method
func NewSetVerificationRelationshipCmd() *cobra.Command {

	// relationships
	var relationships []string
	// if true do not add the default authentication relationship
	var unsafe bool

	cmd := &cobra.Command{
		Use:     "set-verification-relationship [did_id] [verification_method_id_fragment] --relationship NAME [--relationship NAME ...]",
		Short:   "sets one or more verification relationships to a key on a decentralized identifier (did) document.",
		Example: "set-verification-relationship vasp 6f1e0700-6c86-41b6-9e05-ae3cf839cdd0 --relationship capabilityInvocation",
		Args:    cobra.ExactArgs(2),

		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			// document did
			didID := did.NewChainDID(clientCtx.ChainID, args[0])

			// method id
			vmID := didID.NewVerificationMethodID(args[1])

			// signer
			signer := clientCtx.GetFromAddress()

			msg := did.NewMsgSetVerificationRelationships(
				didID.String(),
				vmID,
				relationships,
				signer.String(),
			)

			// make sure that the authentication relationship is preserved
			if !unsafe {
				msg.Relationships = append(msg.Relationships, did.Authentication)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	// add flags to set did relationships
	cmd.Flags().StringSliceVarP(&relationships, "relationship", "r", []string{}, "the relationships to set for the verification method in the DID")
	cmd.Flags().BoolVar(&unsafe, "unsafe", false, fmt.Sprint("do not ensure that '", did.Authentication, "' relationship is set"))

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
