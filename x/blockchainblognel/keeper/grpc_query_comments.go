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

	// get context
	ctx := sdk.UnwrapSDKContext(goCtx)

	// define a variable that will store the list of Comments
	var comments []*types.Comment

	// get the store that keeps comments, using the comment key
	commentsStore := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.CommentKey))

	// check if the post for which we are retrieving the comments exists
	post, found := k.GetPost(ctx, req.Id)
	if !found {
		return nil, sdkerrors.Wrapf(
			types.ErrPostDoesNotExist,
			"Post with ID %v does not exist",
			req.Id,
		)
	}
	postID := post.Id

	// paginate the comments store based on PageRequest
	pageRes, err := query.Paginate(
		commentsStore,
		req.Pagination,
		func(key []byte, value []byte) error {
			var comment types.Comment
			if err := k.cdc.Unmarshal(value, &comment); err != nil {
				return err
			}

			if comment.PostID == postID {
				comments = append(comments, &comment)
			}

			return nil
		},
	)

	// throw an error if pagination failed
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryCommentsResponse{Post: &post, Comments: comments, Pagination: pageRes}, nil
}
