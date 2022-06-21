package main

import (
	"auto-start-service/auto"
	"fmt"
	"os"
	"runtime"

	"github.com/kardianos/service"
)

var logger service.Logger

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
		fmt.Println(err, "service.New() err")
	}

	//服务注册
	if len(os.Args) > 1 {
		if os.Args[1] == "install" {
			err = s.Install()
			if err != nil {
				fmt.Println("install err", err)
			} else {
				fmt.Println("install success")
			}
			return
		}

		if os.Args[1] == "remove" {
			err = s.Uninstall()
			if err != nil {
				fmt.Println("Uninstall err", err)
			} else {
				fmt.Println("Uninstall success")
			}
			return
		}
	}

	err = s.Run() // 运行服务
	if err != nil {
		fmt.Println("s.Run err", err)
	}
}

