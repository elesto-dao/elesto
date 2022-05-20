package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/elesto-dao/elesto/x/credential"
	"github.com/elesto-dao/elesto/x/did"
)

var _ credential.QueryServer = Keeper{}

func (k Keeper) CredentialDefinition(
	c context.Context,
	req *credential.QueryCredentialDefinitionRequest,
) (*credential.QueryCredentialDefinitionResponse, error) {

	if !did.IsValidDID(req.Did) {
		return nil, status.Error(codes.InvalidArgument, "did document id cannot be empty")
	}

	ctx := sdk.UnwrapSDKContext(c)

	cd, found := k.GetCredentialDefinition(ctx, req.Did)
	if !found {
		return nil, status.Error(codes.NotFound, "credential definition not found")
	}

	return &credential.QueryCredentialDefinitionResponse{Definition: &cd}, nil
}

func (k Keeper) CredentialDefinitionsByPublisher(
	c context.Context,
	req *credential.QueryCredentialDefinitionsByPublisherRequest,
) (*credential.QueryCredentialDefinitionsByPublisherResponse, error) {

	return nil, fmt.Errorf("not implemented")
}

func (k Keeper) CredentialDefinitions(
	c context.Context,
	req *credential.QueryCredentialDefinitionsRequest,
) (*credential.QueryCredentialDefinitionsResponse, error) {

	return nil, fmt.Errorf("not implemented")
}

func (k Keeper) PublicCredential(
	c context.Context,
	req *credential.QueryPublicCredentialRequest,
) (*credential.QueryPublicCredentialResponse, error) {

	ctx := sdk.UnwrapSDKContext(c)

	pc, found := k.GetPublicCredential(ctx, req.Id)
	if !found {
		return nil, status.Error(codes.NotFound, "credential definition not found")
	}

	return &credential.QueryPublicCredentialResponse{Credential: &pc}, nil

}

func (k Keeper) PublicCredentialsByHolder(
	c context.Context,
	req *credential.QueryPublicCredentialsByHolderRequest,
) (*credential.QueryPublicCredentialsByHolderResponse, error) {
	return nil, fmt.Errorf("not implemented")
}

func (k Keeper) PublicCredentialsByIssuer(
	c context.Context,
	req *credential.QueryPublicCredentialsByIssuerRequest,
) (*credential.QueryPublicCredentialsByIssuerResponse, error) {
	return nil, fmt.Errorf("not implemented")
}

func (k Keeper) PublicCredentials(
	c context.Context,
	req *credential.QueryPublicCredentialsRequest,
) (*credential.QueryPublicCredentialsResponse, error) {
	return nil, fmt.Errorf("not implemented")
}
