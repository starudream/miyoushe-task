package job

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"

	"github.com/starudream/miyoushe-task/api/common"
	"github.com/starudream/miyoushe-task/api/miyoushe"
)

func TestSignGameRecords_Success(t *testing.T) {
	rs := SignGameRecords{
		{
			GameId:    common.GameIdYS,
			RoleName:  "游戏昵称1",
			RoleUid:   "123456789",
			HasSigned: true,
			IsRisky:   true,
			Award:     "A*5, B*5",
		},
		{
			GameId:    common.GameIdSR,
			RoleName:  "游戏昵称2",
			RoleUid:   "123456789",
			HasSigned: true,
			IsRisky:   false,
			Award:     "ABC*500",
		},
	}
	for i := range rs {
		rs[i].GameName = miyoushe.AllGamesById[rs[i].GameId].Name
	}
	testutil.Log(t, rs, rs.Success())
}
