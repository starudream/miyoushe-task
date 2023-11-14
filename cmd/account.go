package main

import (
	"fmt"

	"github.com/starudream/go-lib/cobra/v2"
	"github.com/starudream/go-lib/core/v2/slog"
	"github.com/starudream/go-lib/core/v2/utils/sliceutil"
	"github.com/starudream/go-lib/tablew/v2"

	"github.com/starudream/miyoushe-task/api/mihoyo"
	"github.com/starudream/miyoushe-task/config"
	"github.com/starudream/miyoushe-task/job"
	"github.com/starudream/miyoushe-task/util"
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

			res, err := mihoyo.GenQRCode(account)
			if err != nil {
				return fmt.Errorf("generate qrcode error: %w", err)
			}
			slog.Info("qrcode content: %s", res.Url)
			fmt.Printf("\n\n%s\n\n", util.QRCode(res.Url))

			account, err = job.WaitQRCodeConfirmed(res.Ticket, account)
			if err != nil {
				return err
			}
			if account.SToken == "" {
				return nil
			}

			account, err = job.RefreshCToken(account)
			if err != nil {
				return err
			}
			account, err = job.RefreshSToken(account)
			if err != nil {
				return err
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
