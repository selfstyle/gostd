package main

import (
	"fmt"
	"testing"
	"time"
)

func Test_time_date(t *testing.T) {
	d := time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC)
	year, month, day := d.Date()
	hour, min, sec := d.Clock()

	fmt.Println("year:", year)
	fmt.Println("month:", month)
	fmt.Println("day:", day)
	fmt.Println("hour:", hour)
	fmt.Println("min:", min)
	fmt.Println("sec:", sec)
	_parse()
	_time()
}

func _parse() {
	fmt.Println("-----parse------------")
	ti, _ := time.Parse("15:04:05", "01:29:00")
	hour, min, sec := ti.Clock()
	fmt.Println("hour:", hour)
	fmt.Println("min:", min)
	fmt.Println("sec:", sec)
	now := time.Now()
	fmt.Println("now:", now)
	znext := now.Add(time.Hour * 24)
	fmt.Println("znext:", znext)
	zznext := time.Date(znext.Year(), znext.Month(), znext.Day(), 0, 0, 0, 0, znext.Location())
	fmt.Println("zznext:", zznext)
	dura := time.Hour*time.Duration(hour) +
		time.Minute*time.Duration(min) +
		time.Second*time.Duration(sec)
	zzznext := zznext.Add(dura)
	fmt.Println("zzznext:", zzznext)
	fmt.Println("sub:", zzznext.Sub(now))

	/*
		dura := time.Hour*24 +
			time.Hour*time.Duration(hour) +
			time.Minute*time.Duration(min) +
			time.Second*time.Duration(sec)
		addtime := now.Add(dura)
		fmt.Println("addtime:", addtime)
		subtime := addtime.Sub(now)
		fmt.Println("subtime:", subtime)
	*/

	fmt.Println("--------------------")
}

func _time() {
	fmt.Println("------time----------")
	now := time.Now()
	fmt.Println("now:", now)
	next := now.Add(time.Hour * 24)
	fmt.Println("next:", next)
	next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
	fmt.Println("next:", next)
	fmt.Println("sub:", next.Sub(now))
	t := time.NewTimer(next.Sub(now))
	fmt.Println("t:", t)
	fmt.Println("--------------------")
}
