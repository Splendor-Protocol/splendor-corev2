package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "earth/testutil/keeper"
	"earth/x/earth/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.EarthKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
