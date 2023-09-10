package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"market/x/market/types"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				ApplesList: []types.Apples{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				ApplesCount: 2,
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated apples",
			genState: &types.GenesisState{
				ApplesList: []types.Apples{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid apples count",
			genState: &types.GenesisState{
				ApplesList: []types.Apples{
					{
						Id: 1,
					},
				},
				ApplesCount: 0,
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
