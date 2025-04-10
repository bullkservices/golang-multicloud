
# Multi-Cloud Bucket Library

This Go module provides a unified interface to manage bucket storage in AWS and GCP.

## Features

- Create/Delete Buckets
- Upload/List/Delete/Move Files
- Factory pattern for selecting the cloud provider
- Unit tested and cleanly structured

## Usage

### Import the Module

```go
import "github.com/your-org/golang-multicloud/bucket"
```

### Create a Client

```go
client, err := bucket.NewBucketClient("aws") // or "gcp"
if err != nil {
    log.Fatal(err)
}
```

### Perform Operations

```go
client.CreateBucket("my-bucket")
client.UploadFile("my-bucket", "hello.txt", strings.NewReader("hello world"))
files, _ := client.ListFiles("my-bucket")
fmt.Println(files)
client.MoveFile("my-bucket", "hello.txt", "renamed.txt")
client.DeleteFile("my-bucket", "renamed.txt")
client.DeleteBucket("my-bucket")
```

## Testing

Run unit tests using:

```bash
go test ./bucket
```
