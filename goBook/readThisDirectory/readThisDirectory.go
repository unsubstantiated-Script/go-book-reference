package readThisDirectory

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

func readThisDirectory(path string) error {

	fmt.Println(path)

	files, err := ioutil.ReadDir(path)

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			err := readThisDirectory(filePath)
			if err != nil {
				return err
			}
		} else {
			fmt.Println(filePath)
		}
	}
	return nil
}

func DirScaner() {
	err := readThisDirectory("goBook/readThisDirectory")
	if err != nil {
		log.Fatal(err)
	}
}
