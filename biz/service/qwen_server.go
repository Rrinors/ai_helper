package service

import (
	"ai_helper/biz/model/module/qwen"
)

func SubmitQwenTask(req *qwen.QwenApiRequest) qwen.QwenApiResponse {
	if req.UserId == uint64(0) {
		return qwen.QwenApiResponse{
			StatusCode: 400,
			StatusMsg:  "empty user_id is invalid",
		}
	}
	return qwen.QwenApiResponse{
		StatusCode: 0,
		StatusMsg:  "submit qwen task success",
	}
}
