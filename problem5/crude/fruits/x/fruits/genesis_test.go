package fruits_test

import (
	"testing"

	keepertest "fruits/testutil/keeper"
	"fruits/testutil/nullify"
	"fruits/x/fruits"
	"fruits/x/fruits/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.FruitsKeeper(t)
	fruits.InitGenesis(ctx, *k, genesisState)
	got := fruits.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
