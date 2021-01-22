package crontabd

import (
	"context"
	"os"
	"runtime/debug"

	"ezreal.com.cn/ez_crontab/crontabd/proc"
	"github.com/ngaut/log"
)

type command struct {
	Ctx         context.Context
	ID          int64
	AutoRestart bool
	Name        string
	Args        []string
	Env         []string
	Dir         string
	User        string
}

func (c *command) Launch() error {
	//todo: 需要添加 ip 校验
	defer func() {
		if err := recover(); err != nil {
			log.Errorf("wrapExecScript error:%v\n%s", err, debug.Stack())
		}
	}()
	//TODO
	err := c.Exec()
	return err
}

func (c *command) Exec() error {
	//TODO 测试
	c.Args = append(c.Args, "/home/nextstage/maya_002_wxrt/job/process/command.sh")
	c.Name = "bash"
	// c.Name = "E:\\mine\\圣诞树给你，放桌面上.exe"
	cmd := proc.CommandContext(c.Ctx, c.Name, c.Args...)
	log.Debugf("cmd exec name:%s args:%v", c.Name, c.Args)
	cmd.SetDir(c.Dir)
	cmd.SetEnv(os.Environ())
	cmd.SetUser(c.User)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	defer stdout.Close()
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}
	defer stderr.Close()
	if err := cmd.Start(); err != nil {
		return err
	}
	//TODO 没有重定向到文件的，写日志
	// reader := bufio.NewReader(stdout)
	// readerErr := bufio.NewReader(stderr)
	if err = cmd.Wait(); err != nil {
		return err
	}
	return nil
}
