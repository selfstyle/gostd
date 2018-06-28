package main

import (
	"flag"
	"fmt"
	"testing"
)

func Test_flag_(t *testing.T) {
	// cmd -flag   只支持bool类型
	// cmd -flag=x
	// cmd -flag x 只支持非bool类型
	//参数有三个：第一个为参数名称，第二个为默认值，第三个是说明
	// (1) 通过flag.String()，Bool()，Int() 等flag.Xxx()方法，该种方式返回一个相应的指针
	username := flag.String("name", "caoyan", "Input your username")
	age := flag.Int("age", 24, "Input your age")
	// (2) 通过flag.XxxVar()方法将flag绑定到一个变量，该种方式返回值类型
	var sex string
	flag.StringVar(&sex, "sex", "男", "man or woman")
	// (3) 通过flag.Var()绑定自定义类型，自定义类型需要实现Value接口(Receives必须为指针)
	//var tel string
	//flag.Var(&tel, "")
	//调用flag.Parse()解析命令行参数到定义的flag
	flag.Parse()
	fmt.Println("Hello, ", *username, *age, sex)
}
