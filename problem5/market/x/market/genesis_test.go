package market_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "market/testutil/keeper"
	"market/testutil/nullify"
	"market/x/market"
	"market/x/market/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		ApplesList: []types.Apples{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		ApplesCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MarketKeeper(t)
	market.InitGenesis(ctx, *k, genesisState)
	got := market.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.ApplesList, got.ApplesList)
	require.Equal(t, genesisState.ApplesCount, got.ApplesCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
