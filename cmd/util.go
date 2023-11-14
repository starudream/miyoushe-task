package main

import (
	"fmt"

	"github.com/starudream/go-lib/core/v2/utils/osutil"
	"github.com/starudream/go-lib/core/v2/utils/sliceutil"

	"github.com/starudream/miyoushe-task/config"
)

func xGetAccount(args []string, i ...int) config.Account {
	if len(i) == 0 {
		i = []int{0}
	}
	phone, _ := sliceutil.GetValue(args, i[0])
	if phone == "" {
		osutil.ExitErr(fmt.Errorf("requires account phone"))
	}
	account, exists := config.GetAccount(phone)
	if !exists {
		osutil.ExitErr(fmt.Errorf("account %s not exists", phone))
	}
	return account
}
