package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/elesto-dao/elesto/x/credentials"
	"github.com/elesto-dao/elesto/x/did"
)

var _ credentials.QueryServer = Keeper{}

func (k Keeper) CredentialDefinition(
	c context.Context,
	req *credentials.QueryCredentialDefinitionRequest,
) (*credentials.QueryCredentialDefinitionResponse, error) {

	if !did.IsValidDID(req.Did) {
		return nil, status.Error(codes.InvalidArgument, "did document id cannot be empty")
	}

	ctx := sdk.UnwrapSDKContext(c)

	cd, found := k.GetCredentialDefinition(ctx, req.Did)
	if !found {
		return nil, status.Error(codes.NotFound, "credential definition not found")
	}

	return &credentials.QueryCredentialDefinitionResponse{Definition: &cd}, nil
}

func (k Keeper) CredentialDefinitionsByPublisher(
	c context.Context,
	req *credentials.QueryCredentialDefinitionsByPublisherRequest,
) (*credentials.QueryCredentialDefinitionsByPublisherResponse, error) {

	return nil, fmt.Errorf("not implemented")
}

func (k Keeper) CredentialDefinitions(
	c context.Context,
	req *credentials.QueryCredentialDefinitionsRequest,
) (*credentials.QueryCredentialDefinitionsResponse, error) {

	return nil, fmt.Errorf("not implemented")
}

func (k Keeper) PublicCredential(
	c context.Context,
	req *credentials.QueryPublicCredentialRequest,
) (*credentials.QueryPublicCredentialResponse, error) {

	ctx := sdk.UnwrapSDKContext(c)

	pc, found := k.GetPublicCredential(ctx, req.Id)
	if !found {
		return nil, status.Error(codes.NotFound, "credential definition not found")
	}

	return &credentials.QueryPublicCredentialResponse{Credential: &pc}, nil

}

func (k Keeper) PublicCredentialsByHolder(
	c context.Context,
	req *credentials.QueryPublicCredentialsByHolderRequest,
) (*credentials.QueryPublicCredentialsByHolderResponse, error) {
	return nil, fmt.Errorf("not implemented")
}

func (k Keeper) PublicCredentialsByIssuer(
	c context.Context,
	req *credentials.QueryPublicCredentialsByIssuerRequest,
) (*credentials.QueryPublicCredentialsByIssuerResponse, error) {
	return nil, fmt.Errorf("not implemented")
}

func (k Keeper) PublicCredentials(
	c context.Context,
	req *credentials.QueryPublicCredentialsRequest,
) (*credentials.QueryPublicCredentialsResponse, error) {
	return nil, fmt.Errorf("not implemented")
}
