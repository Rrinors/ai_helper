package qwen

import (
	"ai_helper/biz/minio"
	"ai_helper/package/config"
	"ai_helper/package/constant"
	"context"
	"fmt"
	"testing"

	"github.com/bytedance/sonic"
)

func TestGetRespMessage(t *testing.T) {
	config.MinioHost = "localhost"
	minio.Init()
	bucket := config.MinioBucketMap[constant.Qwen]
	conf, err := minio.DownloadFile(context.Background(), bucket, "task#4_output.json")
	if err != nil {
		t.Fatal(err)
	}
	confMap := map[string]any{}
	err = sonic.Unmarshal(conf, &confMap)
	if err != nil {
		t.Fatal(err)
	}
	resp := GetRespMessage(confMap)
	fmt.Println(resp.Content)
}
