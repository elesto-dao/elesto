package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/elesto-dao/elesto/v4/x/mint/types"
)

var _ types.QueryServer = Keeper{}

// Params returns params of the mint module.
func (k Keeper) Params(c context.Context, _ *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	params := k.GetParams(ctx)

	return &types.QueryParamsResponse{Params: params}, nil
}

func (k Keeper) Inflation(_ context.Context, req *types.QueryInflationRequest) (*types.QueryInflationResponse, error) {
	inflation, err := types.GetInflation(req.GetHeight()).Float64()
	return &types.QueryInflationResponse{
		Epoch:     types.GetEpoch(req.GetHeight()),
		Inflation: inflation,
	}, err
}
