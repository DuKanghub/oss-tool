package utils

import (
	"fmt"
	"path/filepath"
	"time"
)

// 从全路径中提取文件名同时加入当前时间戳
func parseFileName(file string) string {
	fileName := filepath.Base(file)
	fmt.Println(fileName)
	return fmt.Sprintf("%d_%s", time.Now().Unix(),fileName)
}
