package framework

import (
	"github.com/urfave/cli"
	"github.com/zhanghup/go-framework/action"
	"github.com/zhanghup/go-framework/app"
	"os"
)

func CmdStart(cmds []cli.Command) {
	appconfig := app.GetAppConfig()
	if appconfig == nil {
		panic("配置文件未读取")
	}
	//2. 实例化midway代理
	cmd := cli.NewApp()
	cmd.Name = appconfig.System.Name
	cmd.Description = appconfig.System.Brief
	cmd.Version = appconfig.System.Version
	cmd.Action = func(c *cli.Context) error {
		return action.Run()
	}

	cmd.Commands = []cli.Command{
		//{
		//	Name:  "zip",
		//	Usage: "获取生产环境下的zip包",
		//	Action: func(c *cli.Context) error {
		//		//获取源文件列表
		//		dirstr := MustValue("system", "deploy.dir", "conf/")
		//		customer := MustValue("system", "customer", "")
		//		dirs := strings.Split(dirstr, ",")
		//
		//		files := make([]*os.File, 0)
		//
		//		for _, dir := range dirs {
		//			f, err := os.Open(dir)
		//			if err == nil {
		//				files = append(files, f)
		//			}
		//		}
		//		err := tools.Compress(files, customer+cmd.Version+"_"+time.Now().Format("0601021504")+".zip", cmd.Name)
		//		if err != nil {
		//			panic(err)
		//		}
		//		return nil
		//	},
		//},
		//{
		//	Name:  "install",
		//	Usage: "发布成服务（可根据不同的操作系统只能操作） 哎，暂时实现不了",
		//	Action: func(c *cli.Context) error {
		//		ServiceInstall()
		//		return nil
		//	},
		//},
	}

	if len(cmds) > 0 {
		cmd.Commands = append(cmd.Commands, cmds...)
	}

	if err := cmd.Run(os.Args); err != nil {
		panic(err.Error())
	}
}
