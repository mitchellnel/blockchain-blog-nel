package blockchainblognel_test

import (
	"testing"

	keepertest "blockchain-blog-nel/testutil/keeper"
	"blockchain-blog-nel/testutil/nullify"
	"blockchain-blog-nel/x/blockchainblognel"
	"blockchain-blog-nel/x/blockchainblognel/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		CommentList: []types.Comment{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		CommentCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.BlockchainblognelKeeper(t)
	blockchainblognel.InitGenesis(ctx, *k, genesisState)
	got := blockchainblognel.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.CommentList, got.CommentList)
	require.Equal(t, genesisState.CommentCount, got.CommentCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
