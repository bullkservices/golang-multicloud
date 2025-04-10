package gcp

import (
	"context"
	"io"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

type GCPClient struct {
	storageClient *storage.Client
}

func NewGCPClient() *GCPClient {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		panic("failed to create GCP storage client: " + err.Error())
	}
	return &GCPClient{
		storageClient: client,
	}
}

func (g *GCPClient) CreateBucket(bucketName string) error {
	ctx := context.Background()
	return g.storageClient.Bucket(bucketName).Create(ctx, "", nil)
}

func (g *GCPClient) DeleteBucket(bucketName string) error {
	ctx := context.Background()
	return g.storageClient.Bucket(bucketName).Delete(ctx)
}

func (g *GCPClient) ListFiles(bucketName string) ([]string, error) {
	ctx := context.Background()
	var files []string
	it := g.storageClient.Bucket(bucketName).Objects(ctx, nil)
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		files = append(files, attrs.Name)
	}
	return files, nil
}

func (g *GCPClient) UploadFile(bucketName, objectName string, data io.Reader) error {
	ctx := context.Background()
	wc := g.storageClient.Bucket(bucketName).Object(objectName).NewWriter(ctx)
	if _, err := io.Copy(wc, data); err != nil {
		return err
	}
	return wc.Close()
}

func (g *GCPClient) DeleteFile(bucketName, objectName string) error {
	ctx := context.Background()
	return g.storageClient.Bucket(bucketName).Object(objectName).Delete(ctx)
}

func (g *GCPClient) MoveFile(bucketName, srcObjectName, destObjectName string) error {
	ctx := context.Background()
	src := g.storageClient.Bucket(bucketName).Object(srcObjectName)
	dst := g.storageClient.Bucket(bucketName).Object(destObjectName)
	_, err := dst.CopierFrom(src).Run(ctx)
	if err != nil {
		return err
	}
	return src.Delete(ctx)
}
