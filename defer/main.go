package main

import (
	"os"
	"path/filepath"
	"fmt"
)

var (
	dirname = "newdir"
	filename = "newfile"
)

func dirTest() error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	err = os.MkdirAll(filepath.Join(cwd, dirname), 0755)
	if err != nil {
		return err
	}

	defer os.RemoveAll(dirname)

	fmt.Println(filepath.Join(cwd, dirname, filename))
	f, err := os.Create(filepath.Join(dirname, filename))
	if err != nil {
		return err
	}
	defer f.Close()

	return nil
}

func main() {
	dirTest()
}