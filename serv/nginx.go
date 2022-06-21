package serv

import (
	"fmt"
	"os/exec"
)

func StartNginx() {
	cmd := exec.Command("E:\\nginx-1.20.2\\start.bat")
	err := cmd.Start()
	if err != nil {
		panic(any(nil))
	}
	err = cmd.Wait()
	if err != nil {
		panic(any(nil))
	}
	fmt.Println("运行结束")
}

func StopNginx() {
	cmd := exec.Command("E:\\nginx-1.20.2\\stop.bat")
	err := cmd.Start()
	if err != nil {
		panic(any(nil))
	}
	err = cmd.Wait()
	if err != nil {
		panic(any(nil))
	}
	fmt.Println("运行结束")
}