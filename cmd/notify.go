package main

import (
	"context"

	"github.com/starudream/go-lib/cobra/v2"
	"github.com/starudream/go-lib/core/v2/utils/sliceutil"
	"github.com/starudream/go-lib/ntfy/v2"
)

var (
	notifyCmd = cobra.NewCommand(func(c *cobra.Command) {
		c.Use = "notify"
		c.Short = "Manage notify"
	})

	notifySendCmd = cobra.NewCommand(func(c *cobra.Command) {
		c.Use = "send <message>"
		c.Short = "Send notify"
		c.RunE = func(cmd *cobra.Command, args []string) error {
			msg, _ := sliceutil.GetValue(args, 0, "Hello World")
			return ntfy.Notify(context.Background(), msg)
		}
	})
)

func init() {
	notifyCmd.AddCommand(notifySendCmd)

	rootCmd.AddCommand(notifyCmd)
}
