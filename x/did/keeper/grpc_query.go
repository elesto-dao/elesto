package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/elesto-dao/elesto/x/did"
)

var _ did.QueryServer = Keeper{}

// DidDocument implements the DidDocument gRPC method, it querys the store and returns
// a did document to a gRPC client
func (k Keeper) DidDocument(
	c context.Context,
	req *did.QueryDidDocumentRequest,
) (*did.QueryDidDocumentResponse, error) {
	if did.IsEmpty(req.Id) {
		return nil, status.Error(codes.InvalidArgument, "did document id cannot be empty")
	}

	ctx := sdk.UnwrapSDKContext(c)
	doc, err := k.ResolveDid(ctx, did.DID(req.Id))
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	return &did.QueryDidDocumentResponse{
		DidDocument: doc,
	}, nil
}
