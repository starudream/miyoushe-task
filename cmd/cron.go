package main

import (
	"context"
	"fmt"

	"github.com/starudream/go-lib/cobra/v2"
	"github.com/starudream/go-lib/core/v2/slog"
	"github.com/starudream/go-lib/cron/v2"
	"github.com/starudream/go-lib/ntfy/v2"

	"github.com/starudream/miyoushe-task/config"
	"github.com/starudream/miyoushe-task/job"
)

var cronCmd = cobra.NewCommand(func(c *cobra.Command) {
	c.Use = "cron"
	c.Short = "Run as cron job"
	c.RunE = func(cmd *cobra.Command, args []string) error {
		return cronRun()
	}
})

func init() {
	rootCmd.AddCommand(cronCmd)
}

func cronRun() error {
	if config.C().Cron.Startup {
		cronJob()
	}
	err := cron.AddJob(config.C().Cron.Spec, "miyoushe-cron", cronJob)
	if err != nil {
		return fmt.Errorf("add cron job error: %w", err)
	}
	cron.Run()
	return nil
}

func cronJob() {
	for i := 0; i < len(config.C().Accounts); i++ {
		cronForumAccount(config.C().Accounts[i])
		cronGameAccount(config.C().Accounts[i])
	}
}

func cronForumAccount(account config.Account) (msg string) {
	record, err := job.SignForum(account)
	if err != nil {
		msg = fmt.Sprintf("%s: %v", record.Name(), err)
		slog.Error(msg)
	} else {
		msg = account.Phone + " " + record.Success()
		slog.Info(msg)
	}
	err = ntfy.Notify(context.Background(), msg)
	if err != nil {
		slog.Error("cron miyoushe notify error: %v", err)
	}
	return
}

func cronGameAccount(account config.Account) (msg string) {
	records, err := job.SignGame(account)
	if err != nil {
		msg = fmt.Sprintf("%s: %v", records.Name(), err)
		slog.Error(msg)
	} else {
		msg = account.Phone + " " + records.Success()
		slog.Info(msg)
	}
	err = ntfy.Notify(context.Background(), msg)
	if err != nil {
		slog.Error("cron miyoushe notify error: %v", err)
	}
	return
}
