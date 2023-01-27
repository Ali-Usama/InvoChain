package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "invochain/testutil/keeper"
	"invochain/x/invochain/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.InvochainKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
