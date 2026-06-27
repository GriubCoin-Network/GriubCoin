package ante_test

import (
	"strings"
	"testing"

	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	tmtime "github.com/cometbft/cometbft/types/time"
	sdk "github.com/cosmos/cosmos-sdk/types"
	evmtypes "github.com/evmos/ethermint/x/evm/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/GriubCoin-Network/GriubCoin/app"
	"github.com/GriubCoin-Network/GriubCoin/app/ante"
)

func mustParseDecCoins(value string) sdk.DecCoins {
	coins, err := sdk.ParseDecCoins(strings.ReplaceAll(value, ";", ","))
	if err != nil {
		panic(err)
	}

	return coins
}

func TestEvmMinGasFilter(t *testing.T) {
	tApp := app.NewTestApp()
	handler := ante.NewEvmMinGasFilter(tApp.GetEvmKeeper())

	ctx := tApp.NewContext(true, tmproto.Header{Height: 1, Time: tmtime.Now()})
	tApp.GetEvmKeeper().SetParams(ctx, evmtypes.Params{
		EvmDenom: "agrbc",
	})

	testCases := []struct {
		name                 string
		minGasPrices         sdk.DecCoins
		expectedMinGasPrices sdk.DecCoins
	}{
		{
			"no min gas prices",
			mustParseDecCoins(""),
			mustParseDecCoins(""),
		},
		{
			"zero ugrbc gas price",
			mustParseDecCoins("0ugrbc"),
			mustParseDecCoins("0ugrbc"),
		},
		{
			"non-zero ugrbc gas price",
			mustParseDecCoins("0.001ugrbc"),
			mustParseDecCoins("0.001ugrbc"),
		},
		{
			"zero ugrbc gas price, min agrbc price",
			mustParseDecCoins("0ugrbc;100000agrbc"),
			mustParseDecCoins("0ugrbc"), // agrbc is removed
		},
		{
			"zero ugrbc gas price, min agrbc price, other token",
			mustParseDecCoins("0ugrbc;100000agrbc;0.001other"),
			mustParseDecCoins("0ugrbc;0.001other"), // agrbc is removed
		},
		{
			"non-zero ugrbc gas price, min agrbc price",
			mustParseDecCoins("0.25ugrbc;100000agrbc;0.001other"),
			mustParseDecCoins("0.25ugrbc;0.001other"), // agrbc is removed
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := tApp.NewContext(true, tmproto.Header{Height: 1, Time: tmtime.Now()})

			ctx = ctx.WithMinGasPrices(tc.minGasPrices)
			mmd := MockAnteHandler{}

			_, err := handler.AnteHandle(ctx, nil, false, mmd.AnteHandle)
			require.NoError(t, err)
			require.True(t, mmd.WasCalled)

			assert.NoError(t, mmd.CalledCtx.MinGasPrices().Validate())
			assert.Equal(t, tc.expectedMinGasPrices, mmd.CalledCtx.MinGasPrices())
		})
	}
}
