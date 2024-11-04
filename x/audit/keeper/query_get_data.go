package keeper

import (
	"context"

	"nuchain/x/audit/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetData(goCtx context.Context, req *types.QueryGetDataRequest) (*types.QueryGetDataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	entry, err := k.Db.EntryTable().Get(ctx, req.Datatype, req.Id)
	if err != nil {
		return nil, err
	}

	return &types.QueryGetDataResponse{
		Datatype:  req.Datatype,
		Id:        req.Id,
		Data:      entry.Data,
		Timestamp: entry.Timestamp,
		Actor:     entry.Actor,
	}, nil
}
