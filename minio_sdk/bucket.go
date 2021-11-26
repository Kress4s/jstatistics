package minio_sdk

import (
	"context"
	"time"

	"github.com/minio/minio-go/v7"
)

type Bucket struct {
	Name     string    `json:"name"`
	CreateAt time.Time `json:"create_at"`
}

func (c *client) MakeBucket(bucketName string) error {
	if err := c.client.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{Region: "China"}); err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := c.client.BucketExists(context.Background(), bucketName)
		if errBucketExists == nil && exists {
		} else {
			return err
		}
	}
	return nil
}

func (c *client) ListBuckets() ([]Bucket, error) {
	buckets := make([]Bucket, 0)
	bkts, err := c.client.ListBuckets(context.Background())
	if err != nil {
		return nil, err
	}
	for _, v := range bkts {
		buckets = append(buckets, Bucket{
			Name:     v.Name,
			CreateAt: v.CreationDate,
		})
	}
	return buckets, nil
}

func (c *client) DeleteBucket(bucketName string) error {
	return c.client.RemoveBucket(context.Background(), bucketName)
}
