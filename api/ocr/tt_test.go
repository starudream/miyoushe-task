package ocr

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"

	"github.com/starudream/miyoushe-task/api/common"
	"github.com/starudream/miyoushe-task/config"
)

func TestTT(t *testing.T) {
	data, err := TT(config.TT().Key, "x", "x", common.RefererAct)
	testutil.LogNoErr(t, err, data)
}

func TestTTResult(t *testing.T) {
	data, err := ttResult(config.TT().Key, "x")
	testutil.LogNoErr(t, err, data)
}
