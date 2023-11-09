package util

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"
)

func TestQRCode(t *testing.T) {
	testutil.Log(t, "\n"+QRCode("hello world"))
}
