package mihoyo

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"

	"github.com/starudream/miyoushe-task/config"
)

func TestGetSTokenByGToken(t *testing.T) {
	data, err := GetSTokenByGToken(config.C().FirstAccount())
	testutil.LogNoErr(t, err, data)
}

func TestGetCTokenBySToken(t *testing.T) {
	data, err := GetCTokenBySToken(config.C().FirstAccount())
	testutil.LogNoErr(t, err, data)
}

func TestGetLTokenBySToken(t *testing.T) {
	data, err := GetLTokenBySToken(config.C().FirstAccount())
	testutil.LogNoErr(t, err, data)
}
