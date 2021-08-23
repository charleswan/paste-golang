package main

import (
	"fmt"
)

// go1.17 run -gcflags=-G=3 *.go

func print[T any](s []T) {
	for _, v := range s {
	 fmt.Print(v)
	}
}
   
func main() {
	print([]string{"你好, ", "脑子进了煎鱼\n"})
	print([]int64{1, 2, 3})
}
