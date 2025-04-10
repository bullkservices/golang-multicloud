package bucket

import (
	"errors"

	"github.com/bullkservices/golang-multicloud/bucket/aws"
	"github.com/bullkservices/golang-multicloud/bucket/gcp"
	"github.com/bullkservices/golang-multicloud/common"
)

func NewBucketClient(cloud string) (Bucket, error) {
	switch cloud {
	case common.AWS:
		return aws.NewAWSClient(), nil
	case common.GCP:
		return gcp.NewGCPClient(), nil
	default:
		return nil, errors.New("unsupported cloud provider")
	}
}
