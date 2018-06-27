package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_reflect_(t *testing.T) {
	var circle float64 = 6.28
	var icir interface{}
	icir = circle
	fmt.Println("icir.Value = ", reflect.ValueOf(icir))
	fmt.Println("icir.Type = ", reflect.TypeOf(icir))
}
