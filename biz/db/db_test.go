package db_test

import (
	"ai_helper/biz/db"
	"ai_helper/package/constant"
	"encoding/json"
	"fmt"
	"testing"
)

func TestCreateTask(t *testing.T) {
	db.Init()
	userId := uint64(1)
	moduleType := constant.Qwen
	inputUrl := "test_input.json"
	outputUrl := "test_output.json"
	task, err := db.CreateTask(userId, moduleType, inputUrl, outputUrl)
	if err != nil {
		t.Fatal(err)
	}
	resp, _ := json.Marshal(task)
	fmt.Println(string(resp))
}

func TestLimitedFetchPendingTasks(t *testing.T) {
	db.Init()
	tasks, err := db.LimitedFetchPendingTasks(constant.Qwen, 10)
	if err != nil {
		t.Fatal(err)
	}
	for _, task := range tasks {
		fmt.Printf("fetch task %v\n", task.Id)
	}
}
