package keeper

import (
	"context"

	"nuchain/api/nuchain/audit"
	"nuchain/x/audit/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) PushData(goCtx context.Context, msg *types.MsgPushData) (*types.MsgPushDataResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx
	data := audit.Entry{
		Actor:     msg.Actor,
		Datatype:  msg.DataType,
		Id:        msg.Id,
		Data:      msg.Data,
		Timestamp: msg.Timestamp,
	}
	err := k.Db.EntryTable().Insert(ctx, &data)
	if err != nil {
		return nil, err
	}

	return &types.MsgPushDataResponse{}, nil
}
