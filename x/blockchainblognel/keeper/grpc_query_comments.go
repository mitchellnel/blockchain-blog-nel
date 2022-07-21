package keeper

import (
	"context"

	"blockchain-blog-nel/x/blockchainblognel/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Comments(
	goCtx context.Context,
	req *types.QueryCommentsRequest,
) (*types.QueryCommentsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// define a variable that will store a list of Comments
	var comments []*types.Comment

	// get the key-value module store using the store key
	store := ctx.KVStore(k.storeKey)

	// get the part of the store that keeps Comments, using the comment key
	commentsStore := prefix.NewStore(store, types.KeyPrefix(types.CommentKey))

	// check if the Post we want to see the Comments on exists
	post, found := k.GetPost(ctx, req.PostID)
	if !found {
		return nil, sdkerrors.Wrapf(
			types.ErrPostDoesNotExist,
			"Post with ID %d does not exist",
			req.PostID,
		)
	}
	postID := post.Id

	// paginate the comments store based on PageRequest
	pageRes, err := query.Paginate(
		commentsStore,
		req.Pagination,
		func(key []byte, value []byte) error {
			var comment types.Comment
			if err := k.cdc.Unmarshal(value, &post); err != nil {
				return err
			}

			// only show Comment if it's on the Post for which we want the comments shown
			if comment.PostID == postID {
				comments = append(comments, &comment)
			}

			return nil
		},
	)

	// throw an error if pagination fails
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// return a struct containg a list of Comments and pagination info
	return &types.QueryCommentsResponse{Comments: comments, Pagination: pageRes}, nil
}
