package job

import (
	"fmt"
	"time"

	"github.com/starudream/go-lib/core/v2/slog"

	"github.com/starudream/miyoushe-task/api/common"
	"github.com/starudream/miyoushe-task/api/mihoyo"
	"github.com/starudream/miyoushe-task/config"
)

func WaitQRCodeConfirmed(ticket string, account config.Account) (config.Account, error) {
	ticker := time.NewTicker(3 * time.Second)

	for {
		<-ticker.C
		res, err := mihoyo.QueryQRCode(ticket, account)
		if err != nil {
			if common.IsRetCode(err, common.RetCodeQRCodeExpired) {
				return account, fmt.Errorf("qrcode expired, please try again")
			}
			return account, fmt.Errorf("query qrcode error: %w", err)
		}

		switch {
		case res.Stat.IsInit():
			slog.Info("qrcode not scanned")
			continue
		case res.Stat.IsScanned():
			slog.Info("qrcode scanned, please confirm login")
			continue
		case res.Stat.IsConfirmed():
			slog.Info("qrcode confirmed, login success")
		default:
			return account, fmt.Errorf("unknown qrcode stat: %s", res.Stat)
		}

		account.Uid = res.Payload.Uid
		account.GToken = res.Payload.Token

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
