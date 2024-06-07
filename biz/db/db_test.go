package db

import (
	"ai_helper/package/constant"
	"ai_helper/package/util"
	"fmt"
	"testing"

	"github.com/bytedance/sonic"
)

func TestCreateTask(t *testing.T) {
	Init()
	inputUrl := "task#3_input.json"
	outputUrl := "task#3_output.json"
	task, err := CreateTask(1, constant.Qwen, "qwen-turbo", 10, inputUrl, outputUrl)
	if err != nil {
		t.Fatal(err)
	}
	resp, _ := sonic.Marshal(task)
	fmt.Println(string(resp))
}

func TestLimitedFetchPendingTasks(t *testing.T) {
	Init()
	tasks, err := LimitedFetchPendingTasks(constant.Qwen, 10)
	if err != nil {
		t.Fatal(err)
	}
	for _, task := range tasks {
		fmt.Printf("fetch task %v\n", task.Id)
	}
}

func TestUpdateTask(t *testing.T) {
	Init()
	var task Task
	if err := DB.Model(Task{}).Where("id = ?", 1).First(&task).Error; err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%v\n", task)
	task.Status = constant.TaskPending
	if err := UpdateTask(&task); err != nil {
		t.Fatal(err)
	}
}

func TestGetUserById(t *testing.T) {
	Init()
	user, err := FetchUserById(1)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(util.JsonFmt(user))
}
