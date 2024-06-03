package config

import (
	"ai_helper/package/constant"
	"time"
)

// db_config
const (
	MysqlDSN      = "%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	MysqlUser     = "root"
	MysqlPassword = "mysql"
	MysqlServer   = "localhost:3306"
	MysqlDBName   = "ai_helper"

	DBFetchInterval = time.Second
)

// module_config
var (
	ModuleConcurrencyMap = map[int]int{
		constant.Qwen: 10,
	}
	ModuleReqUrlMap = map[int]string{
		constant.Qwen: "https://dashscope.aliyuncs.com/api/v1/services/aigc/text-generation/generation",
	}
)

// minio_config
const (
	MinioServer   = "localhost:9000"
	MinioUser     = "admin"
	MinioPassword = "minio-key"
	MinioUseSSL   = false
)

var (
	MinioBucketMap = map[int]string{
		constant.Qwen: "qwen",
	}
)
