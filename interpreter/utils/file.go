package utils

import (
	"fmt"
	"io/ioutil"
)

func FileToString(path string) string {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}

	return string(content)
}
