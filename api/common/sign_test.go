package common

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"
)

func TestDS1(t *testing.T) {
	c, s := ds1(1699372800, "123456")
	testutil.Equal(t, "salt=QVu5OdwEWxkq9ygpYBgDprR5tI471HWQ&t=1699372800&r=123456", c)
	testutil.Equal(t, "1699372800,123456,6e98ea06bfae57644337e04d818301db", s)
}

func TestDS2(t *testing.T) {
	c, s := ds2(1699372800, 100001, "", "role_id=222681079&server=cn_gf01")
	testutil.Equal(t, "salt=t0qEgfub6cvueAPgR5m9aQWWVciEer7v&t=1699372800&r=100001&b=&q=role_id=222681079&server=cn_gf01", c)
	testutil.Equal(t, "1699372800,100001,f73a2b0996a7439f11f62b53a44d8f7c", s)
}
