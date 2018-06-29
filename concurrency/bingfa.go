package main

import (
	"fmt"
	"time"
)

//gorouting = 协程（一个cpu） 或 线程goruntime

func task1()  {
	for ; ; {
		fmt.Println("task1")
		time.Sleep(2*time.Second)
	}
}

func task2()  {
	for ; ;  {
		fmt.Println("task2")
		time.Sleep(2*time.Second)
	}
}

func main() {
	fmt.Println("gorouting")
	go task1() //gorouting
	go task2()
	go func() {
		time.Sleep(23)
		fmt.Println("匿名gorouting")
	}() //匿名gorouting
	task2()

	//假如主线程结束，则子的goruting结束
}
