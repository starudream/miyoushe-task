package main

import (
	"github.com/starudream/go-lib/cobra/v2"

	"github.com/starudream/miyoushe-task/job"
)

var (
	signCmd = cobra.NewCommand(func(c *cobra.Command) {
		c.Use = "sign"
		c.Short = "Run sign task"
	})

	signForumCmd = cobra.NewCommand(func(c *cobra.Command) {
		c.Use = "forum <account phone>"
		c.Short = "Miyoushe forum task"
		c.RunE = func(cmd *cobra.Command, args []string) error {
			_, err := job.SignForum(xGetAccount(args))
			return err
		}
	})

	signGameCmd = cobra.NewCommand(func(c *cobra.Command) {
		c.Use = "game <account phone>"
		c.Short = "Miyoushe game award"
		c.RunE = func(cmd *cobra.Command, args []string) error {
			_, err := job.SignGame(xGetAccount(args))
			return err
		}
	})
)

func init() {
	signCmd.AddCommand(signForumCmd)
	signCmd.AddCommand(signGameCmd)

	rootCmd.AddCommand(signCmd)
}
