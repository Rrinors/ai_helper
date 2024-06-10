package service

import (
	"ai_helper/biz/db"
	"ai_helper/biz/minio"
	"ai_helper/biz/model/module/qwen"
	qwen_module "ai_helper/biz/module/qwen"
	"ai_helper/package/config"
	"ai_helper/package/constant"
	"ai_helper/package/util"
	"context"
	"fmt"

	"github.com/bytedance/sonic"
)

func SubmitQwenTask(ctx context.Context, req *qwen.QwenApiRequest) *qwen.QwenApiResponse {
	if req.UserId == uint64(0) {
		return &qwen.QwenApiResponse{
			StatusCode: 400,
			StatusMsg:  "empty user_id is invalid",
		}
	}

	if req.InputContent == "" {
		return &qwen.QwenApiResponse{
			StatusCode: 400,
			StatusMsg:  "empty input is invalid",
		}
	}

	if req.HistoryNum < 0 || req.Timeout < 0 {
		return &qwen.QwenApiResponse{
			StatusCode: 400,
			StatusMsg:  "invalid task params value",
		}
	}

	model := req.InputModel
	if model == "" {
		model = "qwen-long"
	}
	taskDO, err := db.CreateTask(req.UserId, constant.Qwen, model, int(req.HistoryNum), "", "", int(req.Timeout))
	if err != nil {
		return &qwen.QwenApiResponse{
			StatusCode: 500,
			StatusMsg:  fmt.Sprintf("create qwen task failed, err=%v", err),
		}
	}
	taskDO.InputUrl = fmt.Sprintf("task#%v_input.json", taskDO.Id)
	taskDO.OutputUrl = fmt.Sprintf("task#%v_output.json", taskDO.Id)
	err = db.UpdateTask(taskDO)
	if err != nil {
		return &qwen.QwenApiResponse{
			StatusCode: 500,
			StatusMsg:  fmt.Sprintf("update qwen task failed, err=%v", err),
		}
	}

	role := req.InputRole
	if role == "" {
		role = "user"
	}
	inputMap := map[string]string{
		"role":    role,
		"content": req.InputContent,
	}
	inputConfig, _ := sonic.Marshal(inputMap)
	err = minio.UploadFile(ctx, config.MinioBucketMap[constant.Qwen], taskDO.InputUrl, inputConfig)
	if err != nil {
		return &qwen.QwenApiResponse{
			StatusCode: 500,
			StatusMsg:  fmt.Sprintf("upload qwen input failed, err=%v", err),
		}
	}

	return &qwen.QwenApiResponse{
		StatusCode: 0,
		StatusMsg:  util.JsonFmt(taskDO),
	}
}

func QueryQwenTaskResult(ctx context.Context, req *qwen.QwenApiRequest) *qwen.QwenApiResponse {
	if req.Id == uint64(0) {
		return &qwen.QwenApiResponse{
			StatusCode: 400,
			StatusMsg:  "empty task_id is invalid",
		}
	}

	task, err := db.FetchTaskById(req.Id)
	if err != nil {
		return &qwen.QwenApiResponse{
			StatusCode: 500,
			StatusMsg:  fmt.Sprintf("get task failed, err=%v", err),
		}
	}
	if task.Status != constant.TaskSuccess {
		if task.Status == constant.TaskFailed {
			return &qwen.QwenApiResponse{
				StatusCode: 500,
				StatusMsg:  "task failed",
			}
		}
		return &qwen.QwenApiResponse{
			StatusCode: 202,
			StatusMsg:  fmt.Sprintf("task not finished, status=%v", task.Status),
		}
	}

	data, err := minio.DownloadFile(ctx, config.MinioBucketMap[constant.Qwen], task.OutputUrl)
	if err != nil {
		return &qwen.QwenApiResponse{
			StatusCode: 500,
			StatusMsg:  fmt.Sprintf("download resp failed, err=%v", err),
		}
	}

	confMap := map[string]any{}
	err = sonic.Unmarshal(data, &confMap)
	if err != nil {
		return &qwen.QwenApiResponse{
			StatusCode: 500,
			StatusMsg:  fmt.Sprintf("parse resp failed, err=%v", err),
		}
	}
	message := qwen_module.GetRespMessage(confMap)
	return &qwen.QwenApiResponse{
		StatusCode: 0,
		StatusMsg:  message.Content,
	}
}
