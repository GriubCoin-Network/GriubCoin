package types_test

import (
	"os"
	"testing"

	"github.com/GriubCoin-Network/GriubCoin/app"
)

func TestMain(m *testing.M) {
	app.SetSDKConfig()
	os.Exit(m.Run())
}
