package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

func downloadFile(url, filePath string, replace bool) error {
	if !replace {
		if _, err := os.Stat(filePath); err == nil {
			log.Println("File already exists:", filePath)
			return nil
		}
	}

	dir := path.Dir(filePath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err = os.Mkdir(dir, 0777); err != nil {
			return err
		}
	}

	log.Println("Downloading", url, "to", filePath)

	output, err := os.Create(filePath)
	if err != nil {
		log.Println("Error while creating", filePath, "-", err)
		return err
	}
	defer output.Close()

	response, err := http.Get(url)
	if err != nil {
		log.Println("Error while downloading", url, "-", err)
		return err
	}
	defer response.Body.Close()

	n, err := io.Copy(output, response.Body)
	if err != nil {
		log.Println("Error while downloading", url, "-", err)
		return err
	}

	log.Println(n, "bytes downloaded")
	return nil
}

func main() {
	url := "http://www.sample-videos.com/video/mp4/720/big_buck_bunny_720p_10mb.mp4"
	filePath := "/home/dipta/Downloads/big_buck_bunny.mp4"

	err := downloadFile(url, filePath, true)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("File downloaded successfully")
}
