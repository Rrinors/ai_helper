package util

import (
	"ai_helper/package/log"
	"runtime/debug"
)

func GoSafe(f func()) {
	defer func() {
		if err := recover(); err != nil {
			log.Error("recover from panic: err=%v, stack=%v", err, debug.Stack())
		}
	}()
	go f()
}
