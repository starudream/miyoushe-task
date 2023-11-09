package main

import (
	"fmt"

	"github.com/starudream/go-lib/cobra/v2"
	"github.com/starudream/go-lib/core/v2/utils/sliceutil"

	"github.com/starudream/miyoushe-task/config"
	"github.com/starudream/miyoushe-task/job"
)

var (
	signCmd = cobra.NewCommand(func(c *cobra.Command) {
		c.Use = "sign"
		c.Short = "Run sign task"
	})

	signBBSCmd = cobra.NewCommand(func(c *cobra.Command) {
		c.Use = "bbs <account phone> <game id>"
		c.Short = "Miyoushe bbs"
		c.Args = func(cmd *cobra.Command, args []string) error {
			phone, _ := sliceutil.GetValue(args, 0)
			if phone == "" {
				return fmt.Errorf("requires account phone")
			}
			_, exists := config.GetAccount(phone)
			if !exists {
				return fmt.Errorf("account %s not exists", phone)
			}
			return nil
		}
		c.RunE = func(cmd *cobra.Command, args []string) error {
			phone, _ := sliceutil.GetValue(args, 0)
			gameId, _ := sliceutil.GetValue(args, 1)
			account, _ := config.GetAccount(phone)
			return job.SignBBS(gameId, account)
		}
	})

	signLunaCmd = cobra.NewCommand(func(c *cobra.Command) {
		c.Use = "luna <account phone>"
		c.Short = "Game award"
		c.Args = func(cmd *cobra.Command, args []string) error {
			phone, _ := sliceutil.GetValue(args, 0)
			if phone == "" {
				return fmt.Errorf("requires account phone")
			}
			_, exists := config.GetAccount(phone)
			if !exists {
				return fmt.Errorf("account %s not exists", phone)
			}
			return nil
		}
		c.RunE = func(cmd *cobra.Command, args []string) error {
			phone, _ := sliceutil.GetValue(args, 0)
			account, _ := config.GetAccount(phone)
			_, err := job.SignLuna(account)
			return err
		}
	})
)

func init() {
	signCmd.AddCommand(signBBSCmd)
	signCmd.AddCommand(signLunaCmd)

	rootCmd.AddCommand(signCmd)
}
