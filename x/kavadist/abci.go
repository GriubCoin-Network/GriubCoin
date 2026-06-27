package kavadist

import (
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/GriubCoin-Network/GriubCoin/x/kavadist/keeper"
	"github.com/GriubCoin-Network/GriubCoin/x/kavadist/types"
)

func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	err := k.MintPeriodInflation(ctx)
	if err != nil {
		panic(err)
	}
}
