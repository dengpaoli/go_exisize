package main

import (
	"time"
	"fmt"
)

func main() {

	mychan := make(chan int,1)
	go func() {
		time.Sleep(time.Second  *2)
		println("go func")
		mychan<-1
	}()

	fmt.Println("main")
	<-mychan
}
