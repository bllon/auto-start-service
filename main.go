package main

import (
	"auto-start-service/auto"
	. "auto-start-service/common/mylog"
	"go.uber.org/zap"
	"os"
	"runtime"

	"github.com/kardianos/service"
)

func main() {
	NCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(NCPU)

	serConfig := &service.Config{
		Name:        "auto-start-service",
		DisplayName: "auto-start-service",
		Description: "基于go的开机自启动及启动管理",
	}

	pro := &auto.Program{}
	s, err := service.New(pro, serConfig)
	if err != nil {
		Logger.Error("service.New() err: ", zap.Error(err))
	}

	//服务注册
	if len(os.Args) > 1 {
		if os.Args[1] == "install" {
			err = s.Install()
			if err != nil {
				Logger.Error("install err: ", zap.Error(err))
			} else {
				Logger.Info("install success")
			}
			return
		}

		if os.Args[1] == "remove" {
			err = s.Uninstall()
			if err != nil {
				Logger.Error("Uninstall err: ", zap.Error(err))
			} else {
				Logger.Info("Uninstall success")
			}
			return
		}
	}

	err = s.Run() // 运行服务
	if err != nil {
		Logger.Error("service.Run err: ", zap.Error(err))
	}
}
