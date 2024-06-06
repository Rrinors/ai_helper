package qwen

import (
	"ai_helper/biz/db"
	"ai_helper/biz/minio"
	"ai_helper/package/config"
	"ai_helper/package/constant"
	"ai_helper/package/log"
	"ai_helper/package/util"
	"time"

	"github.com/bytedance/sonic"
)

type QwenModule struct {
	ThreadPool *util.ThreadPool
	ResultCh   chan taskResult
}

type taskResult struct {
	task *db.Task
	err  error
}

type messageCarrier struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func (m *QwenModule) HandleTaskReq(task *db.Task) {
	m.ThreadPool.Submit(func() {
		m.ProcessTask(task)
	})
}

func (m *QwenModule) ProcessTask(task *db.Task) {
	log.Info("start process task %v", task.Id)
	var err error
	defer func() {
		result := taskResult{
			task: task,
			err:  err,
		}
		m.ResultCh <- result
	}()

	historyTasks, err := db.FetchUserHistoryTasks(task.Id, constant.Qwen, task.HistoryNum+1)
	if err != nil {
		return
	}
	log.Info("fetch %v history from task %v", len(historyTasks), task.Id)

	bucket := config.MinioBucketMap[constant.Qwen]
	var model string
	messageList := []messageCarrier{}
	for i, historyTask := range historyTasks {
		conf, err := minio.DownloadFile(bucket, historyTask.InputUrl)
		if err != nil {
			return
		}
		var confMap map[string]any
		err = sonic.Unmarshal(conf, confMap)
		if err != nil {
			return
		}
		messageList = append(messageList, messageCarrier{
			Role:    confMap["role"].(string),
			Content: confMap["content"].(string),
		})
		if i < len(historyTasks)-1 {
			conf, err = minio.DownloadFile(bucket, historyTask.OutputUrl)
			if err != nil {
				return
			}
			err = sonic.Unmarshal(conf, confMap)
			if err != nil {
				return
			}
			messageList = append(messageList, getRespMessage(confMap))
		} else {
			model = confMap["model"].(string)
		}
	}
	bodyMap := map[string]any{
		"model": model,
		"input": map[string]any{
			"message": messageList,
		},
	}
	body, _ := sonic.Marshal(bodyMap)

	user, err := db.FetchUserById(task.UserId)
	if err != nil {
		return
	}
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + user.QwenApiKey,
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
		ResultCh:   make(chan taskResult, concurrency),
	}
}

func getRespMessage(confMap map[string]any) messageCarrier {
	choices := confMap["output"].(map[string]any)["choices"].([]any)
	message := choices[0].(map[string]any)["message"].(map[string]string)
	return messageCarrier{
		Role:    message["role"],
		Content: message["content"],
	}
}
