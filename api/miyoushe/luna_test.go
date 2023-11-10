package miyoushe

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"
)

func GetSR(t *testing.T) (string, string, string) {
	data1, err := GetBBSHome(GameIdSR)
	testutil.LogNoErr(t, err, data1)
	actId := data1.GetLunaActId()
	testutil.NotEqual(t, "", actId)

	data, err := ListGameRole(GameBizHKRPGCN, GetAccount(t))
	testutil.LogNoErr(t, err, data)
	testutil.NotEqual(t, 0, len(data.List))
	region, uid := data.List[0].Region, data.List[0].GameUid

	return actId, region, uid
}

func TestSignLuna(t *testing.T) {
	actId, region, uid := GetSR(t)
	data, err := SignLuna(actId, region, uid, GetAccount(t), nil)
	if IsRetCode(err, RetCodeLunaHasSigned) {
		t.Skip("luna has signed")
	}
	testutil.LogNoErr(t, err, data)
	testutil.Equal(t, false, data.IsRisky())
}

func TestGetLunaToday(t *testing.T) {
	actId, region, uid := GetSR(t)
	data, err := GetLunaToday(actId, region, uid, GetAccount(t))
	testutil.LogNoErr(t, err, data)
}

func TestListLuna(t *testing.T) {
	actId, _, _ := GetSR(t)
	data, err := ListLuna(actId, GetAccount(t))
	testutil.LogNoErr(t, err, data)
}

func TestListLunaAward(t *testing.T) {
	actId, region, uid := GetSR(t)
	data, err := ListLunaAward(actId, region, uid, GetAccount(t))
	testutil.LogNoErr(t, err, data, len(data), data.Today(), data.Today().ShortString())
}
