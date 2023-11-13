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
		if config.C().Cron.Startup {
			cronRun()
		}
		err := cron.AddJob(config.C().Cron.Spec, "miyoushe-cron", cronRun)
		if err != nil {
			return fmt.Errorf("add cron job error: %w", err)
		}
		cron.Run()
		return nil
	}
})

func init() {
	rootCmd.AddCommand(cronCmd)
}

func cronRun() {
	for i := 0; i < len(config.C().Accounts); i++ {
		cronBBSAccount(config.C().Accounts[i])
		cronPostAccount(config.C().Accounts[i])
		cronLunaAccount(config.C().Accounts[i])
	}
}

func cronBBSAccount(account config.Account) (msg string) {
	err := job.SignBBS("", account)
	if err != nil {
		msg = fmt.Sprintf("米游社打卡失败: %v", err)
		slog.Error(msg)
	} else {
		msg = account.Phone + " 米游社打卡成功"
		slog.Info(msg)
	}
	err = ntfy.Notify(context.Background(), msg)
	if err != nil {
		slog.Error("cron miyoushe notify error: %v", err)
	}
	return
}

func cronPostAccount(account config.Account) (msg string) {
	err := job.SignPost("", account)
	if err != nil {
		msg = fmt.Sprintf("米游社帖子任务失败: %v", err)
		slog.Error(msg)
	} else {
		msg = account.Phone + " 米游社帖子任务成功"
		slog.Info(msg)
	}
	err = ntfy.Notify(context.Background(), msg)
	if err != nil {
		slog.Error("cron miyoushe notify error: %v", err)
	}
	return
}

func cronLunaAccount(account config.Account) (msg string) {
	awards, err := job.SignLuna(account)
	if err != nil {
		msg = fmt.Sprintf("米游社签到失败: %v", err)
		slog.Error(msg)
	} else {
		msg = account.Phone + " " + job.FormatAwards(awards)
		slog.Info(msg)
	}
	err = ntfy.Notify(context.Background(), msg)
	if err != nil {
		slog.Error("cron miyoushe notify error: %v", err)
	}
	return
}
