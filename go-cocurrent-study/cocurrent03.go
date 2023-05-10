package main

import (
	"fmt"
	"sync"
	"time"
)

// goroutine 的协调器
var wg sync.WaitGroup

func Test02() {
	wg.Add(1) //Goroutine 的个数
	go sayHello02("sayHello02")
	for i := 0; i < 10; i++ {
		fmt.Printf("main 线程=%d  \n", i)
		time.Sleep(1 * time.Second)
	}
	wg.Wait() //等到计算器=0
}
func sayHello02(str string) {
	defer wg.Done() //每个Goroutine执行完要释放这个计算器
	for i := 0; i < 10; i++ {
		fmt.Println(str, "=", i)
		time.Sleep(2 * time.Second)
	}
}
