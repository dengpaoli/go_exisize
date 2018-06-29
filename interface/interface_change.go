package main

import "fmt"

func main() {

	t := 3
	var f float64 = float64(t)
	fmt.Println(f)

	//万能数据类型
	var i interface{}
	i = 4
	fmt.Println(i)

	i = "abv"
	fmt.Println(i)

	m := make(map[string]string,0)
	m["name"] = "xiaoming"
	m["age"] = "young"
	//fmt.Println(m)

	i = m
	fmt.Println(i)

	//类型断言
	n := make(map[string]interface{},0)
	n["int"] = 1
	n["string"] = "asfasf"
	var d1 int
	var ss string

	d1 = n["int"].(int) //类型断言
	ss = n["string"].(string)
	fmt.Println(d1,ss)

}
