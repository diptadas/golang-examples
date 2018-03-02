package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	storage "google.golang.org/api/storage/v1"
)

func getStorageService() (*storage.Service, error) {

	raw, err := ioutil.ReadFile("tigerworks-kube.json")
	if err != nil {
		return nil, err
	}

	conf, err := google.JWTConfigFromJSON(raw, storage.DevstorageReadWriteScope)
	if err != nil {
		return nil, err
	}

	client := conf.Client(oauth2.NoContext)

	service, err := storage.New(client)
	if err != nil {
		return nil, err
	}

	return service, nil
}

func main() {

	projectID := "tigerworks-kube"

	service, err := getStorageService()

	if err != nil {
		log.Fatal(err)
	}

	bucketName := "restic-test-dipta"

	// create bucket if not exists

	if _, err := service.Buckets.Get(bucketName).Do(); err != nil {
		fmt.Println("bucket not exists, creating bucket...")
		if _, err := service.Buckets.Insert(projectID, &storage.Bucket{Name: bucketName}).Do(); err != nil {
			fmt.Println("error creating bucket")
		} else {
			fmt.Println("bucket created")
		}
	} else {
		fmt.Println("bucket exists")
	}

	// list all buckets

	if list, err := service.Buckets.List(projectID).Do(); err != nil {
		fmt.Println(err)
	} else {
		for i := range list.Items {
			fmt.Println(list.Items[i].Name)
		}
	}

	// list objects

	prefix := "abc/"
	if list, err := service.Objects.List(bucketName).Prefix(prefix).Delimiter("/").Do(); err != nil {
		fmt.Println(err)
	} else {
		for _, item := range list.Prefixes {
			name := strings.TrimPrefix(item, prefix)
			fmt.Println(name)
		}
		for _, item := range list.Items {
			name := strings.TrimPrefix(item.Name, prefix)
			if name != "" {
				fmt.Println(name)
			}
		}
	}

	// insert object

	objName := "abc/file.txt"
	object := &storage.Object{Name: objName}

	if f, err := os.Open("notes.txt"); err != nil {
		fmt.Println(err)
	} else {
		if _, err := service.Objects.Insert(bucketName, object).Media(f).Do(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("object inserted")
			f.Close()
		}
	}

	oldName := "abc/file.txt"
	newName := "abc/def/file-new.txt"

	if _, err := service.Objects.Copy(bucketName, oldName, bucketName, newName, &storage.Object{}).Do(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("object copied")
	}

	// delete object

	/*if err := service.Objects.Delete(bucketName, objName).Do(); err != nil {
		fmt.Println(err)
	}else{
		fmt.Println("object deleted")
	}*/
}
