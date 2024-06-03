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
	Id           uint64 `gorm:"primaryKey;autoIncrement"`
	UserId       uint64 `gorm:"index;notNull"`
	ModuleType   int    `gorm:"index;notNull"`
	Status       int    `gorm:"index;notNull"`
	InputUrl     string
	OutputUrl    string
	CreatedTime  time.Time
	FinishedTime time.Time
}

type User struct {
	Id         uint64 `gorm:"primaryKey;autoIncrement"`
	Name       string `gorm:"index;notNull"`
	QwenApiKey string
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
	if err := DB.Model(Task{}).Updates(task).Error; err != nil {
		return err
	}
	return nil
}

func GetUserById(userId uint64) (*User, error) {
	var res *User
	err := DB.Model(User{}).Where("id = ?", userId).First(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func CreateTask(userId uint64, moduleType int, inputUrl, outputUrl string) (*Task, error) {
	task := &Task{
		UserId:       userId,
		ModuleType:   moduleType,
		Status:       constant.TaskPending,
		InputUrl:     inputUrl,
		OutputUrl:    outputUrl,
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
