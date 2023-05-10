package main

import (
	"fmt"
	"time"
)

func Test01() {
	go sayHello("sayHello")
	for i := 0; i < 10; i++ {
		fmt.Printf("main 线程=%d  \n", i)
	}
	time.Sleep(100000)
}
func sayHello(str string) {
	for i := 0; i < 10; i++ {
		fmt.Println(str, "=", i)
	}
}
