package main

import "fmt"

type parent_ interface {
	test() int8
}

type childa interface {
	parent_
	test2() string
}

//定义一个类型
type data struct {

}

//接口属于类型
func (p *data)test() int8 {
	fmt.Println(8)
	return 8
}
func (c *data)test2() string {
	fmt.Println("string")
	return "string"
}

func main() {
	var f *data
	//var c *childa
	f.test()
	f.test2()

}
