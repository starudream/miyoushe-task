package job

import (
	"fmt"
	"time"

	"github.com/starudream/go-lib/core/v2/slog"
	"github.com/starudream/go-lib/core/v2/utils/signalutil"

	"github.com/starudream/miyoushe-task/api/miyoushe"
	"github.com/starudream/miyoushe-task/config"
)

func WaitQRCodeConfirmed(ticket string, account config.Account) (config.Account, error) {
	ticker := time.NewTicker(3 * time.Second)

	for {
		select {
		case <-signalutil.Done():
			return account, nil

		case <-ticker.C:
			res2, err := miyoushe.QueryQRCode(ticket, account)
			if err != nil {
				if miyoushe.IsRetCode(err, miyoushe.RetCodeQRCodeExpired) {
					return account, fmt.Errorf("qrcode expired, please try again")
				}
				return account, fmt.Errorf("query qrcode error: %w", err)
			}

			switch {
			case res2.Stat.IsInit():
				slog.Info("qrcode not scanned")
				continue
			case res2.Stat.IsScanned():
				slog.Info("qrcode scanned, please confirm login")
				continue
			case res2.Stat.IsConfirmed():
				slog.Info("qrcode confirmed, login success")
			default:
				return account, fmt.Errorf("unknown qrcode stat: %s", res2.Stat)
			}

			account.Uid = res2.Payload.Uid
			account.GToken = res2.Payload.Token

			config.UpdateAccount(account.Phone, func(config.Account) config.Account {
				return account
			})
			err = config.Save()
			if err != nil {
				return account, fmt.Errorf("save account error: %w", err)
			}

			return account, nil
		}
	}
}

func Refresh(account config.Account) (_ config.Account, err error) {
	account, err = RefreshSToken(account)
	if err != nil {
		return
	}
	account, err = RefreshCToken(account)
	if err != nil {
		return
	}
	return account, nil
}

func RefreshSToken(account config.Account) (config.Account, error) {
	res, err := miyoushe.GetSTokenByGToken(account)
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
	res, err := miyoushe.GetCTokenBySToken(account)
	if err != nil {
		return account, fmt.Errorf("get stoken error: %w", err)
	}
	account.CToken = res.CookieToken

	config.UpdateAccount(account.Phone, func(config.Account) config.Account { return account })
	err = config.Save()
	if err != nil {
		return account, fmt.Errorf("save account error: %w", err)
	}

	return account, nil
}
