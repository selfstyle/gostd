package main

import (
	"fmt"
	"net/smtp"
	"runtime"
	"strings"
	"testing"
)

func SendEmail(subject, body string) error {
	user := "caoyan@formaxmarket.com"
	password := "8FX7qb8XxojEz9y5"
	host := "smtp.exmail.qq.com:25"
	to := "caoyan@formaxmarket.com"
	//contenttype := "html"

	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	contenttype := "Content-Type: text/html" + "; charset=UTF-8"

	msg := []byte("To: " + to + "\r\nFrom: " + user + "\r\nSubject: " +
		subject + "\r\n" + contenttype + "\r\n\r\n" + body)

	fmt.Println("msg:", string(msg))
	sendto := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, sendto, msg)
	return err
}

func Test_email(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	err := SendEmail("测试", "我正在测试")
	fmt.Println("err: ", err)
}
