package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "splendor/testutil/keeper"
	"splendor/x/splendor/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.SplendorKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
