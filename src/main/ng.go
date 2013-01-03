package main

import (
	"create"
	"dev"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("输入的命令无效\n install: 安装依赖库 \n dev: 设置GOPATH和gocode \n init: 创建一个项目")
	} else if os.Args[1] == "init" {
		create.CreateProject()
	} else if os.Args[1] == "install" {
		// coding
	} else if os.Args[1] == "dev" {
		dev.Gopath()
	} else {
		fmt.Println("输入的命令无效\n install: 安装依赖库 \n dev: 设置GOPATH和gocode \n init: 创建一个项目")
	}
}
