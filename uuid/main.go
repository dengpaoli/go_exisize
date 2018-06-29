package main

import (
	"github.com/satori/go.uuid"
	"fmt"
)

func main() {
	u1,_ := uuid.NewV4()
	fmt.Println(u1)
}
