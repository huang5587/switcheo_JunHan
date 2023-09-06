package keeper

import (
	"fruits/x/fruits/types"
)

var _ types.QueryServer = Keeper{}
