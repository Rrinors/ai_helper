package minio

import (
	"ai_helper/package/config"
	"ai_helper/package/log"
	"bytes"
	"context"
	"io"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var minioClient *minio.Client

func Init() {
	var err error
	minioClient, err = minio.New(config.MinioServer, &minio.Options{
		Creds:  credentials.NewStaticV4(config.MinioUser, config.MinioPassword, ""),
		Secure: config.MinioUseSSL,
	})
	if err != nil {
		log.Fatal("minio init failed: err=%v", err)
	}
	log.Info("minio init success")
}

func UploadFile(bucket, object string, data []byte) error {
	reader := bytes.NewBuffer(data)
	options := minio.PutObjectOptions{
		ContentType: "application/json",
	}
	_, err := minioClient.PutObject(context.Background(), bucket, object, reader, int64(len(data)), options)
	return err
}

func DownloadFile(bucket, object string) ([]byte, error) {
	reader, err := minioClient.GetObject(context.Background(), bucket, object, minio.GetObjectOptions{})
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