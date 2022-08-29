package keeper

import (
	"context"
	"strings"

	"github.com/google/uuid"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	didmod "github.com/elesto-dao/elesto/v2/x/did"
)

const uuid4Length int = 36

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the identity MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) didmod.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ didmod.MsgServer = msgServer{}

// CreateDidDocument creates a new DID document
func (k msgServer) CreateDidDocument(
	goCtx context.Context,
	msg *didmod.MsgCreateDidDocument,
) (*didmod.MsgCreateDidDocumentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	k.Logger(ctx).Info("request to create a did document", "target did", msg.Id)

	// check that the did is not already taken
	found := k.Keeper.HasDidDocument(ctx, []byte(msg.Id))
	if found {
		err := sdkerrors.Wrapf(didmod.ErrDidDocumentFound, "a document with did %s already exists", msg.Id)
		k.Logger(ctx).Error(err.Error())
		return nil, err
	}

	// check that the did is not a key did
	if strings.HasPrefix(msg.Id, didmod.DidKeyPrefix) {
		err := sdkerrors.Wrapf(didmod.ErrInvalidInput, "did documents having id with key format cannot be created %s", msg.Id)
		k.Logger(ctx).Error(err.Error())
		return nil, err
	}

	// setup a new did document (performs input validation)
	didDoc, err := didmod.NewDidDocument(msg.Id,
		didmod.WithServices(msg.Services...),
		didmod.WithVerifications(msg.Verifications...),
		didmod.WithControllers(msg.Controllers...),
	)
	if err != nil {
		k.Logger(ctx).Error(err.Error())
		return nil, err
	}

	// check the last part of the DID is a valid UUID
	if !isValidUUID(didmod.DID(msg.Id)) {
		err := sdkerrors.Wrapf(didmod.ErrInvalidInput, "%s is not a valid UUID format", msg.Id)
		return nil, err
	}

	// persist the did document
	k.Keeper.SetDidDocument(ctx, []byte(msg.Id), didDoc)

	k.Logger(ctx).Info("created did document", "did", msg.Id, "controller", msg.Signer)

	// emit the event
	if err := ctx.EventManager().EmitTypedEvents(didmod.NewDidDocumentCreatedEvent(msg.Id, msg.Signer)); err != nil {
		k.Logger(ctx).Error("failed to emit DidDocumentCreatedEvent", "did", msg.Id, "signer", msg.Signer, "err", err)
	}

	return &didmod.MsgCreateDidDocumentResponse{}, nil
}

// UpdateDidDocument update an existing DID document
func (k msgServer) UpdateDidDocument(
	goCtx context.Context,
	msg *didmod.MsgUpdateDidDocument,
) (*didmod.MsgUpdateDidDocumentResponse, error) {

	if err := executeOnDidWithRelationships(
		goCtx, &k.Keeper,
		didmod.VerificationRelationships{didmod.Authentication},
		msg.Doc.Id, msg.Signer,
		//XXX: check this assignment during audit
		//nolint
		func(didDoc *didmod.DidDocument) (*didmod.DidDocument, error) {
			return msg.Doc, nil
		}); err != nil {
		return nil, err
	}
	return &didmod.MsgUpdateDidDocumentResponse{}, nil
}

// AddVerification adds a verification method and it's relationships to a DID Document
func (k msgServer) AddVerification(
	goCtx context.Context,
	msg *didmod.MsgAddVerification,
) (*didmod.MsgAddVerificationResponse, error) {

	if err := executeOnDidWithRelationships(
		goCtx, &k.Keeper,
		didmod.VerificationRelationships{didmod.Authentication},
		msg.Id, msg.Signer,
		func(didDoc *didmod.DidDocument) (*didmod.DidDocument, error) {
			err := didDoc.AddVerifications(msg.Verification)
			return didDoc, err
		}); err != nil {
		return nil, err
	}
	return &didmod.MsgAddVerificationResponse{}, nil
}

// AddService adds a service to an existing DID document
func (k msgServer) AddService(
	goCtx context.Context,
	msg *didmod.MsgAddService,
) (*didmod.MsgAddServiceResponse, error) {

	if err := executeOnDidWithRelationships(
		goCtx, &k.Keeper,
		didmod.VerificationRelationships{didmod.Authentication},
		msg.Id, msg.Signer,
		func(didDoc *didmod.DidDocument) (*didmod.DidDocument, error) {
			err := didDoc.AddServices(msg.ServiceData)
			return didDoc, err
		}); err != nil {
		return nil, err
	}

	return &didmod.MsgAddServiceResponse{}, nil
}

// RevokeVerification removes a public key and controller from an existing DID document
func (k msgServer) RevokeVerification(
	goCtx context.Context,
	msg *didmod.MsgRevokeVerification,
) (*didmod.MsgRevokeVerificationResponse, error) {

	if err := executeOnDidWithRelationships(
		goCtx, &k.Keeper,
		didmod.VerificationRelationships{didmod.Authentication},
		msg.Id, msg.Signer,
		func(didDoc *didmod.DidDocument) (*didmod.DidDocument, error) {
			err := didDoc.RevokeVerification(msg.MethodId)
			return didDoc, err
		}); err != nil {
		return nil, err
	}

	return &didmod.MsgRevokeVerificationResponse{}, nil
}

// DeleteService removes a service from an existing DID document
func (k msgServer) DeleteService(
	goCtx context.Context,
	msg *didmod.MsgDeleteService,
) (*didmod.MsgDeleteServiceResponse, error) {

	if err := executeOnDidWithRelationships(
		goCtx, &k.Keeper,
		didmod.VerificationRelationships{didmod.Authentication},
		msg.Id, msg.Signer,
		func(didDoc *didmod.DidDocument) (*didmod.DidDocument, error) {
			// Only try to remove service if there are services
			if len(didDoc.Service) == 0 {
				return nil, sdkerrors.Wrapf(didmod.ErrInvalidState, "the did document doesn't have services associated")
			}
			// delete service
			didDoc.DeleteService(msg.ServiceId)
			return didDoc, nil
		}); err != nil {
		return nil, err
	}

	return &didmod.MsgDeleteServiceResponse{}, nil
}

// SetVerificationRelationships set the verification relationships for an existing DID document
func (k msgServer) SetVerificationRelationships(
	goCtx context.Context,
	msg *didmod.MsgSetVerificationRelationships,
) (*didmod.MsgSetVerificationRelationshipsResponse, error) {

	if err := executeOnDidWithRelationships(
		goCtx, &k.Keeper,
		didmod.VerificationRelationships{didmod.Authentication},
		msg.Id, msg.Signer,
		func(didDoc *didmod.DidDocument) (*didmod.DidDocument, error) {
			err := didDoc.SetVerificationRelationships(msg.MethodId, msg.Relationships...)
			return didDoc, err
		}); err != nil {
		return nil, err
	}

	return &didmod.MsgSetVerificationRelationshipsResponse{}, nil
}

// AddController add a new controller to a DID
func (k msgServer) AddController(
	goCtx context.Context,
	msg *didmod.MsgAddController,
) (*didmod.MsgAddControllerResponse, error) {
	if err := executeOnDidWithRelationships(
		goCtx, &k.Keeper,
		didmod.VerificationRelationships{didmod.Authentication},
		msg.Id, msg.Signer,
		func(didDoc *didmod.DidDocument) (*didmod.DidDocument, error) {
			err := didDoc.AddControllers(msg.ControllerDid)
			return didDoc, err
		}); err != nil {
		return nil, err
	}

	return &didmod.MsgAddControllerResponse{}, nil
}

// DeleteController remove an existing controller from a DID document
func (k msgServer) DeleteController(
	goCtx context.Context,
	msg *didmod.MsgDeleteController,
) (*didmod.MsgDeleteControllerResponse, error) {

	if err := executeOnDidWithRelationships(
		goCtx, &k.Keeper,
		didmod.VerificationRelationships{didmod.Authentication},
		msg.Id, msg.Signer,
		func(didDoc *didmod.DidDocument) (*didmod.DidDocument, error) {
			err := didDoc.DeleteControllers(msg.ControllerDid)
			return didDoc, err
		}); err != nil {
		return nil, err
	}
	return &didmod.MsgDeleteControllerResponse{}, nil
}

// executeOnDidWithRelationships updates a did document with by the given parameters
// this function takes care of checking that a did can be updated by a given signer
func executeOnDidWithRelationships(
	goCtx context.Context,
	k *Keeper,
	constraints didmod.VerificationRelationships,
	did, signer string,
	update func(didDoc *didmod.DidDocument) (*didmod.DidDocument, error),
) (err error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	k.Logger(ctx).Info("request to update a did document", "target did", did)
	// check that the did is not a key did
	if strings.HasPrefix(did, didmod.DidKeyPrefix) {
		err = sdkerrors.Wrapf(didmod.ErrInvalidInput, "did documents having id with key format are read only %s", did)
		k.Logger(ctx).Error(err.Error())
		return
	}
	// get the did document
	didDoc, found := k.GetDidDocument(ctx, []byte(did))
	if !found {
		err = sdkerrors.Wrapf(didmod.ErrDidDocumentNotFound, "did document at %s not found", did)
		k.Logger(ctx).Error(err.Error())
		return
	}

	// Any verification method in the authentication relationship can update the DID document
	if !didDoc.HasRelationship(didmod.NewBlockchainAccountID(ctx.ChainID(), signer), constraints...) {
		// check also the controllers
		signerDID := didmod.NewKeyDID(signer)
		if !didDoc.HasController(signerDID) {
			// if also the controller was not set the error
			err = sdkerrors.Wrapf(
				didmod.ErrUnauthorized,
				"signer account %s not authorized to update the target did document at %s",
				signer, did,
			)
			k.Logger(ctx).Error(err.Error())
			return
		}
	}

	// apply the update
	updatedDidDoc, err := update(&didDoc)
	if err != nil || updatedDidDoc == nil {
		k.Logger(ctx).Error(err.Error())
		return
	}

	// check the did doc is valid before storing
	if !didmod.IsValidDIDDocument(updatedDidDoc) {
		return sdkerrors.Wrapf(didmod.ErrInvalidDIDFormat, "invalid did document")
	}

	// persist the did document
	k.SetDidDocument(ctx, []byte(did), *updatedDidDoc)
	k.Logger(ctx).Info("Set did document in the store for", "did", did, "controller", signer)

	// fire the event
	if err := ctx.EventManager().EmitTypedEvent(didmod.NewDidDocumentUpdatedEvent(did, signer)); err != nil {
		k.Logger(ctx).Error("failed to emit DidDocumentUpdatedEvent", "did", did, "signer", signer, "err", err)
	}
	k.Logger(ctx).Info("request to update did document success", "did", updatedDidDoc.Id)
	return
}

// checks if the provided DID has a UUID-conformant namespace-identifier
func isValidUUID(did didmod.DID) bool {
	if len(did) < uuid4Length {
		return false
	}
	// gets the namespace identifier
	namespaceId := did[len(did)-uuid4Length:]
	_, err := uuid.Parse(string(namespaceId))
	return err == nil
}
