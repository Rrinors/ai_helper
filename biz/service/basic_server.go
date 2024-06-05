package service

import (
	"ai_helper/biz/db"
	"ai_helper/biz/model/basic/user"
	"ai_helper/package/constant"
	"ai_helper/package/util"
	"fmt"
)

func RegisterUser(req *user.UserApiRequest) user.UserApiResponse {
	if req.Name == "" {
		return user.UserApiResponse{
			StatusCode: 400,
			StatusMsg:  "empty name is invalid",
		}
	}

	apiKeys := map[int]string{}
	if req.QwenApiKey != "" {
		apiKeys[constant.Qwen] = req.QwenApiKey
	}

	userDO, err := db.CreateUser(req.Name, apiKeys)
	if err != nil {
		return user.UserApiResponse{
			StatusCode: 500,
			StatusMsg:  fmt.Sprintf("create user failed: err=%v", err),
		}
	}
	return user.UserApiResponse{
		StatusCode: 0,
		StatusMsg:  fmt.Sprintf("create user success: %v", util.JsonFmt(userDO)),
	}
}

func BindQwenApiKey(req *user.UserApiRequest) user.UserApiResponse {
	if req.Id == uint64(0) {
		return user.UserApiResponse{
			StatusCode: 400,
			StatusMsg:  "empty user_id is invalid",
		}
	}

	if req.QwenApiKey == "" {
		return user.UserApiResponse{
			StatusCode: 400,
			StatusMsg:  "empty qwen api_key is invalid",
		}
	}

	userDO, err := db.GetUserById(req.Id)
	if err != nil {
		return user.UserApiResponse{
			StatusCode: 500,
			StatusMsg:  fmt.Sprintf("get user failed: err=%v", err),
		}
	}

	userDO.QwenApiKey = req.QwenApiKey
	err = db.UpdateUser(userDO)
	if err != nil {
		return user.UserApiResponse{
			StatusCode: 500,
			StatusMsg:  fmt.Sprintf("update user failed: err=%v", err),
		}
	}
	return user.UserApiResponse{
		StatusCode: 0,
		StatusMsg:  "bind qwen api_key success",
	}
}
