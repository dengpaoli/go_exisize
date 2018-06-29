package main

import (
	"time"
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(10)
	go func() {
		time.Sleep(1000)
		fmt.Println("go func")
		//wg.Done() //或者
		wg.Add(-10)
	}()
	fmt.Println("main")
	wg.Wait()
}
