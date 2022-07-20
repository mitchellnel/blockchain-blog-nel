package keeper_test

import (
	"testing"

	testkeeper "blockchain-blog-nel/testutil/keeper"
	"blockchain-blog-nel/x/blockchainblognel/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.BlockchainblognelKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
