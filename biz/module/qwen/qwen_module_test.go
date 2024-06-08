package qwen

import (
	"ai_helper/biz/db"
	"ai_helper/biz/minio"
	"ai_helper/package/config"
	"ai_helper/package/constant"
	"fmt"
	"testing"

	"github.com/bytedance/sonic"
)

func TestGetRespMessage(t *testing.T) {
	minio.Init()
	bucket := config.MinioBucketMap[constant.Qwen]
	conf, err := minio.DownloadFile(bucket, "test_output.json")
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

func TestMakeRequestBody(t *testing.T) {
	db.Init()
	minio.Init()

	historyTasks, err := db.FetchUserHistoryTasks(1, constant.Qwen, 10)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("count=%v\n", len(historyTasks))

	bucket := config.MinioBucketMap[constant.Qwen]
	messageList := []MessageCarrier{}
	for i := len(historyTasks) - 1; i >= 0; i-- {
		historyTask := historyTasks[i]
		fmt.Printf("task_id=%v\n", historyTask.Id)
		// add history request
		conf, err := minio.DownloadFile(bucket, historyTask.InputUrl)
		if err != nil {
			continue
		}
		confMap := map[string]any{}
		err = sonic.Unmarshal(conf, &confMap)
		if err != nil {
			continue
		}
		role, ok := confMap["role"].(string)
		if !ok {
			continue
		}
		content, ok := confMap["content"].(string)
		if !ok {
			continue
		}
		messageList = append(messageList, MessageCarrier{
			Role:    role,
			Content: content,
		})
		// add history response
		respMessage := MessageCarrier{
			Role: "assistant",
		}
		conf, err = minio.DownloadFile(bucket, historyTask.OutputUrl)
		if err != nil {
			messageList = append(messageList, respMessage)
			continue
		}
		err = sonic.Unmarshal(conf, &confMap)
		if err != nil {
			messageList = append(messageList, respMessage)
			continue
		}
		respMessage = GetRespMessage(confMap)
		messageList = append(messageList, respMessage)
	}

	// add cur request
	conf, err := minio.DownloadFile(bucket, "task#3_input.json")
	if err != nil {
		t.Fatal(err)
	}
	confMap := map[string]any{}
	err = sonic.Unmarshal(conf, &confMap)
	if err != nil {
		t.Fatal(err)
	}
	role, ok := confMap["role"].(string)
	if !ok {
		t.Fatal(err)
	}
	content, ok := confMap["content"].(string)
	if !ok {
		t.Fatal(err)
	}
	messageList = append(messageList, MessageCarrier{
		Role:    role,
		Content: content,
	})

	bodyMap := map[string]any{
		"model": "qwen-turbo",
		"input": map[string]any{
			"messages": messageList,
		},
	}
	body, _ := sonic.Marshal(bodyMap)
	fmt.Println(string(body))
}
