package config

import (
	"ai_helper/package/constant"
	"os"
	"time"
)

// db_config
const (
	MysqlDSN      = "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	MysqlUser     = "root"
	MysqlPassword = "mysql-root-key"
	MysqlDBName   = "ai_helper"
	MysqlPort     = 3306

	DBFetchInterval = time.Second
)

var (
	MysqlHost = os.Getenv("MYSQL_HOST")
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
	MinioUser     = "admin"
	MinioPassword = "minio-admin-key"
	MinioUseSSL   = false
	MinioPort     = 9000
)

var (
	MinioHost = os.Getenv("MINIO_HOST")

	MinioBucketMap = map[int]string{
		constant.Qwen: "qwen",
	}
)
