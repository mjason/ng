package appliction

import (
	"encoding/json"
	"io/ioutil"
)

type Info struct {
	Name         string
	Version      string
	License      string
	Dependencies []string
}

func NewInfo(path string) *Info {
	file, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	}
	var m Info
	json.Unmarshal(file, &m)

	return &m
}
