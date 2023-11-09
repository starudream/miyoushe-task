package config

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"
)

func TestAccount(t *testing.T) {
	phone := "123"

	account := Account{Phone: phone, Device: NewDevice()}
	testutil.Log(t, account)

	AddAccount(account)

	account, ok := GetAccount(phone)
	testutil.Equal(t, true, ok)
	testutil.Log(t, account)

	testutil.Nil(t, Save())
}

func TestSave(t *testing.T) {
	testutil.Nil(t, Save())
}
