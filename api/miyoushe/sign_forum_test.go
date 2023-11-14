package miyoushe

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"

	"github.com/starudream/miyoushe-task/api/common"
	"github.com/starudream/miyoushe-task/config"
)

func TestSignForum(t *testing.T) {
	data, err := SignForum(common.GameIdSR, config.C().FirstAccount(), nil)
	if common.IsRetCode(err, common.RetCodeForumHasSigned) {
		t.Skip("bbs has signed")
	}
	testutil.LogNoErr(t, err, data)
}

func TestGetSignForum(t *testing.T) {
	data, err := GetSignForum(common.GameIdSR, config.C().FirstAccount())
	testutil.LogNoErr(t, err, data)
}
