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
	InputUrl     string    `json:"input_url"`
	OutputUrl    string    `json:"output_url"`
	HistoryNum   int       `json:"history_num"`
	CreatedTime  time.Time `gorm:"index" json:"created_time"`
	FinishedTime time.Time `json:"finished_time"`
}

type User struct {
	Id         uint64 `gorm:"primaryKey;autoIncrement" json:"id"`
	Name       string `gorm:"index;notNull" json:"name"`
	QwenApiKey string `json:"qwen_api_key"`
}

var DB *gorm.DB

func Init() {
	dsn := fmt.Sprintf(config.MysqlDSN, config.MysqlUser, config.MysqlPassword, config.MysqlServer, config.MysqlDBName)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("mysql init failed: err=%v", err)
	}
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

func CreateTask(userId uint64, moduleType int, inputUrl, outputUrl string, history int) (*Task, error) {
	task := &Task{
		UserId:       userId,
		ModuleType:   moduleType,
		Status:       constant.TaskPending,
		InputUrl:     inputUrl,
		OutputUrl:    outputUrl,
		HistoryNum:   history,
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
	var res []*Task
	err := DB.Model(Task{}).Where("module_type = ? AND status = ?", moduleType, constant.TaskSuccess).Order("created_time DESC").Limit(history).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}
