package main

import "fmt"

func main() {
	switch_test(1)
	switch_test(2)
}

func switch_test(a int)  {
	switch a {
	case 1,2,3,4:
		fmt.Println("true")
		fallthrough //继续执行后面的
	//不需要break
	case 0:
		fmt.Println("false")
		fallthrough //继续执行后面的
	default:
		fmt.Println("unknown")
	}
}
