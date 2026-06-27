package types_test

import (
	"testing"

	"github.com/GriubCoin-Network/GriubCoin/x/swap/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
)

func TestKeys(t *testing.T) {
	key := types.PoolKey(types.PoolID("ugrbc", "usdx"))
	assert.Equal(t, types.PoolID("ugrbc", "usdx"), string(key))

	key = types.DepositorPoolSharesKey(sdk.AccAddress("testaddress1"), types.PoolID("ugrbc", "usdx"))
	assert.Equal(t, string(sdk.AccAddress("testaddress1"))+"|"+types.PoolID("ugrbc", "usdx"), string(key))
}
