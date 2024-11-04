package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "nuchain/testutil/keeper"
	"nuchain/x/audit/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.AuditKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
