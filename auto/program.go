package auto

import (
	. "auto-start-service/common/mylog"
	"fmt"
	"github.com/kardianos/service"
	"go.uber.org/zap"
	"os/exec"
)

type Program struct{}

func (p *Program) Start(s service.Service) error {
	Logger.Info("server start")
	go p.run()
	return nil
}

func (p *Program) Stop(s service.Service) error {
	Logger.Info("server stop")
	return nil
}

func (p *Program) run() {
	//fmt.Println("开机自启动服务 - run")
	//serv.StartNginx()
	//serv.StopNginx()

	//读取json配置文件
	//path, err := GetConfigPath()
	//if err != nil {
	//	panic(any(err))
	//}
	path := "./service.json"
	Config, err := GetConfig(path)
	if err != nil {
		Logger.Error("read service.json error: ", zap.Error(err))
	}

	for _, item := range Config.CmdList {
		fmt.Println()
		if item.Status == false {
			continue
		}
		go func() {
			cmd := exec.Command(item.Start)
			err := cmd.Start()
			if err != nil {
				panic(any(nil))
				Logger.Error(item.Name+" service start error: ", zap.Error(err))
			}
			Logger.Info(item.Name + " 已执行")

			//err = cmd.Wait()
			//if err != nil {
			//	Logger.Error(item.Name+" 已停止运行, error: ", zap.Error(err))
			//}
			//
			//Logger.Info(item.Name + " 已结束执行")
		}()
	}
}
