package main

import (
	"github.com/starudream/go-lib/cobra/v2"

	"github.com/starudream/miyoushe-task/config"
)

var (
	configCmd = cobra.NewCommand(func(c *cobra.Command) {
		c.Use = "config"
		c.Short = "Manage config"
	})

	configSaveCmd = cobra.NewCommand(func(c *cobra.Command) {
		c.Use = "save"
		c.RunE = func(cmd *cobra.Command, args []string) error {
			return config.Save()
		}
	})
)

func init() {
	configCmd.AddCommand(configSaveCmd)

	rootCmd.AddCommand(configCmd)
}
