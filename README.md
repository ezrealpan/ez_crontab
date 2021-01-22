参考：https://github.com/iwannay/jiacrontab


// CommandContext is like Command but includes a context.
//
// The provided context is used to kill the process (by calling
// os.Process.Kill) if the context becomes done before the command
// completes on its own.
func CommandContext(ctx context.Context, name string, arg ...string) *Cmd {
	if ctx == nil {
		panic("nil Context")
	}
	cmd := Command(name, arg...)
	cmd.ctx = ctx
	return cmd
}
shell命令很复杂，可以用sh文件。比如：

     command.sh文件内容如下：

     #!/bin/bash

     nohup /home/nextstage/maya_002_wxrt/server/maya wxrt --wxrtCfgFile="/home/nextstage/maya_002_wxrt/server/conf/wxrt.toml"  >/home/nextstage/maya_002_wxrt/maya.log &

     echo $! >> /home/nextstage/maya_002_wxrt/maya.pid

     
	//TODO 测试
	c.Args = append(c.Args, "/home/nextstage/maya_002_wxrt/job/process/command.sh")
	c.Name = "bash"
	// c.Name = "E:\\mine\\圣诞树给你，放桌面上.exe"
	cmd := exec.CommandContext(c.Ctx, c.Name, c.Args...)