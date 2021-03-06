package utils

import (
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func FileBlob(file string) string {
	projectDir, _ := os.Getwd()
	filePath := path.Join(projectDir, file)
	inpBuff, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	return string(inpBuff)
}

func FileRows(file string) []string {
	projectDir, _ := os.Getwd()
	filePath := path.Join(projectDir, file)
	inpBuff, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(inpBuff), "\n")
}
