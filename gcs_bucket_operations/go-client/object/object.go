package object

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"google.golang.org/api/iterator"
	"golang.org/x/net/context"
	"cloud.google.com/go/storage"
)

func Write(ctx context.Context, client *storage.Client, bucket, object string) error {

	f, err := os.Open("notes.txt")
	if err != nil {
		return err
	}
	defer f.Close()

	wc := client.Bucket(bucket).Object(object).NewWriter(ctx)
	if _, err = io.Copy(wc, f); err != nil {
		return err
	}
	if err := wc.Close(); err != nil {
		return err
	}
	return nil
}

func List(ctx context.Context, client *storage.Client, bucket string) error {

	it := client.Bucket(bucket).Objects(ctx, nil)
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
		fmt.Println(attrs.Name)
	}

	return nil
}

func ListByPrefix(ctx context.Context, client *storage.Client, bucket, prefix, delim string) error {

	// Prefixes and delimiters can be used to emulate directory listings.
	// Prefixes can be used filter objects starting with prefix.
	// The delimiter argument can be used to restrict the results to only the
	// objects in the given "directory". Without the delimeter, the entire  tree
	// under the prefix is returned.
	//
	// For example, given these blobs:
	//   /a/1.txt
	//   /a/b/2.txt
	//
	// If you just specify prefix="a/", you'll get back:
	//   /a/1.txt
	//   /a/b/2.txt
	//
	// However, if you specify prefix="a/" and delim="/", you'll get back:
	//   /a/1.txt

	it := client.Bucket(bucket).Objects(ctx, &storage.Query{
		Prefix:    prefix,
		Delimiter: delim,
	})
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
		fmt.Println(attrs.Name)
	}

	return nil
}

func Read(ctx context.Context, client *storage.Client, bucket, object string) ([]byte, error) {

	rc, err := client.Bucket(bucket).Object(object).NewReader(ctx)
	if err != nil {
		return nil, err
	}
	defer rc.Close()

	data, err := ioutil.ReadAll(rc)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func Attrs(ctx context.Context, client *storage.Client, bucket, object string) (*storage.ObjectAttrs, error) {

	o := client.Bucket(bucket).Object(object)
	attrs, err := o.Attrs(ctx)
	if err != nil {
		return nil, err
	}
	return attrs, nil
}

func MakePublic(ctx context.Context, client *storage.Client, bucket, object string) error {

	acl := client.Bucket(bucket).Object(object).ACL()
	if err := acl.Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
		return err
	}

	return nil
}

func Move(ctx context.Context, client *storage.Client, bucket, object string) error {

	dstName := object + "-rename"

	src := client.Bucket(bucket).Object(object)
	dst := client.Bucket(bucket).Object(dstName)

	if _, err := dst.CopierFrom(src).Run(ctx); err != nil {
		return err
	}
	if err := src.Delete(ctx); err != nil {
		return err
	}

	return nil
}

func CopyToBucket(ctx context.Context, client *storage.Client, dstBucket, srcBucket, srcObject string) error {

	dstObject := srcObject + "-copy"
	src := client.Bucket(srcBucket).Object(srcObject)
	dst := client.Bucket(dstBucket).Object(dstObject)

	if _, err := dst.CopierFrom(src).Run(ctx); err != nil {
		return err
	}

	return nil
}

func Delete(ctx context.Context, client *storage.Client, bucket, object string) error {

	o := client.Bucket(bucket).Object(object)
	if err := o.Delete(ctx); err != nil {
		return err
	}

	return nil
}

func WriteEncryptedObject(ctx context.Context, client *storage.Client, bucket, object string, secretKey []byte) error {

	obj := client.Bucket(bucket).Object(object)
	// Encrypt the object's contents.
	wc := obj.Key(secretKey).NewWriter(ctx)
	if _, err := wc.Write([]byte("top secret")); err != nil {
		return err
	}
	if err := wc.Close(); err != nil {
		return err
	}

	return nil
}

func ReadEncryptedObject(ctx context.Context, client *storage.Client, bucket, object string, secretKey []byte) ([]byte, error) {

	obj := client.Bucket(bucket).Object(object)
	rc, err := obj.Key(secretKey).NewReader(ctx)
	if err != nil {
		return nil, err
	}
	defer rc.Close()

	data, err := ioutil.ReadAll(rc)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func RotateEncryptionKey(ctx context.Context, client *storage.Client, bucket, object string, key, newKey []byte) error {

	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}
	obj := client.Bucket(bucket).Object(object)

	// obj is encrypted with key, we are encrypting it with the newKey.
	_, err = obj.Key(newKey).CopierFrom(obj.Key(key)).Run(ctx)
	if err != nil {
		return err
	}

	return nil
}

