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

// RevocationList implements the DidDocument gRPC method
func (k Keeper) RevocationList(
	c context.Context,
	req *credentials.QueryRevocationListRequest,
) (*credentials.QueryRevocationListResponse, error) {

	if !did.IsValidDID(req.Id) {
		return nil, status.Error(codes.InvalidArgument, "did document id cannot be empty")
	}

	ctx := sdk.UnwrapSDKContext(c)
	ci, found := k.GetCredentialIssuer(ctx, []byte(req.Id))
	if !found {
		return nil, status.Error(codes.NotFound, "credential issuer not found")
	}

	return &credentials.QueryRevocationListResponse{RevocationList: ci.Revocations}, nil
}

func (k Keeper) CredentialIssuer(
	c context.Context,
	req *credentials.QueryCredentialIssuerRequest,
) (*credentials.QueryCredentialIssuerResponse, error) {

	if !did.IsValidDID(req.Id) {
		return nil, status.Error(codes.InvalidArgument, "did document id cannot be empty")
	}

	ctx := sdk.UnwrapSDKContext(c)
	ci, found := k.GetCredentialIssuer(ctx, []byte(req.Id))
	if !found {
		return nil, status.Error(codes.NotFound, "credential issuer not found")
	}

	return &credentials.QueryCredentialIssuerResponse{Issuer: &ci}, nil

}

func (k Keeper) PublicCredentials(
	c context.Context,
	req *credentials.QueryPublicCredentialsRequest,
) (*credentials.QueryPublicCredentialsResponse, error) {
	return nil, fmt.Errorf("not implemented")
}