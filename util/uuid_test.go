package util

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"
)

func TestUUID(t *testing.T) {
	testutil.Log(t, UUID())
}
