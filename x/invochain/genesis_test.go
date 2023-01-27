package invochain_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "invochain/testutil/keeper"
	"invochain/testutil/nullify"
	"invochain/x/invochain"
	"invochain/x/invochain/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.InvochainKeeper(t)
	invochain.InitGenesis(ctx, *k, genesisState)
	got := invochain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
