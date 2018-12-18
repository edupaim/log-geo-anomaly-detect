package application

import (
	"os"
	"io/ioutil"
	"time"
)

func WatchFile(filePath string) (chan string, chan error) {
	tailChan := make(chan string)
	errChan := make(chan error)
	go watchFileRoutine(filePath, tailChan, errChan)
	return tailChan, errChan
}

func watchFileRoutine(filePath string, tailChan chan string, errChan chan error) {
	initialStat, err := os.Stat(filePath)
	if err != nil {
		errChan <- err
		return
	}
	for {
		stat, err := os.Stat(filePath)
		if err != nil {
			return
		}
		if stat.Size() != initialStat.Size() || stat.ModTime() != initialStat.ModTime() {
			f, err := os.Open(filePath)
			if err != nil {
				errChan <- err
				return
			}
			content, err := ioutil.ReadAll(f)
			if err != nil {
				errChan <- err
				return
			}
			tailChan <- string(content[initialStat.Size():])
			watchFileRoutine(filePath, tailChan, errChan)
		}
		time.Sleep(1 * time.Second)
	}
}
