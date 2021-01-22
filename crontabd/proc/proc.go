package proc

import (
	"context"
	"os"
	"os/exec"
	"os/user"
	"syscall"

	"github.com/ngaut/log"
)

//Cmd ...
type Cmd struct {
	ctx context.Context
	*exec.Cmd
	isKillChildProcess bool
	done               chan struct{}
}

// CommandContext ...
func CommandContext(ctx context.Context, name string, arg ...string) *Cmd {
	cmd := exec.CommandContext(ctx, name, arg...)
	return &Cmd{
		ctx:                ctx,
		Cmd:                cmd,
		isKillChildProcess: true,
		done:               make(chan struct{}),
	}
}

// SetEnv 设置环境变量
func (c *Cmd) SetEnv(env []string) {
	if len(env) == 0 {
		return
	}
	c.Cmd.Env = env
}

// SetDir 设置工作目录
func (c *Cmd) SetDir(dir string) {
	if dir == "" {
		return
	}
	_, err := os.Stat(dir)
	if err != nil || os.IsExist(err) == false {
		return
	}
	c.Cmd.Dir = dir
}

// SetUser 设置执行用户要保证root权限 TODO
func (c *Cmd) SetUser(username string) {
	if username == "" {
		return
	}
	u, err := user.Lookup(username)
	if err != nil {
		log.Error("setUser error:", err)
		return
	}

	log.Infof("Cmd set uid=%s,gid=%s", u.Uid, u.Gid)
}

// KillAll ...
func (c *Cmd) KillAll() {

	select {
	case c.done <- struct{}{}:
	default:
	}

	if c.Process == nil {
		return
	}

	if c.isKillChildProcess == false {
		return
	}

	group, err := os.FindProcess(c.Process.Pid)
	if err == nil {
		group.Signal(syscall.SIGKILL)
	}
}

// Wait ...
func (c *Cmd) Wait() error {
	defer c.KillAll()
	go func() {
		select {
		case <-c.ctx.Done():
			c.KillAll()
		case <-c.done:
		}
	}()
	return c.Cmd.Wait()
}
