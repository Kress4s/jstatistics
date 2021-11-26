package minio_sdk

import (
	"context"
	"io"
	"time"

	"github.com/minio/minio-go/v7"
)

type Object struct {
	Name        string    `json:"name"`
	Size        int64     `json:"size"`
	ExpireAt    time.Time `json:"expire_at"`
	ContentType string    `json:"content_type"`
}

func (c *client) UploadObject(bucketName, objName string, content []byte) error {
	if err := c.MakeBucket(bucketName); err != nil {
		return err
	}
	buff := buffer{
		content: content,
	}
	defer buff.Close()
	_, err := c.client.PutObject(context.Background(),
		bucketName, objName, &buff, buff.Size(), minio.PutObjectOptions{})
	return err
}

func (c *client) UploadObjectFromReader(bucketName, objName string, reader io.Reader, objSize int64) error {
	if err := c.MakeBucket(bucketName); err != nil {
		return err
	}
	_, err := c.client.PutObject(context.Background(), bucketName, objName, reader, objSize, minio.PutObjectOptions{})
	return err
}

func (c *client) UploadObjectFromFile(bucketName, objName, filePath string) error {
	if err := c.MakeBucket(bucketName); err != nil {
		return err
	}
	_, err := c.client.FPutObject(context.Background(), bucketName, objName, filePath, minio.PutObjectOptions{})
	return err
}

func (c *client) DownloadObject(bucketName, objName string) ([]byte, error) {
	obj, err := c.client.GetObject(context.Background(), bucketName, objName, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	defer obj.Close()
	return io.ReadAll(obj)
}

func (c *client) DeleteObject(bucketName, objName string) error {
	return c.client.RemoveObject(context.Background(), bucketName, objName, minio.RemoveObjectOptions{})
}

func (c *client) ListObjects(bucketName string) ([]Object, error) {
	objects := make([]Object, 0)
	objs := c.client.ListObjects(context.Background(), bucketName, minio.ListObjectsOptions{})
	for obj := range objs {
		objects = append(objects, Object{
			Name:        obj.Key,
			Size:        obj.Size,
			ExpireAt:    obj.Expires,
			ContentType: obj.ContentType,
		})
	}
	return objects, nil
}
