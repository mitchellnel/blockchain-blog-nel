package keeper

import (
	"context"

	"blockchain-blog-nel/x/blockchainblognel/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) DeleteComment(
	goCtx context.Context,
	msg *types.MsgDeleteComment,
) (*types.MsgDeleteCommentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check if the Post that the message sender says the Comment-to-delete is on exists
	post, exists := k.GetPost(ctx, msg.PostID)
	if !exists {
		return nil, sdkerrors.Wrapf(
			types.ErrPostDoesNotExist,
			"Post with ID %d does not exist",
			msg.PostID,
		)
	}

	// check if the Comment that we want to delete exists
	comment, exists := k.GetComment(ctx, msg.CommentID)
	if !exists {
		return nil, sdkerrors.Wrapf(
			types.ErrCommentDoesNotExists,
			"Comment with ID %d does not exist",
			msg.CommentID,
		)
	}

	// check if the Post specified in the message is the Post that the Comment is on
	if msg.PostID != comment.PostID {
		return nil, sdkerrors.Wrapf(
			types.ErrCommentNotFoundOnPost,
			"Comment with ID %d cannot be found on the Post with ID %d",
			msg.CommentID,
			msg.PostID,
		)
	}

	// use Ignite generated CRUD function to delete the comment
	k.RemoveComment(ctx, msg.CommentID)

	_ = post

	return &types.MsgDeleteCommentResponse{}, nil
}
