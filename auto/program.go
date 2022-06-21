package auto

import (
	"fmt"
	"github.com/kardianos/service"
	"os/exec"
)

type Program struct{}

func (p *Program) Start(s service.Service) error {
	fmt.Println("server start")
	go p.run()
	return nil
}

func (p *Program) Stop(s service.Service) error {
	fmt.Println("server stop")
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
		panic(any(err))
	}

	for _, item := range Config.CmdList {
		if item.Status == false {
			continue
		}
		go func() {
			cmd := exec.Command(item.Start)
			err := cmd.Start()
			if err != nil {
				panic(any(nil))
			}
			//err = cmd.Wait()
			//if err != nil {
			//	panic(any(nil))
			//}
			fmt.Println(item.Name + " 已执行start命令")
		}()
	}
}