package miyoushe

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"

	"github.com/starudream/miyoushe-task/api/common"
	"github.com/starudream/miyoushe-task/config"
)

func TestGetHome(t *testing.T) {
	t.Run(common.GameIdDBY, func(t *testing.T) {
		data, err := GetHome(common.GameIdDBY, config.C().FirstAccount())
		testutil.LogNoErr(t, err, data)
	})

	t.Run(common.GameIdYS, func(t *testing.T) {
		data, err := GetHome(common.GameIdYS, config.C().FirstAccount())
		testutil.LogNoErr(t, err, data, data.GetSignActId())
	})

	t.Run(common.GameIdSR, func(t *testing.T) {
		data, err := GetHome(common.GameIdSR, config.C().FirstAccount())
		testutil.LogNoErr(t, err, data, data.GetSignActId())
	})

	t.Run(common.GameIdBH3, func(t *testing.T) {
		data, err := GetHome(common.GameIdBH3, config.C().FirstAccount())
		testutil.LogNoErr(t, err, data, data.GetSignActId())
	})
}

func TestGetBusinesses(t *testing.T) {
	data, err := GetBusinesses(config.C().FirstAccount())
	testutil.LogNoErr(t, err, data)
}
