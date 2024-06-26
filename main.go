// Code generated by hertz generator.

package main

import (
	"ai_helper/biz/db"
	"ai_helper/biz/minio"
	"ai_helper/biz/module"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func Init() {
	db.Init()
	minio.Init()
	module.Init()
}

func main() {
	Init()

	h := server.Default()

	register(h)
	h.Spin()
}
