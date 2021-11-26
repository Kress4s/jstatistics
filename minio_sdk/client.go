package minio_sdk

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type client struct {
	client *minio.Client
}

func New(host, accessKeyID, secretAccessKey string, ssl bool) (Client, error) {
	c, err := minio.New(
		host, &minio.Options{
			Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
			Secure: ssl,
		})
	if err != nil {
		return nil, err
	}
	return &client{client: c}, nil
}
