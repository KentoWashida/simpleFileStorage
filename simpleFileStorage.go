package simpleFileStorage

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

var mu sync.Mutex

func GetTopDirs(mainpath string) (dirNameSlice []string, err error) {
	if !CheckExists(mainpath) {
		if err := os.MkdirAll(mainpath, 0774); err != nil {
			return nil, err
		}
		return nil, nil
	} else {
		files, err := ioutil.ReadDir(mainpath)
		if err != nil {
			return nil, err
		}
		filelist := []string{}

		for _, file := range files {
			filelist = append(filelist, file.Name())
		}
		return filelist, nil
	}
}

func CheckExists(fileName string) bool {
	_, err := os.Stat(fileName)
	return err == nil
}

func ReadFile(fileName string) ([]byte, error) {
	rdata, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return rdata, nil
}

func WriteFile(fileName string, wdata []byte) error {
	mu.Lock()
	os.MkdirAll(filepath.Dir(fileName), 0774)
	err := ioutil.WriteFile(fileName, wdata, 0664)
	mu.Unlock()
	return err
}

func AddFile(fileName string, wdata []byte) error {
	mu.Lock()
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0664)
	if err != nil {
		return err
	}
	fmt.Fprint(file, string(wdata))
	file.Close()
	mu.Unlock()
	return nil
}

func DeleteFile(fileName string) error {
	mu.Lock()
	err := os.Remove(fileName)
	mu.Unlock()
	return err
}
