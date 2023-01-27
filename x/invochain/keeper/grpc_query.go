package keeper

import (
	"invochain/x/invochain/types"
)

var _ types.QueryServer = Keeper{}
