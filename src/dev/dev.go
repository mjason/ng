package dev

import (
	"fmt"
	"os"
)

func Gopath() {
	file, err := os.Create("./script/dev.sh")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	pwd, _ := os.Getwd()
	file.WriteString("export GOPATH=" + pwd)
	fmt.Printf("创建dev脚本完成, %v \n", pwd)
}
