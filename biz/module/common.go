package module

import (
	"ai_helper/biz/db"
	"ai_helper/biz/module/qwen"
	"ai_helper/package/config"
	"ai_helper/package/constant"
	"ai_helper/package/log"
	"time"
)

type BaseModule interface {
	HandleTaskReq(*db.Task)
	ProcessTask(*db.Task)
	HandleTaskResult()
}

var ModuleMap map[int]BaseModule

func Init() {
	ModuleMap = map[int]BaseModule{
		constant.Qwen: qwen.NewQwenModule(),
	}
	for k, v := range ModuleMap {
		go v.HandleTaskResult()
		go dispatchTask(k)
	}
}

func dispatchTask(moduleType int) {
	concurrency := config.ModuleConcurrencyMap[moduleType]
	module := ModuleMap[moduleType]
	for {
		time.Sleep(config.DBFetchInterval)
		tasks, err := db.LimitedFetchPendingTasks(moduleType, concurrency)
		if err != nil {
			log.Error("fetch %v tasks failed: err=%v", moduleType, err)
			continue
		}
		for _, task := range tasks {
			module.HandleTaskReq(task)
		}
	}
}
