package singleton

import (
	"fmt"
	"sync"
)

// AppConfig 是单例模式接口，导出的。
// 通过该接口可以避免 GetAppConfig 返回一个包私有类型的指针
type AppConfig interface {
	readConfig()
}

// appConfig 是单例模式类，包私有
type appConfig struct{}

func (appConfig) readConfig() {
	fmt.Println("already get config")
}

var (
	instance *appConfig
	once     sync.Once
)

// GetAppConfig 获取单例模式对象
func GetAppConfig() AppConfig {
	once.Do(func() {
		instance = &appConfig{}
	})
	return instance
}
