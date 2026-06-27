package pricefeed_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/GriubCoin-Network/GriubCoin/x/pricefeed"
	"github.com/GriubCoin-Network/GriubCoin/x/pricefeed/keeper"
	"github.com/GriubCoin-Network/GriubCoin/x/pricefeed/testutil"
)

func TestEndBlocker_UpdatesMultipleMarkets(t *testing.T) {
	testutil.SetCurrentPrices_PriceCalculations(t, func(ctx sdk.Context, keeper keeper.Keeper) {
		pricefeed.EndBlocker(ctx, keeper)
	})

	testutil.SetCurrentPrices_EventEmission(t, func(ctx sdk.Context, keeper keeper.Keeper) {
		pricefeed.EndBlocker(ctx, keeper)
	})
}
