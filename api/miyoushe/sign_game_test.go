package miyoushe

import (
	"strings"
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"

	"github.com/starudream/miyoushe-task/api/common"
	"github.com/starudream/miyoushe-task/config"
)

var gameIdByName = map[string]string{
	common.GameNameBH3: common.GameIdBH3,
	common.GameNameYS:  common.GameIdYS,
	common.GameNameBH2: common.GameIdBH2,
	common.GameNameWD:  common.GameIdWD,
	common.GameNameSR:  common.GameIdSR,
	common.GameNameZZZ: common.GameIdZZZ,
}

func GetRole(t *testing.T, gameBiz string) (string, string, string, string) {
	gameName := strings.Split(gameBiz, "_")[0]
	gameId := gameIdByName[gameName]

	data1, err := GetHome(gameId, config.C().FirstAccount())
	testutil.LogNoErr(t, err, data1)
	actId := data1.GetSignActId()
	testutil.MustNotEqual(t, "", actId)

	data2, err := ListGameRole(gameBiz, config.C().FirstAccount())
	testutil.LogNoErr(t, err, data2)
	testutil.MustNotEqual(t, 0, len(data2.List))
	region, uid := data2.List[0].Region, data2.List[0].GameUid

	return gameName, actId, region, uid
}

var gameBiz = common.GameBizZZZCN

func TestSignGame(t *testing.T) {
	gameName, actId, region, uid := GetRole(t, gameBiz)
	data, err := SignGame(gameName, actId, region, uid, config.C().FirstAccount(), nil)
	if common.IsRetCode(err, common.RetCodeGameHasSigned) {
		t.Skip("game has signed")
	}
	testutil.LogNoErr(t, err, data)
	testutil.Equal(t, false, data.IsRisky())
}

func TestGetSignGame(t *testing.T) {
	gameName, actId, region, uid := GetRole(t, gameBiz)
	data, err := GetSignGame(gameName, actId, region, uid, config.C().FirstAccount())
	testutil.LogNoErr(t, err, data)
}

func TestListSignGame(t *testing.T) {
	gameName, actId, _, _ := GetRole(t, gameBiz)
	data, err := ListSignGame(gameName, actId, config.C().FirstAccount())
	testutil.LogNoErr(t, err, data)
}

func TestListSignGameAward(t *testing.T) {
	gameName, actId, region, uid := GetRole(t, gameBiz)
	data, err := ListSignGameAward(gameName, actId, region, uid, config.C().FirstAccount())
	testutil.LogNoErr(t, err, data, len(data), data.Today(), data.Today().ShortString())
}
