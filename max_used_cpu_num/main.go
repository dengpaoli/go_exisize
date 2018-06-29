package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.NumCPU()) //cpu的个数
	runtime.GOMAXPROCS(runtime.NumCPU()) //设置运行时暂用的cpu个数
}
