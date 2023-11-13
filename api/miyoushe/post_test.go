package miyoushe

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"
)

func TestListPost(t *testing.T) {
	data, err := ListPost(ForumIdSR, "", GetAccount(t))
	testutil.LogNoErr(t, err, data, data.LastId)
}

func TestGetPost(t *testing.T) {
	data, err := GetPost("45453046", GetAccount(t))
	testutil.LogNoErr(t, err, data, data.Post.IsUpvote())
}

func TestUpvotePost(t *testing.T) {
	err := UpvotePost("45453046", GetAccount(t))
	testutil.LogNoErr(t, err)
}

func TestSharePost(t *testing.T) {
	data, err := SharePost("45453046", GetAccount(t))
	testutil.LogNoErr(t, err, data)
}
