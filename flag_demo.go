package main

import (
	"fmt"
	"flag"
)

func main() {
	// flag.String() 最终返回的是指针，*string
	userName := flag.String("name", "", "input your name")
	flag.Parse()
	fmt.Println("hello", *userName)
	//fmt.Println("hello world")
}
