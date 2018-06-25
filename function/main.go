package main

import "fmt"

/*
1 无需前置声明
2 内部定义函数
3 同一类型（参数和返回值相同）可以赋值
4 局部变量可返回
5 函数名清晰可见，不用缩写
*/
func main() {
	a := sayHello
	b := sayHello2

	a = b
	a()

	var x *int
	x = localvarret()

	fmt.Println(*x)
}

func sayHello()  {
	println("hello")

	//2 内部定义函数
	a := func() {fmt.Println(2222)}
	a()
}

func sayHello2()  {
	println("hello")

	//2 内部定义函数
	a := func() {fmt.Println("bbbbb")}
	a()
}

func localvarret() * int  {
	var a int
	a = 1

	return &a //会缓存到heap中，所以保留了
}