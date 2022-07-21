package keeper

import (
	"context"
	"fmt"

	"blockchain-blog-nel/x/blockchainblognel/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateComment(
	goCtx context.Context,
	msg *types.MsgCreateComment,
) (*types.MsgCreateCommentResponse, error) {
	// get the context
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check if the Post being commented on exists
	post, found := k.GetPost(ctx, msg.PostID)
	if !found {
		return nil, sdkerrors.Wrap(
			sdkerrors.ErrKeyNotFound,
			fmt.Sprintf("Post with ID %d (the key) doesn't exist", msg.Id),
		)
	}

	// create a variable of type Comment
	var comment = types.Comment{
		Creator:   msg.Creator,
		Id:        msg.Id,
		Title:     msg.Title,
		Body:      msg.Body,
		PostID:    msg.PostID,
		CreatedAt: ctx.BlockHeight(),
	}

	// check if the comment is older than the post
	// if more than 100 blocks, return an error
	if comment.CreatedAt > post.CreatedAt+100 {
		return nil, sdkerrors.Wrapf(
			types.ErrCommentOld,
			"Comment created at %d is more than 100 blocks older than Post created at %d",
			comment.CreatedAt,
			post.CreatedAt,
		)
	}

	id := k.AppendComment(ctx, comment)

	return &types.MsgCreateCommentResponse{Id: id}, nil
}
