
package bucket

import (
    "bytes"
    "testing"
)

func TestFactory_AWS(t *testing.T) {
    client, err := NewBucketClient("aws")
    if err != nil {
        t.Fatal(err)
    }

    err = client.CreateBucket("test-bucket")
    if err != nil {
        t.Error("CreateBucket failed:", err)
    }

    files, err := client.ListFiles("test-bucket")
    if err != nil || len(files) == 0 {
        t.Error("ListFiles failed:", err)
    }

    err = client.UploadFile("test-bucket", "file.txt", bytes.NewBufferString("dummy"))
    if err != nil {
        t.Error("UploadFile failed:", err)
    }

    err = client.MoveFile("test-bucket", "file.txt", "file-renamed.txt")
    if err != nil {
        t.Error("MoveFile failed:", err)
    }

    err = client.DeleteFile("test-bucket", "file-renamed.txt")
    if err != nil {
        t.Error("DeleteFile failed:", err)
    }

    err = client.DeleteBucket("test-bucket")
    if err != nil {
        t.Error("DeleteBucket failed:", err)
    }
}

func TestFactory_GCP(t *testing.T) {
    client, err := NewBucketClient("gcp")
    if err != nil {
        t.Fatal(err)
    }

    err = client.CreateBucket("test-bucket")
    if err != nil {
        t.Error("CreateBucket failed:", err)
    }

    files, err := client.ListFiles("test-bucket")
    if err != nil || len(files) == 0 {
        t.Error("ListFiles failed:", err)
    }

    err = client.UploadFile("test-bucket", "file.txt", bytes.NewBufferString("dummy"))
    if err != nil {
        t.Error("UploadFile failed:", err)
    }

    err = client.MoveFile("test-bucket", "file.txt", "file-renamed.txt")
    if err != nil {
        t.Error("MoveFile failed:", err)
    }

    err = client.DeleteFile("test-bucket", "file-renamed.txt")
    if err != nil {
        t.Error("DeleteFile failed:", err)
    }

    err = client.DeleteBucket("test-bucket")
    if err != nil {
        t.Error("DeleteBucket failed:", err)
    }
}
