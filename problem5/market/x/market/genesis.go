package market

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"market/x/market/keeper"
	"market/x/market/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the apples
	for _, elem := range genState.ApplesList {
		k.SetApples(ctx, elem)
	}

	// Set apples count
	k.SetApplesCount(ctx, genState.ApplesCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.ApplesList = k.GetAllApples(ctx)
	genesis.ApplesCount = k.GetApplesCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
