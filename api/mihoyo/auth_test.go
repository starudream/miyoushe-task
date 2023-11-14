package mihoyo

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"

	"github.com/starudream/miyoushe-task/config"
	"github.com/starudream/miyoushe-task/util"
)

func TestGenQRCode(t *testing.T) {
	data, err := GenQRCode(config.C().FirstAccount())
	testutil.LogNoErr(t, err, data, data.Url, util.QRCode(data.Url))
}

func TestQueryQRCode(t *testing.T) {
	data, err := QueryQRCode("654b81644b7cf30567d6d20c", config.C().FirstAccount())
	testutil.LogNoErr(t, err, data)
}
