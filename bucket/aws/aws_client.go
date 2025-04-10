package aws

import (
	"context"

	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type AWSClient struct {
	s3Client *s3.Client
}

func NewAWSClient() *AWSClient {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("unable to load AWS SDK config, " + err.Error())
	}

	return &AWSClient{
		s3Client: s3.NewFromConfig(cfg),
	}
}

func (a *AWSClient) CreateBucket(bucketName string) error {
	_, err := a.s3Client.CreateBucket(context.TODO(), &s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
	})
	return err
}

func (a *AWSClient) DeleteBucket(bucketName string) error {
	_, err := a.s3Client.DeleteBucket(context.TODO(), &s3.DeleteBucketInput{
		Bucket: aws.String(bucketName),
	})
	return err
}

func (a *AWSClient) ListFiles(bucketName string) ([]string, error) {
	resp, err := a.s3Client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		return nil, err
	}

	var files []string
	for _, item := range resp.Contents {
		files = append(files, *item.Key)
	}
	return files, nil
}

func (a *AWSClient) UploadFile(bucketName, objectName string, data io.Reader) error {
	_, err := a.s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectName),
		Body:   data,
	})
	return err
}

func (a *AWSClient) DeleteFile(bucketName, objectName string) error {
	_, err := a.s3Client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectName),
	})
	return err
}

func (a *AWSClient) MoveFile(bucketName, srcObjectName, destObjectName string) error {
	_, err := a.s3Client.CopyObject(context.TODO(), &s3.CopyObjectInput{
		Bucket:     aws.String(bucketName),
		CopySource: aws.String(bucketName + "/" + srcObjectName),
		Key:        aws.String(destObjectName),
	})
	if err != nil {
		return err
	}
	return a.DeleteFile(bucketName, srcObjectName)
}
