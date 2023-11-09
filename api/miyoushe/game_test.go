package miyoushe

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/codec/json"
	"github.com/starudream/go-lib/core/v2/utils/testutil"
)

func TestListGame(t *testing.T) {
	data, err := ListGame()
	testutil.LogNoErr(t, err, data, json.MustMarshalString(data))
}

func TestListGameRole(t *testing.T) {
	t.Run("all", func(t *testing.T) {
		data, err := ListGameRole("", GetAccount(t))
		testutil.LogNoErr(t, err, data)
	})

	t.Run(GameBizHKRPGCN, func(t *testing.T) {
		data, err := ListGameRole(GameBizHKRPGCN, GetAccount(t))
		testutil.LogNoErr(t, err, data)
	})
}
