package pricefeed

import (
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/GriubCoin-Network/GriubCoin/x/pricefeed/keeper"
	"github.com/GriubCoin-Network/GriubCoin/x/pricefeed/types"
)

// EndBlocker updates the current pricefeed
func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyEndBlocker)

	k.SetCurrentPricesForAllMarkets(ctx)
}
