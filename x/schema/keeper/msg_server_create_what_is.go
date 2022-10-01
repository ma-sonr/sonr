package keeper

import (
	"context"

    "github.com/sonr-io/sonr/x/schema/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)


func (k msgServer) CreateWhatIs(goCtx context.Context,  msg *types.MsgCreateWhatIs) (*types.MsgCreateWhatIsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

    // TODO: Handling the message
    _ = ctx

	return &types.MsgCreateWhatIsResponse{}, nil
}
