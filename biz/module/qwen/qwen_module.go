package qwen

import (
	"ai_helper/biz/db"
	"ai_helper/biz/minio"
	"ai_helper/package/config"
	"ai_helper/package/constant"
	"ai_helper/package/log"
	"ai_helper/package/util"
	"time"
)

type QwenModule struct {
	ThreadPool *util.ThreadPool
	ResultCh   chan qwenResult
}

type qwenResult struct {
	task *db.Task
	err  error
}

func (m *QwenModule) HandleTaskReq(task *db.Task) {
	m.ThreadPool.Submit(func() {
		m.ProcessTask(task)
	})
}

func (m *QwenModule) ProcessTask(task *db.Task) {
	var err error
	defer func() {
		result := qwenResult{
			task: task,
			err:  err,
		}
		m.ResultCh <- result
	}()

	bucket := config.MinioBucketMap[constant.Qwen]
	body, err := minio.DownloadFile(bucket, task.InputUrl)
	if err != nil {
		return
	}

	user, err := db.GetUserById(task.UserId)
	if err != nil {
		return
	}
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": user.QwenApiKey,
	}

	url := config.ModuleReqUrlMap[constant.Qwen]
	resp, err := util.RequestHttp("POST", url, headers, body)
	if err != nil {
		return
	}

	err = minio.UploadFile(bucket, task.OutputUrl, resp)
}

func (m *QwenModule) HandleTaskResult() {
	for result := range m.ResultCh {
		result.task.FinishedTime = time.Now()
		if result.err != nil {
			result.task.Status = constant.TaskFailed
			log.Error("task %v failed: err=%v", result.task.Id, result.err)
		} else {
			result.task.Status = constant.TaskSuccess
			log.Info("task %v success", result.task.Id)
		}
		if err := db.UpdateTask(result.task); err != nil {
			log.Error("update task %v status failed: err=%v", result.task.Id, err)
		}
	}
}

func NewQwenModule() *QwenModule {
	moduleType := constant.Qwen
	concurrency := config.ModuleConcurrencyMap[moduleType]
	return &QwenModule{
		ThreadPool: util.NewThreadPool(concurrency),
		ResultCh:   make(chan qwenResult, concurrency),
	}
}
