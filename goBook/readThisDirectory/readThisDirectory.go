package readThisDirectory

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func readThisDirectory(path string) error {

	fmt.Println(path)

	files, err := ioutil.ReadDir(path)

	if err != nil {
		panic(err)
	}

	for _, file := range files {
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			readThisDirectory(filePath)
		} else {
			fmt.Println(filePath)
		}
	}
	return nil
}

func DirScaner() {
	readThisDirectory("goBook/readThisDirectory")
}
