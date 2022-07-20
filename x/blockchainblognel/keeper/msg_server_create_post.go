package keeper

import (
	"context"

	"blockchain-blog-nel/x/blockchainblognel/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreatePost(
	goCtx context.Context,
	msg *types.MsgCreatePost,
) (*types.MsgCreatePostResponse, error) {
	// get the context
	ctx := sdk.UnwrapSDKContext(goCtx)

	// create a variable of type Post
	var post = types.Post{
		Creator: msg.Creator,
		Title:   msg.Title,
		Body:    msg.Body,
	}

	// add a post to the store and get back an ID
	id := k.AppendPost(ctx, post)

	// return the ID of the post
	return &types.MsgCreatePostResponse{Id: id}, nil
}
