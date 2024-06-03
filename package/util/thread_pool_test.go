package util_test

import (
	"ai_helper/package/util"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestThreadPool(t *testing.T) {
	threadPool := util.NewThreadPool(4)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		threadPool.Submit(func() {
			defer wg.Done()
			fmt.Printf("start %v\n", i)
			time.Sleep(3 * time.Second)
			fmt.Printf("finish %v\n", i)
		})
	}
	wg.Wait()
}
