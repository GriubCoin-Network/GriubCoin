package community_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/GriubCoin-Network/GriubCoin/x/community"
	"github.com/GriubCoin-Network/GriubCoin/x/community/keeper"
	"github.com/GriubCoin-Network/GriubCoin/x/community/testutil"
	"github.com/stretchr/testify/suite"
)

func TestABCIPayoutAccumulatedStakingRewards(t *testing.T) {
	testFunc := func(ctx sdk.Context, k keeper.Keeper) {
		community.BeginBlocker(ctx, k)
	}
	suite.Run(t, testutil.NewStakingRewardsTestSuite(testFunc))
}
