package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[1:]) // args 는 슬라이스 타입, Args == arguments
	fmt.Println(os.Args[2])
}
