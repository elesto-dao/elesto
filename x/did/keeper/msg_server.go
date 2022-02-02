package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/elesto-dao/elesto/x/did"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the identity MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) did.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ did.MsgServer = msgServer{}

// CreateDidDocument creates a new DID document
func (k msgServer) CreateDidDocument(
	goCtx context.Context,
	msg *did.MsgCreateDidDocument,
) (*did.MsgCreateDidDocumentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	k.Logger(ctx).Info("request to create a did document", "target did", msg.Id)
	// setup a new did document (performs input validation)
	didDoc, err := did.NewDidDocument(msg.Id,
		did.WithServices(msg.Services...),
		did.WithVerifications(msg.Verifications...),
		did.WithControllers(msg.Controllers...),
	)
	if err != nil {
		k.Logger(ctx).Error(err.Error())
		return nil, err
	}

	// check that the did is not already taken
	_, found := k.Keeper.GetDidDocument(ctx, []byte(msg.Id))
	if found {
		err := sdkerrors.Wrapf(did.ErrDidDocumentFound, "a document with did %s already exists", msg.Id)
		k.Logger(ctx).Error(err.Error())
		return nil, err
	}

	// persist the did document
	k.Keeper.SetDidDocument(ctx, []byte(msg.Id), didDoc)

	// now create and persist the metadata
	didM := did.NewDidMetadata(ctx.TxBytes(), ctx.BlockTime())
	k.Keeper.SetDidMetadata(ctx, []byte(msg.Id), didM)

	k.Logger(ctx).Info("created did document", "did", msg.Id, "controller", msg.Signer)

	// emit the event
	if err := ctx.EventManager().EmitTypedEvents(did.NewDidDocumentCreatedEvent(msg.Id, msg.Signer)); err != nil {
		k.Logger(ctx).Error("failed to emit DidDocumentCreatedEvent", "did", msg.Id, "signer", msg.Signer, "err", err)
	}

	return &did.MsgCreateDidDocumentResponse{}, nil
}

// UpdateDidDocument update an existing DID document
func (k msgServer) UpdateDidDocument(
	goCtx context.Context,
	msg *did.MsgUpdateDidDocument,
) (*did.MsgUpdateDidDocumentResponse, error) {

	if err := executeOnDidWithRelationships(
		goCtx, &k.Keeper,
		newConstraints(did.Authentication),
		msg.Doc.Id, msg.Signer,
		//XXX: check this assignment during audit
		//nolint
		func(didDoc *did.DidDocument) error {
			if !did.IsValidDIDDocument(msg.Doc) {
				return sdkerrors.Wrapf(did.ErrInvalidDIDFormat, "invalid did document")
			}
			didDoc = msg.Doc
			return nil
		}); err != nil {
		return nil, err
	}
	return &did.MsgUpdateDidDocumentResponse{}, nil
}

// AddVerification adds a verification method and it's relationships to a DID Document
func (k msgServer) AddVerification(
	goCtx context.Context,
	msg *did.MsgAddVerification,
) (*did.MsgAddVerificationResponse, error) {

	if err := executeOnDidWithRelationships(
		goCtx, &k.Keeper,
		newConstraints(did.Authentication),
		msg.Id, msg.Signer,
		func(didDoc *did.DidDocument) error {
			return didDoc.AddVerifications(msg.Verification)
		}); err != nil {
		return nil, err
	}
	return &did.MsgAddVerificationResponse{}, nil
}

// AddService adds a service to an existing DID document
func (k msgServer) AddService(
	goCtx context.Context,
	msg *did.MsgAddService,
) (*did.MsgAddServiceResponse, error) {

	if err := executeOnDidWithRelationships(
		goCtx, &k.Keeper,
		newConstraints(did.Authentication),
		msg.Id, msg.Signer,
		func(didDoc *did.DidDocument) error {
			return didDoc.AddServices(msg.ServiceData)
		}); err != nil {
		return nil, err
	}

	return &did.MsgAddServiceResponse{}, nil
}

// RevokeVerification removes a public key and controller from an existing DID document
func (k msgServer) RevokeVerification(
	goCtx context.Context,
	msg *did.MsgRevokeVerification,
) (*did.MsgRevokeVerificationResponse, error) {

	if err := executeOnDidWithRelationships(
		goCtx, &k.Keeper,
		newConstraints(did.Authentication),
		msg.Id, msg.Signer,
		func(didDoc *did.DidDocument) error {
			return didDoc.RevokeVerification(msg.MethodId)
		}); err != nil {
		return nil, err
	}

	return &did.MsgRevokeVerificationResponse{}, nil
}

// DeleteService removes a service from an existing DID document
func (k msgServer) DeleteService(
	goCtx context.Context,
	msg *did.MsgDeleteService,
) (*did.MsgDeleteServiceResponse, error) {

	if err := executeOnDidWithRelationships(
		goCtx, &k.Keeper,
		newConstraints(did.Authentication),
		msg.Id, msg.Signer,
		func(didDoc *did.DidDocument) error {
			// Only try to remove service if there are services
			if len(didDoc.Service) == 0 {
				return sdkerrors.Wrapf(did.ErrInvalidState, "the did document doesn't have services associated")
			}
			// delete service
			didDoc.DeleteService(msg.ServiceId)
			return nil
		}); err != nil {
		return nil, err
	}

	return &did.MsgDeleteServiceResponse{}, nil
}

// SetVerificationRelationships set the verification relationships for an existing DID document
func (k msgServer) SetVerificationRelationships(
	goCtx context.Context,
	msg *did.MsgSetVerificationRelationships,
) (*did.MsgSetVerificationRelationshipsResponse, error) {

	if err := executeOnDidWithRelationships(
		goCtx, &k.Keeper,
		newConstraints(did.Authentication),
		msg.Id, msg.Signer,
		func(didDoc *did.DidDocument) error {
			return didDoc.SetVerificationRelationships(msg.MethodId, msg.Relationships...)
		}); err != nil {
		return nil, err
	}

	return &did.MsgSetVerificationRelationshipsResponse{}, nil
}

// AddController add a new controller to a DID
func (k msgServer) AddController(
	goCtx context.Context,
	msg *did.MsgAddController,
) (*did.MsgAddControllerResponse, error) {
	if err := executeOnDidWithRelationships(
		goCtx, &k.Keeper,
		newConstraints(did.Authentication),
		msg.Id, msg.Signer,
		func(didDoc *did.DidDocument) error {
			return didDoc.AddControllers(msg.ControllerDid)
		}); err != nil {
		return nil, err
	}

	return &did.MsgAddControllerResponse{}, nil
}

// DeleteController remove an existing controller from a DID document
func (k msgServer) DeleteController(
	goCtx context.Context,
	msg *did.MsgDeleteController,
) (*did.MsgDeleteControllerResponse, error) {

	if err := executeOnDidWithRelationships(
		goCtx, &k.Keeper,
		newConstraints(did.Authentication),
		msg.Id, msg.Signer, func(didDoc *did.DidDocument) error {
			return didDoc.DeleteControllers(msg.ControllerDid)
		}); err != nil {
		return nil, err
	}
	return &did.MsgDeleteControllerResponse{}, nil
}

// helper function to update the did metadata
func updateDidMetadata(keeper *Keeper, ctx sdk.Context, didID string) (err error) {
	didMeta, found := keeper.GetDidMetadata(ctx, []byte(didID))
	if found {
		did.UpdateDidMetadata(&didMeta, ctx.TxBytes(), ctx.BlockTime())
		keeper.SetDidMetadata(ctx, []byte(didID), didMeta)
	} else {
		err = fmt.Errorf("(warning) did metadata not found")
	}
	return
}

// VerificationRelationships for did document manipulation
type VerificationRelationships []string

func newConstraints(relationships ...string) VerificationRelationships {
	return relationships
}

func executeOnDidWithRelationships(goCtx context.Context, k *Keeper, constraints VerificationRelationships, didID, signer string, update func(document *did.DidDocument) error) (err error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	k.Logger(ctx).Info("request to update a did document", "target did", didID)
	// TODO: fail if the input did is of type KEY (immutable)
	// eg: ErrInvalidState, "did document key is immutable"

	// get the did document
	didDoc, found := k.GetDidDocument(ctx, []byte(didID))
	if !found {
		err = sdkerrors.Wrapf(did.ErrDidDocumentNotFound, "did document at %s not found", didID)
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
				did.ErrUnauthorized,
				"signer account %s not authorized to update the target did document at %s",
				signer, didID,
			)
			k.Logger(ctx).Error(err.Error())
			return
		}
	}

	// apply the update
	err = update(&didDoc)
	if err != nil {
		k.Logger(ctx).Error(err.Error())
		return
	}

	// persist the did document
	k.SetDidDocument(ctx, []byte(didID), didDoc)
	k.Logger(ctx).Info("Set verification relationship from did document for", "did", didID, "controller", signer)

	// update the Metadata
	if err = updateDidMetadata(k, ctx, didDoc.Id); err != nil {
		k.Logger(ctx).Error(err.Error(), "did", didDoc.Id)
		return
	}
	// fire the event
	if err := ctx.EventManager().EmitTypedEvent(did.NewDidDocumentUpdatedEvent(didID, signer)); err != nil {
		k.Logger(ctx).Error("failed to emit DidDocumentUpdatedEvent", "did", didID, "signer", signer, "err", err)
	}
	k.Logger(ctx).Info("request to update did document success", "did", didDoc.Id)
	return
}
