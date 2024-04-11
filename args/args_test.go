package testutil

import (
	"fmt"
	"os"
	"testing"
)

// 获取命令行参数1
func TestArgs(t *testing.T) {
	s := os.Args[5] //获取第一位命令参数
	fmt.Println("输出的参数：", s)
}
