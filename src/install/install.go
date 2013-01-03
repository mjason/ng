package install

import (
	"appliction"
	"fmt"
	"os"
)

func Install() {
	// pwd, _ := os.Getwd()
	file := "project.json"
	info := appliction.NewInfo(file)
	createFile(info.Dependencies)
	fmt.Println("创建脚本完毕")
}

func createFile(packages []string) {
	file, err := os.Create("./script/goinstall.sh")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for _, v := range packages {
		name := fmt.Sprintf("go get -u %v \n", v)
		file.WriteString(name)
	}

}
