package keeper_test

import (
	"context"
	"testing"

	keepertest "fruits/testutil/keeper"
	"fruits/x/fruits/keeper"
	"fruits/x/fruits/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.FruitsKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
