package earth_test

import (
	"testing"

	keepertest "earth/testutil/keeper"
	"earth/testutil/nullify"
	earth "earth/x/earth/module"
	"earth/x/earth/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.EarthKeeper(t)
	earth.InitGenesis(ctx, k, genesisState)
	got := earth.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
