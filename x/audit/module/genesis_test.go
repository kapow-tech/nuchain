package audit_test

import (
	"testing"

	keepertest "nuchain/testutil/keeper"
	"nuchain/testutil/nullify"
	audit "nuchain/x/audit/module"
	"nuchain/x/audit/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.AuditKeeper(t)
	audit.InitGenesis(ctx, k, genesisState)
	got := audit.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
