package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/elesto-dao/elesto/v3/x/credential"
)

func NewPublicCredentialProposalHandler(k Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *credential.ProposePublicCredentialID:
			return handlePublicProposalChange(ctx, k, c)

		case *credential.ProposeRemovePublicCredentialID:
			return handleRemovePublicProposalChange(ctx, k, c)

		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized credential proposal content type: %T", c)
		}
	}
}

func handlePublicProposalChange(ctx sdk.Context, k Keeper, proposal *credential.ProposePublicCredentialID) error {
	if _, found := k.GetCredentialDefinition(ctx, proposal.CredentialDefinitionID); !found {
		return fmt.Errorf("proposal with %s id not found", proposal.CredentialDefinitionID)
	}

	if k.IsPublicCredentialDefinitionAllowed(ctx, proposal.CredentialDefinitionID) {
		return fmt.Errorf("credential definition with id %s is already public", proposal.CredentialDefinitionID)
	}

	k.AllowPublicCredential(ctx, proposal.CredentialDefinitionID)
	return nil
}

func handleRemovePublicProposalChange(ctx sdk.Context, k Keeper, proposal *credential.ProposeRemovePublicCredentialID) error {
	if _, found := k.GetCredentialDefinition(ctx, proposal.CredentialDefinitionID); !found {
		return fmt.Errorf("proposal with %s id not found", proposal.CredentialDefinitionID)
	}

	if !k.IsPublicCredentialDefinitionAllowed(ctx, proposal.CredentialDefinitionID) {
		return fmt.Errorf("credential definition with id %s is not public", proposal.CredentialDefinitionID)
	}

	k.RemovePublicCredentialFromAllowedList(ctx, proposal.CredentialDefinitionID)
	return nil
}
