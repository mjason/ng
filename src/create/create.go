package create

import (
	"dev"
	"fmt"
	"os"
)

func CreateProject() {
	dirLists := [5]string{"src", "bin", "pkg", "doc", "script"}
	fileLists := [2]string{"project.json", "README.md"}

	for _, v := range dirLists {
		err := os.Mkdir(v, 0777)
		if err != nil {
			panic(err)
		}
		fmt.Printf("创建 %s, ......... ok !!!!", v)
	}

	for _, v := range fileLists {
		err := createFile(v)
		if err != nil {
			panic(err)
		}
		fmt.Printf("创建 %s, ......... ok !!!! \n", v)
	}

	dev.Gopath()

}

func createFile(name string) (err error) {
	file, err := os.Create(name)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}
