package miyoushe

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"
)

func TestGetUser(t *testing.T) {
	t.Run("no-auth", func(t *testing.T) {
		data, err := GetUser("75596302")
		testutil.LogNoErr(t, err, data)
	})

	t.Run("by-auth", func(t *testing.T) {
		data, err := GetUser("", GetAccount(t))
		testutil.LogNoErr(t, err, data)
	})
}
