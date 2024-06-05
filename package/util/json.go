package util

import "github.com/bytedance/sonic"

func JsonFmt(value any) string {
	res, err := sonic.Marshal(value)
	if err != nil {
		return ""
	}
	return string(res)
}
