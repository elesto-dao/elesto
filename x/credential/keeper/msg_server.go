package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/xeipuuv/gojsonschema"

	"github.com/elesto-dao/elesto/x/credential"
	"github.com/elesto-dao/elesto/x/did"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the identity MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) credential.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ credential.MsgServer = msgServer{}

func (k msgServer) PublishCredentialDefinition(
	goCtx context.Context,
	msg *credential.MsgPublishCredentialDefinitionRequest,
) (*credential.MsgPublishCredentialDefinitionResponse, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)
	k.Logger(ctx).Info("request to register a CredentialDefinition", "credential Definition ID", msg.CredentialDefinition.Id)

	// check if the credential definition exists
	if _, found := k.GetCredentialDefinition(ctx, msg.CredentialDefinition.Id); found {
		err := sdkerrors.Wrapf(credential.ErrCredentialDefinitionFound, "a credential definition with did %s already exists", msg.CredentialDefinition.Id)
		k.Logger(ctx).Error(err.Error())
		return nil, err
	}

	// resolve the publisher DID
	if _, err := k.did.ResolveDid(ctx, did.DID(msg.CredentialDefinition.PublisherId)); err != nil {
		err = sdkerrors.Wrapf(did.ErrDidDocumentFound, "the credential publisher DID %v cannot be resolved", msg.CredentialDefinition.PublisherId)
		k.Logger(ctx).Error(err.Error())
		return nil, err
	}

	// persist the credential definition
	k.SetCredentialDefinition(ctx, msg.CredentialDefinition)

	k.Logger(ctx).Info("created CredentialDefinition", "definitionId", msg.CredentialDefinition.Id, "publisher", msg.CredentialDefinition.PublisherId)

	// TODO: events

	return &credential.MsgPublishCredentialDefinitionResponse{}, nil
}

func (k msgServer) UpdateCredentialDefinition(
	goCtx context.Context,
	msg *credential.MsgUpdateCredentialDefinitionRequest,
) (*credential.MsgUpdateCredentialDefinitionResponse, error) {

	return nil, fmt.Errorf("not implemented")
}

func (k msgServer) IssuePublicVerifiableCredential(
	goCtx context.Context,
	msg *credential.MsgIssuePublicVerifiableCredentialRequest,
) (*credential.MsgIssuePublicVerifiableCredentialResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	k.Logger(ctx).Info("request to issuer a PublicCredential", "credential Definition ID", msg.Credential.Id)

	var (
		err error
		cd  credential.CredentialDefinition
		wc  *credential.WrappedCredential
	)

	// fetch the credential definition
	var found bool
	if cd, found = k.GetCredentialDefinition(ctx, msg.CredentialDefinitionDid); !found {
		err = sdkerrors.Wrapf(credential.ErrCredentialDefinitionNotFound, "credential definition %s not found", msg.CredentialDefinitionDid)
		k.Logger(ctx).Error(err.Error())
		return nil, err
	}
	// verify that can be published
	if !cd.IsPublic {
		err = sdkerrors.Wrapf(credential.ErrCredentialNotIssuable, "the credential definition %s is defined as non-public", msg.CredentialDefinitionDid)
		k.Logger(ctx).Error(err.Error())
		return nil, err
	}
	// verify that is not suspended
	if !cd.IsActive {
		err = sdkerrors.Wrapf(credential.ErrCredentialNotIssuable, "the credential definition %s issuance is suspended", msg.CredentialDefinitionDid)
		k.Logger(ctx).Error(err.Error())
		return nil, err
	}
	// Wrap the credential
	if wc, err = credential.NewWrappedCredential(msg.Credential); err != nil {
		err = sdkerrors.Wrapf(credential.ErrInvalidCredential, "the credential %s is malformed: %v", msg.Credential, err)
		k.Logger(ctx).Error(err.Error())
		return nil, err
	}
	// verify the credential against the schema
	schema, err := gojsonschema.NewSchema(gojsonschema.NewBytesLoader(cd.Schema))
	if err != nil {
		err = sdkerrors.Wrapf(credential.ErrCredentialDefinitionCorrupted, "the credential definition %s is corrupted: %v", msg.CredentialDefinitionDid, cd)
		k.Logger(ctx).Error(err.Error())
		return nil, err
	}
	crL := gojsonschema.NewBytesLoader(wc.GetBytes())
	dataValidator, err := schema.Validate(crL)
	if err != nil {
		err = sdkerrors.Wrapf(credential.ErrInvalidCredential, "the credential doesn't match the schema: %v", msg.CredentialDefinitionDid)
		k.Logger(ctx).Error(err.Error())
		return nil, err
	}

	if !dataValidator.Valid() {
		err = sdkerrors.Wrapf(credential.ErrCredentialSchema, "schema: %s, errors: %v", msg.CredentialDefinitionDid, dataValidator.Errors())
		k.Logger(ctx).Error(err.Error())
		return nil, err
	}

	// validate the proof
	if err = ValidateProof(ctx, k.Keeper, wc, did.Authentication, did.AssertionMethod); err != nil {
		err = sdkerrors.Wrapf(credential.ErrInvalidProof, "%v", err)
		k.Logger(ctx).Error(err.Error())
		return nil, err
	}
	k.SetPublicCredential(ctx, msg.Credential)

	// TODO fire events
	return &credential.MsgIssuePublicVerifiableCredentialResponse{}, err
}

// ValidateProof validate the proof of a verifiable credential
func ValidateProof(ctx sdk.Context, k Keeper, wc *credential.WrappedCredential, verificationRelationships ...string) (err error) {
	// resolve the issuer
	doc, err := k.did.ResolveDid(ctx, wc.GetIssuerDID())
	if err != nil {
		return fmt.Errorf("issuer DID not resolvable %w", err)
	}

	// see if the subject is a did
	// TODO: fix this GetSubjectID
	if id, isDID := wc.GetSubjectID(); isDID {
		// resolve the subject
		if _, err = k.did.ResolveDid(ctx, did.DID(id)); err != nil {
			return fmt.Errorf("subject DID not resolvable %w", err)
		}
	}

	// verify the signature
	if wc.Proof == nil {
		return fmt.Errorf("missing credential proof")
	}
	//check relationships
	authorized := false
	methodRelationships := doc.GetVerificationRelationships(wc.Proof.VerificationMethod)
Outer:
	for _, gotR := range methodRelationships {
		for _, wantR := range verificationRelationships {
			if gotR == wantR {
				authorized = true
				break Outer
			}
		}
	}
	// verify the relationships
	if !authorized {
		return fmt.Errorf("unauthorized, verification method ID not listed in any of the required relationships in the issuer did (want %v, got %v) ", verificationRelationships, methodRelationships)
	}
	// get the address in the verification method
	issuerAddress, err := doc.GetVerificationMethodBlockchainAddress(wc.Proof.VerificationMethod)
	if err != nil {
		return fmt.Errorf("the issuer address cannot be retrieved due to %w", err)
	}

	// verify that is the same of the vc
	issuerAccount, err := sdk.AccAddressFromBech32(issuerAddress)
	if err != nil {
		return fmt.Errorf("failed to convert the issuer address to account %v due to %w", issuerAddress, err)
	}
	// get the public key from the account
	pk, err := k.account.GetPubKey(ctx, issuerAccount)
	if err != nil || pk == nil {
		return fmt.Errorf("issuer public key not found %w", err)
	}
	//
	if err = wc.Validate(pk); err != nil {
		return err

	}
	return nil
}
