package readThisDirectory

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func reportPanic() {
	p := recover()

	if p == nil {
		return
	}

	err, ok := p.(error)

	if ok {
		fmt.Println(err)
	} else {
		panic(p)
	}
}
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
	//Calling this deferred so that will still log the error panic, but not be hindered by a code crash.
	defer reportPanic()
	readThisDirectory("goBook/readThisDirectory")
}
