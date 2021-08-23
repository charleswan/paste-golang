package main

import (
	"fmt"
	"reflect"
)

type testStruct struct {
	a string
	b int64
}

type testStruct2 struct {
	a string
	b int64
	m map[int64]string
}

// 检测是否可以比较
func IsComparable(key interface{}) bool {
	return reflect.TypeOf(key).Comparable()
}

func main() {
	m := make(map[string]int64)
	fmt.Println(IsComparable(m)) // false

	s := testStruct{}
	fmt.Println(IsComparable(s)) // true

	s2 := testStruct2{}
	fmt.Println(IsComparable(s2)) // false
}
