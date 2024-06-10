package db

import (
	"ai_helper/package/config"
	"ai_helper/package/constant"
	"ai_helper/package/log"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Task struct {
	Id           uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	UserId       uint64    `gorm:"index;notNull" json:"user_id"`
	ModuleType   int       `gorm:"index;notNull" json:"module_type"`
	Status       int       `gorm:"index;notNull" json:"status"`
	ModelName    string    `json:"model_name"`
	HistoryNum   int       `json:"history_num"`
	InputUrl     string    `json:"input_url"`
	OutputUrl    string    `json:"output_url"`
	Timeout      int       `json:"timeout"`
	CreatedTime  time.Time `json:"created_time"`
	FinishedTime time.Time `json:"finished_time"`
}

type User struct {
	Id         uint64 `gorm:"primaryKey;autoIncrement" json:"id"`
	Name       string `gorm:"index;notNull" json:"name"`
	QwenApiKey string `json:"qwen_api_key"`
}

var DB *gorm.DB

func Init() {
	dsn := fmt.Sprintf(config.MysqlDSN, config.MysqlUser, config.MysqlPassword, config.MysqlHost, config.MysqlPort, config.MysqlDBName)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("mysql init failed, err=%v", err)
	}
	// update tables
	DB.AutoMigrate(User{}, Task{})
	log.Info("mysql init success")
}

func LimitedFetchPendingTasks(moduleType, limit int) ([]*Task, error) {
	var res []*Task
	db := DB.Model(Task{}).Where("module_type = ? AND status = ?", moduleType, constant.TaskPending)
	if err := db.Limit(limit).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func UpdateTask(task *Task) error {
	if err := DB.Model(Task{}).Where("id = ?", task.Id).Updates(task).Error; err != nil {
		return err
	}
	return nil
}

func FetchUserById(userId uint64) (*User, error) {
	var res *User
	err := DB.Model(User{}).First(&res, "id = ?", userId).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func CreateTask(userId uint64, moduleType int, model string, history int, inputUrl, outputUrl string, timeout int) (*Task, error) {
	task := &Task{
		UserId:       userId,
		ModuleType:   moduleType,
		Status:       constant.TaskPending,
		ModelName:    model,
		HistoryNum:   history,
		InputUrl:     inputUrl,
		OutputUrl:    outputUrl,
		Timeout:      timeout,
		CreatedTime:  time.Now(),
		FinishedTime: time.Unix(0, 0),
	}
	if err := DB.Model(Task{}).Create(&task).Error; err != nil {
		return nil, err
	}
	return task, nil
}

func CreateUser(name string, apiKeyMap map[int]string) (*User, error) {
	user := &User{
		Name:       name,
		QwenApiKey: apiKeyMap[constant.Qwen],
	}
	if err := DB.Model(User{}).Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateUser(user *User) error {
	if err := DB.Model(User{}).Where("id = ?", user.Id).Updates(user).Error; err != nil {
		return err
	}
	return nil
}

func FetchUserHistoryTasks(userId uint64, moduleType int, history int) ([]*Task, error) {
	db := DB.Model(Task{})
	db.Where("user_id = ? AND module_type = ? AND status = ?", userId, moduleType, constant.TaskSuccess)
	db.Order("id DESC")
	db.Limit(history)
	var res []*Task
	if err := db.Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func FetchTaskById(taskId uint64) (*Task, error) {
	var res *Task
	err := DB.Model(Task{}).First(&res, "id = ?", taskId).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}
