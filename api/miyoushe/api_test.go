package miyoushe

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"

	"github.com/starudream/miyoushe-task/config"
)

func GetAccount(t *testing.T) config.Account {
	accounts := config.C().Accounts
	if len(accounts) == 0 {
		t.SkipNow()
	}
	account := accounts[0]
	testutil.Log(t, account)
	return account
}

func SaveAccount(t *testing.T, account config.Account) {
	config.UpdateAccount(account.Phone, func(config.Account) config.Account { return account })
	testutil.Nil(t, config.Save())
}

func TestGameById(t *testing.T) {
	testutil.Log(t, GameById)
}

func TestDS1(t *testing.T) {
	c, s := ds1(1699372800, "123456")
	testutil.Equal(t, "salt=pIlzNr5SAZhdnFW8ZxauW8UlxRdZc45r&t=1699372800&r=123456", c)
	testutil.Equal(t, "1699372800,123456,1338659bf2f390470717646a658d15e1", s)
}

func TestDS2(t *testing.T) {
	c, s := ds2(1699372800, 100001, "", "role_id=222681079&server=cn_gf01")
	testutil.Equal(t, "salt=t0qEgfub6cvueAPgR5m9aQWWVciEer7v&t=1699372800&r=100001&b=&q=role_id=222681079&server=cn_gf01", c)
	testutil.Equal(t, "1699372800,100001,f73a2b0996a7439f11f62b53a44d8f7c", s)
}
