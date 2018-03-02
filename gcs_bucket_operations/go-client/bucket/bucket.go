package bucket

import (
	"fmt"

	"cloud.google.com/go/storage"
	"golang.org/x/net/context"
	"google.golang.org/api/iterator"
)

func Create(ctx context.Context, client *storage.Client, projectID, bucketName string) error {

	if err := client.Bucket(bucketName).Create(ctx, projectID, nil); err != nil {
		return err
	}

	return nil
}

func CreateWithAttrs(ctx context.Context, client *storage.Client, projectID, bucketName string) error {

	bucket := client.Bucket(bucketName)
	if err := bucket.Create(ctx, projectID, &storage.BucketAttrs{
		StorageClass: "COLDLINE",
		Location:     "asia",
	}); err != nil {
		return err
	}

	return nil
}

func List(ctx context.Context, client *storage.Client, projectID string) ([]string, error) {

	var buckets []string
	it := client.Buckets(ctx, projectID)
	for {
		battrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		buckets = append(buckets, battrs.Name)
	}

	return buckets, nil
}

func Delete(ctx context.Context, client *storage.Client, bucketName string) error {

	if err := client.Bucket(bucketName).Delete(ctx); err != nil {
		return err
	}

	return nil
}

func GetACLs(ctx context.Context, client *storage.Client, bucketName string) error {

	acls, err := client.Bucket(bucketName).ACL().List(ctx)
	if err != nil {
		return err
	}

	fmt.Println("ACL:")
	for _, rule := range acls {
		fmt.Printf("%s has role %s\n", rule.Entity, rule.Role)
	}

	return nil
}
