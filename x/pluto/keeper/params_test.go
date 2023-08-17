package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "pluto/testutil/keeper"
	"pluto/x/pluto/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.PlutoKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
