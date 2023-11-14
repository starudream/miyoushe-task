package miyoushe

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"

	"github.com/starudream/miyoushe-task/api/common"
	"github.com/starudream/miyoushe-task/config"
)

func TestListPost(t *testing.T) {
	data, err := ListPost(common.ForumIdSR, "", config.C().FirstAccount())
	testutil.LogNoErr(t, err, data, data.LastId)
}

func TestListFeedPost(t *testing.T) {
	data, err := ListFeedPost(common.GameIdDBY, config.C().FirstAccount())
	testutil.LogNoErr(t, err, data, data.LastId)
}

func TestGetPost(t *testing.T) {
	data, err := GetPost("45453046", config.C().FirstAccount())
	testutil.LogNoErr(t, err, data, data.Post.IsUpvote())
}

func TestUpvotePost(t *testing.T) {
	err := UpvotePost("45453046", false, config.C().FirstAccount())
	testutil.LogNoErr(t, err)
}

func TestCollectPost(t *testing.T) {
	err := CollectPost("45453046", true, config.C().FirstAccount())
	testutil.LogNoErr(t, err)
}

func TestSharePost(t *testing.T) {
	data, err := SharePost("45453046", config.C().FirstAccount())
	testutil.LogNoErr(t, err, data)
}
