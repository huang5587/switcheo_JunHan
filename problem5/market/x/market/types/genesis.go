package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		ApplesList: []Apples{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in apples
	applesIdMap := make(map[uint64]bool)
	applesCount := gs.GetApplesCount()
	for _, elem := range gs.ApplesList {
		if _, ok := applesIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for apples")
		}
		if elem.Id >= applesCount {
			return fmt.Errorf("apples id should be lower or equal than the last id")
		}
		applesIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
