package app

import (
	rice "github.com/GeertJohan/go.rice"
	"github.com/zhanghup/go-framework/ctx/cfg"
)

//func StartCmd(afg ctx.ICfg, box *rice.Box, cmds []cli.Command) {
//	ctx.InitCfg(afg, box)
//	appconfig := ctx.GetCfg()
//	if appconfig == nil {
//		panic("配置文件未读取")
//	}
//	//2. 实例化midway代理
//	cmd := cli.NewApp()
//	cmd.Name = appconfig.System.Name
//	cmd.Description = appconfig.System.Brief
//	cmd.Version = appconfig.System.Version
//	cmd.Action = func(c *cli.Context) error {
//		return action.Run()
//	}
//
//	cmd.Commands = []cli.Command{
//
//	}
//
//	if len(cmds) > 0 {
//		cmd.Commands = append(cmd.Commands, cmds...)
//	}
//
//	if err := cmd.Run(os.Args); err != nil {
//		panic(err.Error())
//	}
//}

func StartCommon(afg cfg.ICfg, box *rice.Box) {
	cfg.InitCfg(afg, box)
}
