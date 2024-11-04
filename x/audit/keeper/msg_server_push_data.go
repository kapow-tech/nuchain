package keeper

import (
	"context"

	"nuchain/x/audit/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) PushData(goCtx context.Context, msg *types.MsgPushData) (*types.MsgPushDataResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgPushDataResponse{}, nil
}
