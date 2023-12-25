package miyoushe

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"

	"github.com/starudream/miyoushe-task/api/common"
	"github.com/starudream/miyoushe-task/config"
)

func GetSR(t *testing.T) (string, string, string, string) {
	data1, err := GetHome(common.GameIdYS)
	testutil.LogNoErr(t, err, data1)
	actId := data1.GetSignActId()
	testutil.NotEqual(t, "", actId)

	data2, err := ListGameRole(common.GameBizYSCN, config.C().FirstAccount())
	testutil.LogNoErr(t, err, data2)
	testutil.NotEqual(t, 0, len(data2.List))
	region, uid := data2.List[0].Region, data2.List[0].GameUid

	return common.GameNameYS, actId, region, uid
}

func TestSignGame(t *testing.T) {
	gameName, actId, region, uid := GetSR(t)
	data, err := SignGame(gameName, actId, region, uid, config.C().FirstAccount(), nil)
	if common.IsRetCode(err, common.RetCodeGameHasSigned) {
		t.Skip("game has signed")
	}
	testutil.LogNoErr(t, err, data)
	testutil.Equal(t, false, data.IsRisky())
}

func TestGetSignGame(t *testing.T) {
	gameName, actId, region, uid := GetSR(t)
	data, err := GetSignGame(gameName, actId, region, uid, config.C().FirstAccount())
	testutil.LogNoErr(t, err, data)
}

func TestListSignGame(t *testing.T) {
	gameName, actId, _, _ := GetSR(t)
	data, err := ListSignGame(gameName, actId, config.C().FirstAccount())
	testutil.LogNoErr(t, err, data)
}

func TestListSignGameAward(t *testing.T) {
	gameName, actId, region, uid := GetSR(t)
	data, err := ListSignGameAward(gameName, actId, region, uid, config.C().FirstAccount())
	testutil.LogNoErr(t, err, data, len(data), data.Today(), data.Today().ShortString())
}
