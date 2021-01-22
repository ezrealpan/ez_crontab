package crontabd

import (
	"context"

	"ezreal.com.cn/ez_crontab/crontabd"
	"ezreal.com.cn/ez_crontab/models"
	"github.com/spf13/cobra"
)

// NewCrontabdCmd ...
func NewCrontabdCmd() *cobra.Command {
	var serverCmd = &cobra.Command{
		Use:   "crontabd [string to echo]",
		Short: "web server",
		Run:   runCrontabd,
	}

	return serverCmd
}

func runCrontabd(cmd *cobra.Command, args []string) {
	ctx, _ := context.WithCancel(context.Background())
	dj := crontabd.DaemonJob{
		Job: &models.DaemonJob{},
	}
	dj.Do(ctx)
}
