package main

import (
	"context"

	"github.com/starudream/go-lib/core/v2/config"
	"github.com/starudream/go-lib/core/v2/utils/osutil"
	"github.com/starudream/go-lib/service/v2"
)

func init() {
	args := []string{"cron"}
	if c := config.LoadedFile(); c != "" {
		args = append(args, "-c", c)
	}
	service.AddCommand(rootCmd, service.New("miyoushe-task", serviceCron, service.WithArguments(args...)))
}

func serviceCron(context.Context) {
	osutil.ExitErr(cronRun())
}
