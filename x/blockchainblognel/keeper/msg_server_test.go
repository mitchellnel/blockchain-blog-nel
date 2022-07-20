package keeper_test

import (
	"context"
	"testing"

	keepertest "blockchain-blog-nel/testutil/keeper"
	"blockchain-blog-nel/x/blockchainblognel/keeper"
	"blockchain-blog-nel/x/blockchainblognel/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.BlockchainblognelKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
