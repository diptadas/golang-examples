package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func Write() error {
	dockerDir := filepath.Join(os.Getenv("HOME"), ".docker")
	dockerConfig := filepath.Join(dockerDir, "config.json")
	fmt.Println(dockerConfig)
	if err := os.MkdirAll(dockerDir, os.ModePerm); err != nil {
		return err
	}
	content := "hello world"
	return ioutil.WriteFile(dockerConfig, []byte(content), 0600)
}

func main() {
	fmt.Println(">>>>>>>>")
	if err := Write(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("=========")
}
