package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

//sync.once可以控制函数只能被调用一次。不能多次重复调用。

func do(o *sync.Once) {
	fmt.Println("Start do")
	o.Do(func() {
		fmt.Println("Doing something...")
	})
	fmt.Println("Do end")
}

func Test_sync_once(t *testing.T) {
	once := &sync.Once{}
	go do(once)
	go do(once)
	time.Sleep(time.Second * 2)
}
