package minio_sdk

import "io"

type Client interface {
	UploadObject(bucketName, objName string, content []byte) error
	UploadObjectFromReader(bucketName, objName string, reader io.Reader, objSize int64) error
	UploadObjectFromFile(bucketName, objName, filePath string) error
	DownloadObject(bucketName, objName string) ([]byte, error)
	DeleteObject(bucketName, objName string) error
	ListBuckets() ([]Bucket, error)
	ListObjects(bucketName string) ([]Object, error)
	DeleteBucket(bucketName string) error
}
