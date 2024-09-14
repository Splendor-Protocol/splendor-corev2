package splendor_test

import (
	"testing"

	keepertest "splendor/testutil/keeper"
	"splendor/testutil/nullify"
	splendor "splendor/x/splendor/module"
	"splendor/x/splendor/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SplendorKeeper(t)
	splendor.InitGenesis(ctx, k, genesisState)
	got := splendor.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
