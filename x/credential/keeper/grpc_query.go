package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/elesto-dao/elesto/v2/x/credential"
	"github.com/elesto-dao/elesto/v2/x/did"
)

var _ credential.QueryServer = Keeper{}

func (k Keeper) CredentialDefinition(
	c context.Context,
	req *credential.QueryCredentialDefinitionRequest,
) (*credential.QueryCredentialDefinitionResponse, error) {

	if credential.IsEmpty(req.Id) {
		return nil, status.Error(codes.InvalidArgument, "credential definition id must not be empty")
	}

	ctx := sdk.UnwrapSDKContext(c)

	cd, found := k.GetCredentialDefinition(ctx, req.Id)
	if !found {
		return nil, status.Error(codes.NotFound, "credential definition not found")
	}

	return &credential.QueryCredentialDefinitionResponse{Definition: &cd}, nil
}

func (k Keeper) CredentialDefinitionsByPublisher(
	c context.Context,
	req *credential.QueryCredentialDefinitionsByPublisherRequest,
) (*credential.QueryCredentialDefinitionsByPublisherResponse, error) {
	if !did.IsValidDID(req.Did) {
		return nil, status.Error(codes.InvalidArgument, "publisher DID must be a valid DID")
	}
	ctx := sdk.UnwrapSDKContext(c)
	cds := k.GetCredentialDefinitionsWithFilter(ctx, func(cd *credential.CredentialDefinition) bool {
		return cd.PublisherId == req.Did
	})
	return &credential.QueryCredentialDefinitionsByPublisherResponse{Definitions: cds}, nil
}

func (k Keeper) CredentialDefinitions(
	c context.Context,
	req *credential.QueryCredentialDefinitionsRequest,
) (*credential.QueryCredentialDefinitionsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	cd, pr, err := k.GetCredentialDefinitions(ctx, req)
	return &credential.QueryCredentialDefinitionsResponse{Definitions: cd, Pagination: pr}, err
}

func (k Keeper) PublicCredential(
	c context.Context,
	req *credential.QueryPublicCredentialRequest,
) (*credential.QueryPublicCredentialResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	pc, found := k.GetPublicCredential(ctx, req.Id)
	if !found {
		return nil, status.Error(codes.NotFound, "public verifiable credential not found")
	}
	return &credential.QueryPublicCredentialResponse{Credential: &pc}, nil

}

func (k Keeper) PublicCredentialsByHolder(
	c context.Context,
	req *credential.QueryPublicCredentialsByHolderRequest,
) (*credential.QueryPublicCredentialsByHolderResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	pvcs := k.GetPublicCredentialWithFilter(ctx, func(vc *credential.PublicVerifiableCredential) bool {
		wc, err := credential.NewWrappedCredential(vc)
		if err != nil {
			return false
		}
		subjectID, hasIt := wc.GetSubjectID()
		if !hasIt {
			return false
		}
		return subjectID == req.Did
	})

	return &credential.QueryPublicCredentialsByHolderResponse{Credential: pvcs}, nil
}

func (k Keeper) PublicCredentialsByIssuer(
	c context.Context,
	req *credential.QueryPublicCredentialsByIssuerRequest,
) (*credential.QueryPublicCredentialsByIssuerResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	pvcs := k.GetPublicCredentialWithFilter(ctx, func(vc *credential.PublicVerifiableCredential) bool {
		return vc.Issuer == req.Did
	})
	return &credential.QueryPublicCredentialsByIssuerResponse{Credential: pvcs}, nil
}

func (k Keeper) PublicCredentials(
	c context.Context,
	req *credential.QueryPublicCredentialsRequest,
) (*credential.QueryPublicCredentialsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	pvcs := k.GetPublicCredentialWithFilter(ctx, func(vc *credential.PublicVerifiableCredential) bool {
		return true
	})
	return &credential.QueryPublicCredentialsResponse{Credential: pvcs}, nil
}
