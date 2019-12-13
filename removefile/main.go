package main

import (
	"io/ioutil"
	"log"
	"os"
	"path"
)

const (
	testdir  = "./testdata"
	filename = "myfile"
)

func main() {
	if err := os.Mkdir(testdir, 0777); err != nil {
		log.Println(err)
	}
	if err := ioutil.WriteFile(path.Join(testdir, filename), []byte("hello"), 0644); err != nil {
		log.Fatal(err)
	}
	if err := os.Remove(path.Join(testdir, filename)); err != nil {
		log.Fatal(err)
	}
}
