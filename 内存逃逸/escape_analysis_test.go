package testutil

import (
	"fmt"
	"testing"
)

// 内存逃逸分析主要针对一些本来在 栈 中的数据被 堆中的数据引用 ， 比如 数组 ，而堆数据主要由make，new ，append 这种引用类型来动态分配内存的，当栈中的数据逃逸到堆里，则形成内存逃逸

// 优点：更灵活的分配共享参数
// 缺点：回收开销增大，存在内存泄露的风险

// 如何在必须使用内存逃逸现象的时候避免内存泄露 ：
// 1.defer
// 2.避免循环使用，迭代类，双向链表，树
// 3.工具排查：pprof

// 返回指针给别的函数引用，导致变量s没法被回收，就行成内存逃逸
func TestAnalysis1(t *testing.T) {
	s := foo()
	sa := *s + `1`
	fmt.Println(sa)
}

func foo() *string {
	s := `foo`
	return &s
}

// 数组被切片引用
func TestAnalysis2(t *testing.T) {
	s := make([]int, 0, 10) //这时还在栈中安全的被分配内存
	s = append(s, 1)        //被引用成切片，
	fmt.Println(s)
}
