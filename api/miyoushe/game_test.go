package miyoushe

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/codec/json"
	"github.com/starudream/go-lib/core/v2/utils/testutil"

	"github.com/starudream/miyoushe-task/api/common"
	"github.com/starudream/miyoushe-task/config"
)

func TestListGame(t *testing.T) {
	data, err := ListGame()
	testutil.LogNoErr(t, err, data)
	testutil.Equal(t, json.MustMarshalString(AllGames), json.MustMarshalString(data.List))
}

func TestListGameRole(t *testing.T) {
	t.Run("all", func(t *testing.T) {
		data, err := ListGameRole("", config.C().FirstAccount())
		testutil.LogNoErr(t, err, data)
	})

	t.Run(common.GameBizSRCN, func(t *testing.T) {
		data, err := ListGameRole(common.GameBizSRCN, config.C().FirstAccount())
		testutil.LogNoErr(t, err, data)
	})
}

func TestListGameCard(t *testing.T) {
	data, err := ListGameCard(config.C().FirstAccount())
	testutil.LogNoErr(t, err, data)
}
