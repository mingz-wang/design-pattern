package singleton

import (
	"sync"
	"testing"
)

func TestSingleton(t *testing.T) {
	appConfig1 := GetAppConfig()
	appConfig2 := GetAppConfig()
	if appConfig1 != appConfig2 {
		t.Fatal("appConfig is not equal")
	}
}

const parCount = 100

func TestParallelSingleton(t *testing.T) {
	start := make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(parCount)
	appConfigs := [parCount]AppConfig{}
	for i := 0; i < parCount; i++ {
		go func(idx int) {
			// 协程阻塞，等待 chan 被关闭才继续运行
			<-start
			appConfigs[idx] = GetAppConfig()
			wg.Done()
		}(i)
	}
	// 关闭 chan，所有协程同时开始运行，实现并行
	close(start)
	wg.Wait()
	for i := 1; i < parCount; i++ {
		if appConfigs[i] != appConfigs[i-1] {
			t.Fatal("appConfig is not equal")
		}
	}
}
