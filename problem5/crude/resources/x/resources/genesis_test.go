package resources_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "resources/testutil/keeper"
	"resources/testutil/nullify"
	"resources/x/resources"
	"resources/x/resources/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ResourcesKeeper(t)
	resources.InitGenesis(ctx, *k, genesisState)
	got := resources.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
