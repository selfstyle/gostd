package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.NewTimer(time.Second * 2)
	t2 := time.NewTimer(time.Second * 5)

	for {
		select {
		case <-t1.C:
			fmt.Println("2s timer")
			t1.Reset(time.Second * 2)

		case <-t2.C:
			fmt.Println("5s timer")
			t2.Reset(time.Second * 5)
		}
	}
}
