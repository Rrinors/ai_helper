package module

import (
	"ai_helper/biz/db"
	"ai_helper/biz/module/qwen"
	"ai_helper/package/config"
	"ai_helper/package/constant"
	"ai_helper/package/log"
	"ai_helper/package/util"
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
		util.GoSafe(v.HandleTaskResult)
		util.GoSafe(func() {
			dispatchTask(k)
		})
	}
}

func dispatchTask(moduleType int) {
	concurrency := config.ModuleConcurrencyMap[moduleType]
	module := ModuleMap[moduleType]
	for {
		time.Sleep(config.DBFetchInterval)
		tasks, err := db.LimitedFetchPendingTasks(moduleType, 2*concurrency)
		if err != nil {
			log.Error("fetch %v tasks failed, err=%v", moduleType, err)
			continue
		}
		if len(tasks) == 0 {
			continue
		}
		log.Info("fetch %v pending tasks", len(tasks))
		for _, task := range tasks {
			task.Status = constant.TaskRunning
			db.UpdateTask(task)
			module.HandleTaskReq(task)
		}
	}
}
