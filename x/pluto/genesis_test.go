package pluto_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "pluto/testutil/keeper"
	"pluto/testutil/nullify"
	"pluto/x/pluto"
	"pluto/x/pluto/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.PlutoKeeper(t)
	pluto.InitGenesis(ctx, *k, genesisState)
	got := pluto.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
