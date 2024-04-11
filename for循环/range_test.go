package testutil

import (
	"fmt"
	"testing"
)

// 1.22新特性，可以使用int作为for范围，同py
// range不再使用同一个loop，每次循环都是使用不同的内存存储值
func TestFor(t *testing.T) {
	for i := range 10 {
		fmt.Println(i)
		fmt.Println(&i)
	}
}
