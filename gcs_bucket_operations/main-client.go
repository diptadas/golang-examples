package main

import (
	"fmt"
	"log"
	"golang.org/x/net/context"
	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
	"gcs-bucket-operations/go-client/bucket"
	"gcs-bucket-operations/go-client/object"
)

func main() {
	ctx := context.Background()

	projectID := "tigerworks-kube"

	client, err := storage.NewClient(ctx, option.WithServiceAccountFile("tigerworks-kube.json"))

	if err != nil {
		log.Fatal(err)
	}

	bucketName := "restic-test-dipta"

	if err := bucket.Create(ctx, client, projectID, bucketName); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("created bucket: %v\n", bucketName)
	}

	buckets, err := bucket.List(ctx, client, projectID)
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println("bucket list:")
		for _,bucket := range buckets {
			fmt.Println(bucket)
		}
	}

	if err := bucket.GetACLs(ctx, client, bucketName); err != nil {
		fmt.Println(err)
	}

	if err := bucket.Delete(ctx, client, bucketName); err != nil {
		fmt.Println(err)
	}else{
		fmt.Printf("deleted bucket: %v\n", bucketName)
	}

	if err := object.Write(ctx, client, bucketName, "new-file.txt"); err != nil {
		fmt.Println(err)
	}else {
		fmt.Println("file uploaded")
	}

	if err := object.List(ctx, client, bucketName); err != nil {
		fmt.Println(err)
	}
}

