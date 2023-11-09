package miyoushe

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"
)

func TestGetBBSHome(t *testing.T) {
	t.Run(GameIdYS, func(t *testing.T) {
		data, err := GetBBSHome(GameIdYS)
		testutil.LogNoErr(t, err, data, data.GetLunaActId())
	})

	t.Run(GameIdSR, func(t *testing.T) {
		data, err := GetBBSHome(GameIdSR)
		testutil.LogNoErr(t, err, data, data.GetLunaActId())
	})
}

func TestSignBBS(t *testing.T) {
	data, err := SignBBS(GameIdSR, GetAccount(t))
	if IsRetCode(err, RetCodeBBSHasSigned) {
		t.Skip("bbs has signed")
	}
	testutil.LogNoErr(t, err, data)
}
