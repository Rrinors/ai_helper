package minio

import (
	"fmt"
	"testing"
)

const qwenInput = `{
    "role": "user",
    "content": "你和GPT4、文心一言、kimi助手比谁更厉害？"
}`

func TestDownload(t *testing.T) {
	Init()
	bucket := "qwen"
	object := "test_input.json"
	data, err := DownloadFile(bucket, object)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(data))
}

func TestUpload(t *testing.T) {
	Init()
	bucket := "qwen"
	object := "test_input.json"
	data := []byte(qwenInput)
	err := UploadFile(bucket, object, data)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreateBucket(t *testing.T) {
	Init()
	err := initBucket("test-create")
	if err != nil {
		t.Fatal(err)
	}
}
