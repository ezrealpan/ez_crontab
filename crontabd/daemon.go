package crontabd

import (
	"context"
	"time"

	"ezreal.com.cn/ez_crontab/models"
	"github.com/ngaut/log"
)

// DaemonJob ...
type DaemonJob struct {
	Job        *models.DaemonJob
	Ctx        context.Context
	cancel     context.CancelFunc
	processNum int
}

// Do ...
func (d *DaemonJob) Do(ctx context.Context) {
	if d.Job == nil {
		return
	}
	t := time.NewTicker(1 * time.Second)
	defer t.Stop()
	retryNum := d.Job.RetryNum

	for {
		var (
			stop bool
			err  error
		)
		myCmd := command{
			Ctx:  ctx,
			Name: d.Job.CommandName,
			Args: d.Job.Command,
			Env:  d.Job.WorkEnv,
			Dir:  d.Job.WorkDir,
			User: d.Job.WorkUser,
		}

		log.Info("exec daemon job, jobName:", d.Job.Name, " jobID", d.Job.ID)

		err = myCmd.Launch()
		retryNum--

		select {
		case <-ctx.Done():
			stop = true
		case <-t.C:
		}

		if stop || d.Job.FailRestart == false || (d.Job.RetryNum > 0 && retryNum == 0) {
			break
		}
		log.Error(err)
	}
	t.Stop()
	log.Info("daemon task end", d.Job.Name)
}
