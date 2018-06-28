package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func Test_bufio_peek(t *testing.T) {
	s := strings.NewReader("ABCDEFG")
	br := bufio.NewReader(s)

	b, _ := br.Peek(5)
	fmt.Printf("%s\n", b)

	b[0] = 'a'
	b, _ = br.Peek(5)
	fmt.Printf("%s\n", b)
}

func Test_bufio_read(t *testing.T) {
	s := strings.NewReader("ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	br := bufio.NewReader(s)
	b := make([]byte, 20)

	n, err := br.Read(b)
	fmt.Printf("%-20s %-2v %v\n", b[:n], n, err)

	n, err = br.Read(b)
	fmt.Printf("%-20s %-2v %v\n", b[:n], n, err)

	n, err = br.Read(b)
	fmt.Printf("%-20s %-2v %v\n", b[:n], n, err)
}

//ReadByte 从 b 中读出一个字节并返回,如果 b 中无可读数据，则返回一个错误
//UnreadByte 撤消最后一次读出的字节,只有最后读出的字节可以被撤消
func Test_bufio_readbyte(t *testing.T) {
	s := strings.NewReader("ABCDEFG")
	br := bufio.NewReader(s)

	c, _ := br.ReadByte()
	fmt.Printf("%c\n", c)

	c, _ = br.ReadByte()
	fmt.Printf("%c\n", c)

	br.UnreadByte()
	c, _ = br.ReadByte()
	fmt.Printf("%c\n", c)

}

func Test_bufio_readrune(t *testing.T) {
	s := strings.NewReader("你好，世界")
	br := bufio.NewReader(s)
	c, size, _ := br.ReadRune()
	fmt.Printf("%c %v\n", c, size)

	c, size, _ = br.ReadRune()
	fmt.Printf("%c %v\n", c, size)

	br.UnreadRune()
	c, size, _ = br.ReadRune()
	fmt.Printf("%c %v\n", c, size)
}

func Test_bufio_buffer(t *testing.T) {
	s := strings.NewReader("你好，世界")
	br := bufio.NewReader(s)
	fmt.Println(br.Buffered())

	br.Peek(1)
	fmt.Println(br.Buffered())
}

func Test_bufio_readslice(t *testing.T) {
	s := strings.NewReader("ABC DEF GHI JKL")
	br := bufio.NewReader(s)
	w, err := br.ReadSlice(' ')
	fmt.Printf("%q\n", w)
	fmt.Println(err)

	w, err = br.ReadSlice(' ')
	fmt.Printf("%q\n", w)
	fmt.Println(err)

	w, err = br.ReadSlice(' ')
	fmt.Printf("%q\n", w)
	fmt.Println(err)

	w, err = br.ReadSlice(' ')
	fmt.Printf("%q\n", w)
	fmt.Println(err)

	w, err = br.ReadSlice(' ')
	fmt.Printf("%q\n", w)
	fmt.Println(err)
}
