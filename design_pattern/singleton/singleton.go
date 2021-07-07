package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// 单例模式优缺点
// 1、优点
//   仅在首次请求单例对象时对其进行初始化。
//   提供对唯一实例的受控访问
// 2、缺点
//   该模式在多线程环境下需要进行特殊处理， 避免多个线程多次创建单例对象。
//   缺少抽象层而难以扩展
//   违背了“单一职责原则”。

type SingleLoad struct {
	name string
}

var (
	singleLoadIns *SingleLoad
	mu            sync.Mutex
	initialized   uint32
	once          = &sync.Once{}
	ch            = make(chan int, 1) // 定义缓冲型的 channel
)

// 懒汉式 mutex
func GetLoadInstance1() *SingleLoad {
	if singleLoadIns == nil {
		mu.Lock()
		defer mu.Unlock()

		if singleLoadIns == nil {
			singleLoadIns = &SingleLoad{
				name: "singleton",
			}
			time.Sleep(3 * time.Second)
		}
	}
	return singleLoadIns
}

// 懒汉式 channel
func GetLoadInstance2() *SingleLoad {
	if singleLoadIns == nil {
		ch <- 1
		defer func() {
			<-ch
		}()

		if singleLoadIns == nil {
			singleLoadIns = &SingleLoad{
				name: "singleton",
			}
			time.Sleep(3 * time.Second)
		}
	}
	return singleLoadIns
}

// 设置初始化标志
func GetLoadInstance3() *SingleLoad {
	if atomic.LoadUint32(&initialized) == 1 {
		return singleLoadIns
	}

	mu.Lock()
	defer mu.Unlock()

	if singleLoadIns == nil {
		singleLoadIns = &SingleLoad{
			name: "singleton",
		}
		time.Sleep(3 * time.Second)
		atomic.StoreUint32(&initialized, 1)
	}
	return singleLoadIns
}

// 懒汉式
func GetLoadInstance4() *SingleLoad {
	if singleLoadIns == nil {
		once.Do(func() {
			singleLoadIns = &SingleLoad{
				name: "singleton",
			}
			time.Sleep(3 * time.Second)
		})
	}
	return singleLoadIns
}

func main() {
	var pt1, pt2 *SingleLoad
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		// pt1 = GetLoadInstance1()
		// pt1 = GetLoadInstance2()
		// pt1 = GetLoadInstance3()
		pt1 = GetLoadInstance4()
	}()

	go func() {
		defer wg.Done()
		// pt2 = GetLoadInstance1()
		// pt2 = GetLoadInstance2()
		// pt2 = GetLoadInstance3()
		pt2 = GetLoadInstance4()
	}()

	wg.Wait()
	if pt1 == pt2 {
		fmt.Println("pt1 == pt2") // always here
	} else {
		fmt.Println("pt1 != pt2")
	}
	fmt.Println("end")
}
