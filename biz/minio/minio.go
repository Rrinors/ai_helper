package minio

import (
	"ai_helper/package/config"
	"ai_helper/package/log"
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var minioClient *minio.Client

func Init() {
	var err error
	minioClient, err = minio.New(fmt.Sprintf("%s:%d", config.MinioHost, config.MinioPort), &minio.Options{
		Creds:  credentials.NewStaticV4(config.MinioUser, config.MinioPassword, ""),
		Secure: config.MinioUseSSL,
	})
	if err != nil {
		log.Fatal("minio init failed, err=%v", err)
	}
	// init all buckets
	ctx := context.Background()
	for _, bucket := range config.MinioBucketMap {
		if err = initBucket(ctx, bucket); err != nil {
			log.Fatal("init bucket %v failed, err=%v", bucket, err)
		}
	}
	log.Info("minio init success")
}

func UploadFile(ctx context.Context, bucket, object string, data []byte) error {
	reader := bytes.NewBuffer(data)
	options := minio.PutObjectOptions{
		ContentType: "application/json",
	}
	_, err := minioClient.PutObject(ctx, bucket, object, reader, int64(len(data)), options)
	return err
}

func DownloadFile(ctx context.Context, bucket, object string) ([]byte, error) {
	reader, err := minioClient.GetObject(ctx, bucket, object, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func initBucket(ctx context.Context, bucket string) error {
	exists, err := minioClient.BucketExists(ctx, bucket)
	if err != nil {
		return err
	}
	if exists {
		return nil
	}
	err = minioClient.MakeBucket(ctx, bucket, minio.MakeBucketOptions{})
	if err != nil {
		return err
	}
	return nil
}
