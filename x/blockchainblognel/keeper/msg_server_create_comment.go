package keeper

import (
	"context"

	"blockchain-blog-nel/x/blockchainblognel/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateComment(goCtx context.Context, msg *types.MsgCreateComment) (*types.MsgCreateCommentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateCommentResponse{}, nil
}
