package main

import (
	"os/user"
	"log"
	"path/filepath"
	"os"
	"fmt"
)

func main() {
	u, err := user.Current()

	fmt.Println(u)

	if err != nil {
		log.Fatal(err)
	}

	dir := filepath.Join(u.HomeDir, ".config", "mydir")
	fmt.Println(dir)
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		log.Fatal(err)
	}
}