package keeper

import (
	"blockchain-blog-nel/x/blockchainblognel/types"
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) AppendPost(ctx sdk.Context, post types.Post) uint64 {
	// get the current number of posts in the store
	postCount := k.GetPostCount(ctx)

	// assign an ID to the post based on the current number of posts in the store
	post.Id = postCount

	// get the store
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.PostKey))

	// convert the post ID into bytes
	postID_bytes := make([]byte, 8)
	binary.BigEndian.PutUint64(postID_bytes, post.Id)

	// marshal the post into bytes
	post_bytes := k.cdc.MustMarshal(&post)

	// insert the post bytes into the store using postID as the key
	store.Set(postID_bytes, post_bytes)

	// update the post count
	k.SetPostCount(ctx, postCount+1)

	return post.Id
}

func (k Keeper) GetPostCount(ctx sdk.Context) uint64 {
	// get the store using storeKey (which is "blockchainblognel") and PostCountKey (which is
	// "Post-count-")
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.PostCountKey))

	// convert the PostCountKey to bytes
	byteKey := []byte(types.PostCountKey)

	// get the value of the count
	postCount_bytes := store.Get(byteKey)
	if postCount_bytes == nil {
		// return 0 if postCount_bytes is nil, i.e. no posts created yet
		return 0
	}

	// convert the count into a uint64, and return
	return binary.BigEndian.Uint64(postCount_bytes)
}

func (k Keeper) SetPostCount(ctx sdk.Context, postCount uint64) {
	// get the store using storeKey (which is "blockchainblognel") and PostCountKey (which is
	// "Post-count-")
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.PostCountKey))

	// convert the PostCountKey to bytes
	byteKey := []byte(types.PostCountKey)

	// convert postCount from uint64 to bytes
	postCount_bytes := make([]byte, 8)
	binary.BigEndian.PutUint64(postCount_bytes, postCount)

	// set the value of Post-count- to postCount
	store.Set(byteKey, postCount_bytes)
}
