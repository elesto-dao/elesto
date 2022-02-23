package keeper

import (
	"context"
	"fmt"
	"strings"

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

func (k msgServer) RegisterCredentialIssuer(
	goCtx context.Context,
	msg *credentials.MsgRegisterCredentialIssuerRequest,
) (*credentials.MsgRegisterCredentialIssuerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	k.Logger(ctx).Info("request to register a CredentialIssuer", "did", msg.Issuer.Did)
	// check that the did is not a key did
	if strings.HasPrefix(msg.Issuer.Did, did.DidKeyPrefix) {
		err := sdkerrors.Wrapf(credentials.ErrInvalidInput, "did documents having id with key format cannot be used for credential issuers %s", msg.Issuer.Did)
		k.Logger(ctx).Error(err.Error())
		return nil, err
	}
	// retrieve the did document
	didDoc, found := k.DIDs.GetDidDocument(ctx, []byte(msg.Issuer.Did))
	if !found {
		err := sdkerrors.Wrapf(did.ErrDidDocumentNotFound, "the issuer did document was not found %s", msg.Issuer.Did)
		k.Logger(ctx).Error(err.Error())
		return nil, err
	}
	// check if the signer can modify the did document
	if !didDoc.HasRelationship(did.NewBlockchainAccountID(ctx.ChainID(), msg.Signer), did.Authentication) {
		// if also the controller was not set the error
		err := sdkerrors.Wrapf(
			credentials.ErrUnauthorized,
			"signer account %s not authorized to update the target did document at %s",
			msg.Signer, msg.Issuer.Did,
		)
		k.Logger(ctx).Error(err.Error())
		return nil, err
	}
	// create a new service for the revocation service url
	revocationCheckService := did.NewService(
		fmt.Sprint(msg.Issuer.Did, "/revocations"),
		"RevocationServiceURL",
		msg.RevocationServiceURL,
	)
	if err := didDoc.AddServices(revocationCheckService); err != nil {
		k.Logger(ctx).Error(err.Error())
		return nil, err
	}
	// now save the credential issuer
	// TODO: implement validity checks
	k.SetCredentialIssuer(ctx, *msg.Issuer)
	k.DIDs.SetDidDocument(ctx, []byte(msg.Issuer.Did), didDoc)

	// log and emit the events
	k.Logger(ctx).Info("created credential issuer", "did", msg.Issuer.Did, "controller", msg.Signer)
	if err := ctx.EventManager().EmitTypedEvents(
		did.NewDidDocumentCreatedEvent(msg.Issuer.Did, msg.Signer),
		credentials.NewCredentialIssuerRegisteredEvent(msg.Issuer.Did),
	); err != nil {
		k.Logger(ctx).Error("failed to emit events", "did", msg.Issuer.Did, "signer", msg.Signer, "err", err)
	}

	return &credentials.MsgRegisterCredentialIssuerResponse{}, nil
}

func (k msgServer) UpdateRevocationList(
	goCtx context.Context,
	msg *credentials.MsgUpdateRevocationListRequest,
) (*credentials.MsgUpdateRevocationListResponse, error) {

	if err := executeOnCredentialIssuer(
		goCtx, &k,
		did.VerificationRelationships{did.Authentication},
		msg.IssuerDid, msg.Signer,
		func(issuer *credentials.CredentialIssuer) error {
			issuer.Revocations = msg.Revocation
			return nil
		}); err != nil {
		return nil, err
	}
	return &credentials.MsgUpdateRevocationListResponse{}, nil
}

func (k msgServer) AddCredentialIssuance(
	goCtx context.Context,
	msg *credentials.MsgAddCredentialIssuanceRequest,
) (*credentials.MsgAddCredentialIssuanceResponse, error) {

	return nil, fmt.Errorf("not implemented")
}

func (k msgServer) RemoveCredentialIssuance(
	goCtx context.Context,
	msg *credentials.MsgRemoveCredentialIssuanceRequest,
) (*credentials.MsgRemoveCredentialIssuanceResponse, error) {

	return nil, fmt.Errorf("not implemented")
}

func (k msgServer) AddCredentialConstraint(
	goCtx context.Context,
	msg *credentials.MsgAddCredentialConstraintRequest,
) (*credentials.MsgAddCredentialConstraintResponse, error) {

	return nil, fmt.Errorf("not implemented")
}

func (k msgServer) RemoveCredentialConstraint(
	goCtx context.Context,
	msg *credentials.MsgRemoveCredentialConstraintRequest,
) (*credentials.MsgRemoveCredentialConstraintResponse, error) {

	return nil, fmt.Errorf("not implemented")
}

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

	return nil, fmt.Errorf("not implemented")
}

func executeOnCredentialIssuer(goCtx context.Context, ms *msgServer, constraints did.VerificationRelationships, issuerDID, signer string, update func(issuer *credentials.CredentialIssuer) error) (err error) {

	k := ms.Keeper
	ctx := sdk.UnwrapSDKContext(goCtx)
	k.Logger(ctx).Info("request to update a did document", "target did", issuerDID)
	// check that the did is not a key did
	if strings.HasPrefix(issuerDID, did.DidKeyPrefix) {
		err = sdkerrors.Wrapf(credentials.ErrInvalidInput, "did documents having id with key format are read only %s", issuerDID)
		k.Logger(ctx).Error(err.Error())
		return
	}
	// get the did document
	didDoc, found := ms.DIDs.GetDidDocument(ctx, []byte(issuerDID))
	if !found {
		err = sdkerrors.Wrapf(did.ErrDidDocumentNotFound, "did document at %s not found", issuerDID)
		k.Logger(ctx).Error(err.Error())
		return
	}

	// Any verification method in the authentication relationship can update the DID document
	if !didDoc.HasRelationship(did.NewBlockchainAccountID(ctx.ChainID(), signer), constraints...) {
		// check also the controllers
		signerDID := did.NewKeyDID(signer)
		if !didDoc.HasController(signerDID) {
			// if also the controller was not set the error
			err = sdkerrors.Wrapf(
				credentials.ErrUnauthorized,
				"signer account %s not authorized to update the target did document at %s",
				signer, issuerDID,
			)
			k.Logger(ctx).Error(err.Error())
			return
		}
	}

	issuer, found := k.GetCredentialIssuer(ctx, []byte(issuerDID))
	if !found {
		err = sdkerrors.Wrapf(credentials.ErrCredentialIssuerNotFound, "credential issuer definition for %s not found", issuerDID)
		k.Logger(ctx).Error(err.Error())
		return
	}

	// apply the update
	err = update(&issuer)
	if err != nil {
		k.Logger(ctx).Error(err.Error())
		return
	}

	// persist the did document
	k.SetCredentialIssuer(ctx, issuer)
	k.Logger(ctx).Info("credential issuer updated", "did", issuerDID, "controller", signer)

	//TODO: fire the event
	//if err := ctx.EventManager().EmitTypedEvent(credentials.NewCredentialIssuerUpdated(did, signer)); err != nil {
	//	k.Logger(ctx).Error("failed to emit DidDocumentUpdatedEvent", "did", did, "signer", signer, "err", err)
	//}
	k.Logger(ctx).Info("request to update did document success", "did", didDoc.Id)
	return
}
