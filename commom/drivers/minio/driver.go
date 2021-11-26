package minio

import (
	"js_statistics/config"
	"js_statistics/minio_sdk"
	"sync"
)

var (
	instance minio_sdk.Client
	once     sync.Once
)

func GetDriver() minio_sdk.Client {
	var err error
	once.Do(func() {
		cfg := config.GetConfig()
		instance, err = minio_sdk.New(
			cfg.MinIO.ADDR, cfg.MinIO.AccessKeyID, cfg.MinIO.SecretAccessKey, cfg.MinIO.SSL)
	})
	if err != nil {
		return nil
	}
	return instance
}
