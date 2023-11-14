package job

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"

	"github.com/starudream/miyoushe-task/api/common"
	"github.com/starudream/miyoushe-task/api/miyoushe"
)

func TestSignForumRecord_Success(t *testing.T) {
	r := SignForumRecord{
		GameId:     common.GameIdDBY,
		HasSigned:  false,
		IsRisky:    true,
		Points:     50,
		PostView:   PostView,
		PostUpvote: PostUpvote,
		PostShare:  PostShare,
	}
	r.GameName = miyoushe.AllGamesById[r.GameId].Name
	testutil.Log(t, r, r.Success())
}
