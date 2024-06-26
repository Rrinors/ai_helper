// Code generated by hertz generator.

package user

import (
	"context"

	user "ai_helper/biz/model/basic/user"
	"ai_helper/biz/service"
	"ai_helper/package/log"
	"ai_helper/package/util"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// RegisterUser .
// @router api/v1/user/register [POST]
func RegisterUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UserApiRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	log.Info("/api/v1/user/register request: %v", util.JsonFmt(&req))
	resp := service.RegisterUser(ctx, &req)

	log.Info("/api/v1/user/register response: %v", util.JsonFmt(resp))
	c.JSON(consts.StatusOK, resp)
}

// BindQwenApiKey .
// @router api/v1/user/bind_qwen [POST]
func BindQwenApiKey(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UserApiRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	log.Info("/api/v1/user/bind_qwen request: %v", util.JsonFmt(&req))
	resp := service.BindQwenApiKey(ctx, &req)

	log.Info("/api/v1/user/bind_qwen response: %v", util.JsonFmt(resp))
	c.JSON(consts.StatusOK, resp)
}
