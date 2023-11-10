package ocr

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"

	"github.com/starudream/miyoushe-task/api/miyoushe"
	"github.com/starudream/miyoushe-task/config"
)

func TestRR(t *testing.T) {
	data, err := RR(config.C().RROCRKey, "x", "x", miyoushe.RefererAct)
	testutil.LogNoErr(t, err, data)
}
