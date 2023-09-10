package market

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"market/testutil/sample"
	marketsimulation "market/x/market/simulation"
	"market/x/market/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = marketsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateApples = "op_weight_msg_apples"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateApples int = 100

	opWeightMsgUpdateApples = "op_weight_msg_apples"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateApples int = 100

	opWeightMsgDeleteApples = "op_weight_msg_apples"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteApples int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	marketGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		ApplesList: []types.Apples{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		ApplesCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&marketGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateApples int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateApples, &weightMsgCreateApples, nil,
		func(_ *rand.Rand) {
			weightMsgCreateApples = defaultWeightMsgCreateApples
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateApples,
		marketsimulation.SimulateMsgCreateApples(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateApples int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateApples, &weightMsgUpdateApples, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateApples = defaultWeightMsgUpdateApples
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateApples,
		marketsimulation.SimulateMsgUpdateApples(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteApples int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteApples, &weightMsgDeleteApples, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteApples = defaultWeightMsgDeleteApples
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteApples,
		marketsimulation.SimulateMsgDeleteApples(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
