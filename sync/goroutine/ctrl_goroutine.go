package main

import (
	"fmt"
	"time"
)

// 使用 chan+select 方式通知 goroutine 退出，非常优雅。
// 但有局限性：如果有很多goroutine都需要控制结束怎么办呢？
// 如果这些goroutine又衍生了其他更多的goroutine怎么办呢？
// 如果一层层的无穷尽的goroutine呢？
// 这就非常复杂了，即使我们定义很多chan也很难解决这个问题，因为goroutine的关系链就导致了这种场景非常复杂。
// 这个时候就可以考虑使用 context 了。

func main() {
	stop := make(chan bool)

	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("监控退出，停止了...")
				return
			default:
				fmt.Println("goroutine监控中...")
				time.Sleep(2 * time.Second)
			}
		}
	}()

	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	stop <- true
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}
