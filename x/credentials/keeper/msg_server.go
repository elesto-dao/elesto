package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/elesto-dao/elesto/x/credentials"
	"github.com/elesto-dao/elesto/x/did"
)

type msgServer struct {
	Keeper
	DIDs credentials.DidKeeper
}

// NewMsgServerImpl returns an implementation of the identity MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper, dids credentials.DidKeeper) credentials.MsgServer {
	return &msgServer{Keeper: keeper, DIDs: dids}
}

var _ credentials.MsgServer = msgServer{}

func (k msgServer) IssuePublicVerifiableCredential(
	goCtx context.Context,
	msg *credentials.MsgIssuePublicVerifiableCredentialRequest,
) (*credentials.MsgIssuePublicVerifiableCredentialResponse, error) {

	return nil, fmt.Errorf("not implemented")
}

func (k msgServer) PublishCredentialDefinition(
	goCtx context.Context,
	msg *credentials.MsgPublishCredentialDefinitionRequest,
) (*credentials.MsgPublishCredentialDefinitionResponse, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)
	k.Logger(ctx).Info("request to register a CredentialDefinition", "credential Definition ID", msg.CredentialDefinition.Id)

	// check if the credential definition exists
	if _, found := k.GetCredentialDefinition(ctx, msg.CredentialDefinition.Id); found {
		err := sdkerrors.Wrapf(credentials.ErrCredentialDefinitionFound, "a credential definition with did %s already exists", msg.CredentialDefinition.Id)
		k.Logger(ctx).Error(err.Error())
		return nil, err
	}

	// resolve the publisher DID
	if _, err := k.DIDs.ResolveDid(ctx, did.DID(msg.CredentialDefinition.PublisherId)); err != nil {
		err := sdkerrors.Wrapf(did.ErrDidDocumentFound, "the credential publisher DID cannot be resolved: %v", msg.CredentialDefinition.PublisherId)
		k.Logger(ctx).Error(err.Error())
		return nil, err
	}

	// persist the credential definition
	k.SetCredentialDefinition(ctx, msg.CredentialDefinition)

	k.Logger(ctx).Info("created CredentialDefinition", "definitionId", msg.CredentialDefinition.Id, "publisher", msg.CredentialDefinition.PublisherId)

	// TODO: events

	return &credentials.MsgPublishCredentialDefinitionResponse{}, nil
}

func (k msgServer) UpdateCredentialDefinition(
	goCtx context.Context,
	msg *credentials.MsgUpdateCredentialDefinitionRequest,
) (*credentials.MsgUpdateCredentialDefinitionResponse, error) {

	return nil, fmt.Errorf("not implemented")
}

// TODO: probably reuse for credential verification
//func executeOnCredentialIssuer(goCtx context.Context, ms *msgServer, constraints did.VerificationRelationships, issuerDID, signer string, update func(issuer *credentials.CredentialIssuer) error) (err error) {
//
//	k := ms.Keeper
//	ctx := sdk.UnwrapSDKContext(goCtx)
//	k.Logger(ctx).Info("request to update a did document", "target did", issuerDID)
//	// check that the did is not a key did
//	if strings.HasPrefix(issuerDID, did.DidKeyPrefix) {
//		err = sdkerrors.Wrapf(did.ErrInvalidInput, "did documents having id with key format are read only %s", issuerDID)
//		k.Logger(ctx).Error(err.Error())
//		return
//	}
//	// get the did document
//	didDoc, found := ms.DIDs.GetDidDocument(ctx, []byte(issuerDID))
//	if !found {
//		err = sdkerrors.Wrapf(did.ErrDidDocumentNotFound, "did document at %s not found", issuerDID)
//		k.Logger(ctx).Error(err.Error())
//		return
//	}
//
//	// Any verification method in the authentication relationship can update the DID document
//	if !didDoc.HasRelationship(did.NewBlockchainAccountID(ctx.ChainID(), signer), constraints...) {
//		// check also the controllers
//		signerDID := did.NewKeyDID(signer)
//		if !didDoc.HasController(signerDID) {
//			// if also the controller was not set the error
//			err = sdkerrors.Wrapf(
//				did.ErrUnauthorized,
//				"signer account %s not authorized to update the target did document at %s",
//				signer, issuerDID,
//			)
//			k.Logger(ctx).Error(err.Error())
//			return
//		}
//	}
//
//	issuer, found := k.GetCredentialIssuer(ctx, []byte(issuerDID))
//	if !found {
//		err = sdkerrors.Wrapf(credentials.ErrCredentialIssuerNotFound, "credential issuer definition for %s not found", issuerDID)
//		k.Logger(ctx).Error(err.Error())
//		return
//	}
//
//	// apply the update
//	err = update(&issuer)
//	if err != nil {
//		k.Logger(ctx).Error(err.Error())
//		return
//	}
//
//	// persist the did document
//	k.SetCredentialIssuer(ctx, issuer)
//	k.Logger(ctx).Info("credential issuer updated", "did", issuerDID, "controller", signer)
//
//	//TODO: fire the event
//	//if err := ctx.EventManager().EmitTypedEvent(credentials.NewCredentialIssuerUpdated(did, signer)); err != nil {
//	//	k.Logger(ctx).Error("failed to emit DidDocumentUpdatedEvent", "did", did, "signer", signer, "err", err)
//	//}
//	k.Logger(ctx).Info("request to update did document success", "did", didDoc.Id)
//	return
//}
