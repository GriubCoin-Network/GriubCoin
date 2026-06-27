package hard

import (
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/GriubCoin-Network/GriubCoin/x/hard/keeper"
	"github.com/GriubCoin-Network/GriubCoin/x/hard/types"
)

// BeginBlocker updates interest rates
func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	k.ApplyInterestRateUpdates(ctx)
}
