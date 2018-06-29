package main

import (
	"runtime"
	"fmt"
)

func main() {
	//gosched把调度器让个别人
	runtime.GOMAXPROCS(1)
	exit := make(chan int)
	go func() {
		defer func() {
			exit<-1
		}()
		go func() {
			fmt.Println("b")
		}()
		fmt.Println("a1")
		runtime.Gosched()
		fmt.Println("a2")
	}()

	<-exit
}
