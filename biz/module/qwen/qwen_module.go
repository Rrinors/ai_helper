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

type MessageCarrier struct {
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

	historyTasks, err := db.FetchUserHistoryTasks(task.UserId, constant.Qwen, task.HistoryNum)
	if err != nil {
		return
	}
	log.Info("fetch %v history from task %v", len(historyTasks), task.Id)

	bucket := config.MinioBucketMap[constant.Qwen]
	messageList := []MessageCarrier{}
	for i := len(historyTasks) - 1; i >= 0; i-- {
		// add history request
		historyTask := historyTasks[i]
		conf, err := minio.DownloadFile(bucket, historyTask.InputUrl)
		if err != nil {
			continue
		}
		confMap := map[string]any{}
		err = sonic.Unmarshal(conf, &confMap)
		if err != nil {
			continue
		}
		role, ok := confMap["role"].(string)
		if !ok {
			continue
		}
		content, ok := confMap["content"].(string)
		if !ok {
			continue
		}
		reqMessage := MessageCarrier{
			Role:    role,
			Content: content,
		}
		// add history response
		conf, err = minio.DownloadFile(bucket, historyTask.OutputUrl)
		if err != nil {
			continue
		}
		err = sonic.Unmarshal(conf, &confMap)
		if err != nil {
			continue
		}
		respMessage := GetRespMessage(confMap)
		if respMessage.Content != "" {
			messageList = append(messageList, reqMessage, respMessage)
		}
	}
	// add cur request
	conf, err := minio.DownloadFile(bucket, task.InputUrl)
	if err != nil {
		return
	}
	confMap := map[string]any{}
	err = sonic.Unmarshal(conf, &confMap)
	if err != nil {
		return
	}
	role, ok := confMap["role"].(string)
	if !ok {
		return
	}
	content, ok := confMap["content"].(string)
	if !ok {
		return
	}
	messageList = append(messageList, MessageCarrier{
		Role:    role,
		Content: content,
	})

	model := task.ModelName
	if model == "" {
		model = "qwen-long"
	}
	bodyMap := map[string]any{
		"model": model,
		"input": map[string]any{
			"messages": messageList,
		},
	}
	body, _ := sonic.Marshal(bodyMap)
	log.Info("qwen request body: %v", string(body))

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
			log.Error("task %v failed, err=%v", result.task.Id, result.err)
		} else {
			result.task.Status = constant.TaskSuccess
			log.Info("task %v success", result.task.Id)
		}
		if err := db.UpdateTask(result.task); err != nil {
			log.Error("update task %v status failed, err=%v", result.task.Id, err)
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

func GetRespMessage(confMap map[string]any) MessageCarrier {
	output, ok := confMap["output"].(map[string]any)
	if !ok {
		return MessageCarrier{}
	}
	content, ok := output["text"].(string)
	if !ok {
		return getLongRespMessage(output)
	}
	return MessageCarrier{
		Role:    "assistant",
		Content: content,
	}
}

func getLongRespMessage(output map[string]any) MessageCarrier {
	choices, ok := output["choices"].([]any)
	if !ok || len(choices) == 0 {
		return MessageCarrier{}
	}
	choice, ok := choices[0].(map[string]any)
	if !ok {
		return MessageCarrier{}
	}
	message, ok := choice["message"].(map[string]any)
	if !ok {
		return MessageCarrier{}
	}
	content, ok := message["content"].(string)
	if !ok {
		return MessageCarrier{}
	}
	return MessageCarrier{
		Role:    "assistant",
		Content: content,
	}
}
