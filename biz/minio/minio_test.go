package minio_test

import (
	"ai_helper/biz/minio"
	"fmt"
	"testing"
)

const qwenInput = `{
    "model": "qwen-turbo",
    "input": {
        "messages": [
            {
                "role": "system",
                "content": "You are a helpful assistant."
            },
            {
                "role": "user",
                "content": "你和GPT4、文心一言、kimi助手比谁更厉害？"
            }
        ]
    }
}`

func TestDownload(t *testing.T) {
	minio.Init()
	bucket := "qwen"
	object := "test_output.json"
	data, err := minio.DownloadFile(bucket, object)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(data))
}

func TestUpload(t *testing.T) {
	minio.Init()
	bucket := "qwen"
	object := "test_input.json"
	data := []byte(qwenInput)
	err := minio.UploadFile(bucket, object, data)
	if err != nil {
		t.Fatal(err)
	}
}
