package keeper

import (
	"context"

	"blockchain-blog-nel/x/blockchainblognel/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Posts(
	goCtx context.Context,
	req *types.QueryPostsRequest,
) (*types.QueryPostsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	// get context
	ctx := sdk.UnwrapSDKContext(goCtx)

	// define a variable that will store a list of Posts
	var posts []*types.Post

	// get the key-value module store using the store key (in this case it's "chain")
	store := ctx.KVStore(k.storeKey)

	// get the part of the store that stores keeps Posts using the PostKey (in this case it's
	// "Post-value-")
	postsStore := prefix.NewStore(store, []byte(types.PostKey))

	// paginate the postsStore based on PageRequest
	pageRes, err := query.Paginate(
		postsStore,
		req.Pagination,
		func(key []byte, value []byte) error {
			var post types.Post
			if err := k.cdc.Unmarshal(value, &post); err != nil {
				return err
			}

			posts = append(posts, &post)
			return nil
		},
	)
	if err != nil {
		// throw an error if pagination failed
		return nil, status.Error(codes.Internal, err.Error())
	}

	// return a struct containing a list of posts and pagination info
	return &types.QueryPostsResponse{Post: posts, Pagination: pageRes}, nil
}
