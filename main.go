package main

import (
	"ai_helper/biz/db"
	"ai_helper/biz/minio"
	"ai_helper/biz/module"
)

func main() {
	db.Init()
	minio.Init()
	module.Init()
}