package miyoushe

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"

	"github.com/starudream/miyoushe-task/config"
)

func TestCreateVerification(t *testing.T) {
	data, err := CreateVerification(config.C().FirstAccount())
	testutil.LogNoErr(t, err, data)
}

func TestVerifyVerification(t *testing.T) {
	data, err := VerifyVerification("", "", config.C().FirstAccount())
	testutil.LogNoErr(t, err, data)
}
