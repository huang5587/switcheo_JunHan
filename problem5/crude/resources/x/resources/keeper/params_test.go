package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "resources/testutil/keeper"
	"resources/x/resources/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.ResourcesKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
