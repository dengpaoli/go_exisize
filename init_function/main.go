package main

import (
	"fmt"
	"log"
	"os"
	"github.com/dengpaoli/go_exisize/init_function/child" //如果子包有init函数，优先执行
)


func init() {
	log.SetOutput(os.Stdout) //这个很重要，没有日志打印有问题
	fmt.Println("init")
}

func main() {
	fmt.Println("main")
	child.Run() //如果子包有init函数，优先执行

	file,err := os.Open("hell.txt")
	if err != nil {
		fmt.Println("open faild")
		return
	}

	buf := make([]byte,256)
	s,es := file.Read(buf )
	if s<0 || es != nil{
		fmt.Println("read ")
		return
	}
}
