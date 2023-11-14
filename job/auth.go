package job

import (
	"fmt"

	"github.com/starudream/miyoushe-task/api/mihoyo"
	"github.com/starudream/miyoushe-task/api/miyoushe"
	"github.com/starudream/miyoushe-task/config"
)

func Refresh(account config.Account) (config.Account, error) {
	_, err := miyoushe.GetUser("", account)
	if err != nil {
		return RefreshSToken(account)
	}
	return account, nil
}

func RefreshSToken(account config.Account) (config.Account, error) {
	res, err := mihoyo.GetSTokenByGToken(account)
	if err != nil {
		return account, fmt.Errorf("get stoken error: %w", err)
	}
	account.Mid = res.UserInfo.Mid
	account.SToken = res.Token.Token

	config.UpdateAccount(account.Phone, func(config.Account) config.Account { return account })
	err = config.Save()
	if err != nil {
		return account, fmt.Errorf("save account error: %w", err)
	}

	return account, nil
}

func RefreshCToken(account config.Account) (config.Account, error) {
	res, err := mihoyo.GetCTokenBySToken(account)
	if err != nil {
		return account, fmt.Errorf("get ctoken error: %w", err)
	}
	account.CToken = res.CookieToken

	config.UpdateAccount(account.Phone, func(config.Account) config.Account { return account })
	err = config.Save()
	if err != nil {
		return account, fmt.Errorf("save account error: %w", err)
	}

	return account, nil
}
