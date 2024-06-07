package service

import (
	"ai_helper/biz/db"
	"ai_helper/biz/minio"
	"ai_helper/biz/model/module/qwen"
	"ai_helper/package/config"
	"ai_helper/package/constant"
	"ai_helper/package/util"
	"fmt"

	"github.com/bytedance/sonic"
)

func SubmitQwenTask(req *qwen.QwenApiRequest) *qwen.QwenApiResponse {
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

	model := req.InputModel
	if model == "" {
		model = "qwen-turbo"
	}
	taskDO, err := db.CreateTask(req.UserId, constant.Qwen, model, int(req.HistoryNum), "", "")
	if err != nil {
		return &qwen.QwenApiResponse{
			StatusCode: 500,
			StatusMsg:  fmt.Sprintf("create qwen task failed: err=%v", err),
		}
	}
	taskDO.InputUrl = fmt.Sprintf("task#%v_input.json", taskDO.Id)
	taskDO.OutputUrl = fmt.Sprintf("task#%v_output.json", taskDO.Id)
	err = db.UpdateTask(taskDO)
	if err != nil {
		return &qwen.QwenApiResponse{
			StatusCode: 500,
			StatusMsg:  fmt.Sprintf("update qwen task failed: err=%v", err),
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
	err = minio.UploadFile(config.MinioBucketMap[constant.Qwen], taskDO.InputUrl, inputConfig)
	if err != nil {
		return &qwen.QwenApiResponse{
			StatusCode: 500,
			StatusMsg:  fmt.Sprintf("upload qwen input failed: err=%v", err),
		}
	}

	return &qwen.QwenApiResponse{
		StatusCode: 0,
		StatusMsg:  fmt.Sprintf("submit qwen task success: %v", util.JsonFmt(taskDO)),
	}
}

func QueryQwenTask(req *qwen.QwenApiRequest) *qwen.QwenApiResponse {
	// TODO
	return &qwen.QwenApiResponse{
		StatusCode: 0,
		StatusMsg:  "query qwen task success",
	}
}
