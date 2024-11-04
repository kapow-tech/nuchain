package keeper

import (
	"nuchain/x/audit/types"
)

var _ types.QueryServer = Keeper{}
