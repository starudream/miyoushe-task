package main

import (
	"encoding/base64"
	"fmt"

	"github.com/starudream/go-lib/cobra/v2"
	"github.com/starudream/go-lib/core/v2/slog"
	"github.com/starudream/go-lib/core/v2/utils/fmtutil"
	"github.com/starudream/go-lib/core/v2/utils/sliceutil"
	"github.com/starudream/go-lib/tablew/v2"

	"github.com/starudream/miyoushe-task/api/common"
	"github.com/starudream/miyoushe-task/api/mihoyo"
	"github.com/starudream/miyoushe-task/config"
	"github.com/starudream/miyoushe-task/job"
)

var (
	accountCmd = cobra.NewCommand(func(c *cobra.Command) {
		c.Use = "account"
		c.Short = "Manage accounts"
	})

	accountInitCmd = cobra.NewCommand(func(c *cobra.Command) {
		c.Use = "init <account phone>"
		c.Short = "Init account device information"
		c.Args = func(cmd *cobra.Command, args []string) error {
			phone, _ := sliceutil.GetValue(args, 0)
			if phone == "" {
				return fmt.Errorf("requires account phone")
			}
			_, exists := config.GetAccount(phone)
			if exists {
				return fmt.Errorf("account %s already exists", phone)
			}
			return nil
		}
		c.RunE = func(cmd *cobra.Command, args []string) error {
			phone, _ := sliceutil.GetValue(args, 0)
			config.AddAccount(config.Account{Phone: phone, Device: config.NewDevice()})
			return config.Save()
		}
	})

	accountLoginCmd = cobra.NewCommand(func(c *cobra.Command) {
		c.Use = "login <account phone>"
		c.Short = "Login account"
		c.RunE = func(cmd *cobra.Command, args []string) error {
			account := xGetAccount(args)

			_, err := mihoyo.SendPhoneCode("", account)
			if err != nil {
				if !common.IsRetCode(err, common.RetCodeSendPhoneCodeFrequently) {
					return fmt.Errorf("send phone code error: %w", err)
				}

				aigis, aigisData := common.GetAigisData(err)
				if aigis == nil || aigisData == nil {
					return fmt.Errorf("get aigis data empty")
				}

				slog.Info("aigis gt: %s, challenge: %s", aigisData.Gt, aigisData.Challenge)

				geetest := base64.StdEncoding.EncodeToString([]byte(fmtutil.Scan("please enter GeeTest json string: ")))

				_, err = mihoyo.SendPhoneCode(fmt.Sprintf("%s;%s", aigis.SessionId, geetest), account)
				if err != nil {
					return fmt.Errorf("send phone code error: %w", err)
				}
			}

			code := fmtutil.Scan("please enter the verification code you received (use ctrl+c to exit): ")
			if code == "" {
				return nil
			}

			res, err := mihoyo.LoginByPhoneCode(code, account)
			if err != nil {
				return fmt.Errorf("login by phone code error: %w", err)
			}

			account.Uid = res.UserInfo.Aid
			account.Mid = res.UserInfo.Mid
			account.SToken = res.Token.Token

			_, err = job.RefreshCToken(account)
			if err != nil {
				return fmt.Errorf("refresh token error: %w", err)
			}

			slog.Info("login success")
			return nil
		}
	})

	accountListCmd = cobra.NewCommand(func(c *cobra.Command) {
		c.Use = "list"
		c.Short = "List accounts"
		c.Run = func(cmd *cobra.Command, args []string) {
			fmt.Println(tablew.Structs(config.C().Accounts))
		}
	})
)

func init() {
	accountCmd.AddCommand(accountInitCmd)
	accountCmd.AddCommand(accountLoginCmd)
	accountCmd.AddCommand(accountListCmd)

	rootCmd.AddCommand(accountCmd)
}
