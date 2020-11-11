package general

import (
	"io/ioutil"
	"strings"
)

// GetFiles returns an array of all target files in target dir
func GetFiles(dir, ext string) (files []string) {
	rawfiles, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	for _, file := range rawfiles {
		if !file.IsDir() &&
			file.Name()[0] != '.' &&
			strings.HasSuffix(file.Name(), ext) {
			files = append(files, file.Name())
		}
	}
	return
}

// GetDirs returns an array of all dirs in target dir
func GetDirs(dir string) (dirs []string) {
	rawfiles, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	for _, file := range rawfiles {
		if file.IsDir() &&
			file.Name()[0] != '.' {
			dirs = append(dirs, file.Name())
		}
	}
	return
}
