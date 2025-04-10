
package bucket

import "io"

type Bucket interface {
    CreateBucket(bucketName string) error
    DeleteBucket(bucketName string) error
    ListFiles(bucketName string) ([]string, error)
    UploadFile(bucketName, objectName string, data io.Reader) error
    DeleteFile(bucketName, objectName string) error
    MoveFile(bucketName, srcObjectName, destObjectName string) error
}
