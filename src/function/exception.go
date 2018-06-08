package main

import (
	"fmt"
	"runtime/debug"
)

func CheckPanic() {

	if r := recover(); r != nil {
		fmt.Println("catch an panic %v", r)
		debug.PrintStack()
	}
}

func myPainc() {
	var x = 30
	var y = 0
	var c = x / y
	fmt.Println(c)
}

func main() {
	fmt.Printf("Hello World!")
	defer CheckPanic()
	myPainc()
	fmt.Printf("这里执行不到!")
}
