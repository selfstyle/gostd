package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"runtime"
	"testing"
	"time"
)

func Test_pkg_fmt(t *testing.T) {
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}

func Test_pkg_math(t *testing.T) {
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println(math.Pi)
}

//func add(x int, y int) int {
//	return x + y
//}
func add(x, y int) int {
	return x + y
}

func Test_func_add(t *testing.T) {
	fmt.Println(add(1, 2))
}

//func swap(x, y string) (string, string) {
//	return y, x
//}
func swap(x, y string) (a, b string) {
	b = x
	a = y
	return
}

func Test_func_swap(t *testing.T) {
	fmt.Println(swap("hello", "world"))
}

//---------------------------------------------------------------------

var c, python, java bool

func Test_var_(t *testing.T) {
	var i int
	fmt.Println(i, c, python, java)

}

var i, j int = 1, 2

func Test_var_init(t *testing.T) {
	var c, python, java = true, false, "no!"
	k := 3
	fmt.Println(i, j, k, c, python, java)

}

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func Test_var_type(t *testing.T) {
	const f = "%T(%v)\n"
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)
}

const Pi = 3.14

func Test_var_const(t *testing.T) {
	const World = "世界"
	fmt.Println("Hello", "World")
	fmt.Println("Happy", Pi, "Day")
	const Truth = true
	fmt.Println("Go rules?", Truth)
}

func Test_for_(t *testing.T) {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func Test_for_sum(t *testing.T) {
	sum := 1 //这里一定要从1开始 ????
	for sum < 100 {
		sum += sum
	}
	fmt.Println(sum)
}

func Test_for_death(t *testing.T) {
	for {

	}
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))

}
func Test_if_(t *testing.T) {
	fmt.Println(sqrt(2), sqrt(-4))

}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func Test_if_else(t *testing.T) {
	fmt.Println(pow(3, 2, 10), pow(3, 3, 20))
}

func Test_switch_(t *testing.T) {
	fmt.Print("Go runs on: ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Println("%S.", os)
	}

}

func Test_switch_true(t *testing.T) {
	h := time.Now()
	switch {
	case h.Hour() < 12:
		fmt.Println("Good morning!")
	case h.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func Test_defer_(t *testing.T) {
	defer fmt.Println("world")
	fmt.Println("hello")
}

func Test_defer_for(t *testing.T) {
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}

//---------------------------------------------------------------------

func Test_pointer_(t *testing.T) {
	i, j := 42, 2701
	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)
	p = &j
	*p = *p / 37
	fmt.Println(j)

}

type Vertex struct {
	X int
	Y int
}

func Test_struct_(t *testing.T) {
	fmt.Println(Vertex{1, 2})

}

func Test_struct_access(t *testing.T) {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}

func Test_struct_pointer(t *testing.T) {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}

func Test_struct_syntax(t *testing.T) {
	v1 := Vertex{1, 2}
	v2 := Vertex{X: 1}
	v3 := Vertex{}
	p := &Vertex{1, 2}

	fmt.Println(v1, p, v2, v3)
}

//类型 [n]T 是一个有 n 个类型为 T 的值的数组
func Test_array_(t *testing.T) {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
}

//slice 可以重新切片，创建一个新的 slice 值指向相同的数组。
//s[lo:hi] 表示从 lo 到 hi-1 的 slice 元素，含两端
func Test_array_slice(t *testing.T) {
	p := []int{
		2, 3, 5, 7, 11, 13}
	fmt.Println("p ==", p)
	fmt.Println("p[1:4] ==", p[1:4])
	// 省略下标代表 从 0 开始
	fmt.Println("p[:3] ==", p[:3])
	// 省略上标代表到 len(s)结束
	fmt.Println("p[4:] ==", p[4:])
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

//slice 由函数 make 创建。这会分配一个零长度的数组并且返回一个 slice 指向这个数组
//a := make([]int, 5)  // len(a)=5
//为了指定容量，可传递第三个参数到 `make`：
//b := make([]int, 0, 5) // len(b)=0, cap(b)=5
func Test_slice_(t *testing.T) {
	a := make([]int, 5)
	printSlice("a", a)
	b := make([]int, 0, 5)
	printSlice("b", b)
	c := b[:2]
	printSlice("c", c)
	d := c[2:5]
	printSlice("d", d)
}

func Test_slice_nil(t *testing.T) {
	var z []int
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("nil!")
	}
}

//向 slice 添加元素是一种常见的操作，因此 Go 提供了一个内建函数 `append`。 内建函数的文档对 append 有详细介绍。
//func append(s []T, vs ...T) []T
//append 的第一个参数 s 是一个类型为 T 的数组，其余类型为 T 的值将会添加到 slice。
//append 的结果是一个包含原 slice 所有元素加上新添加的元素的 slice。
//如果 s 的底层数组太小，而不能容纳所有值时，会分配一个更大的数组。 返回的 slice 会指向这个新分配的数组。
func Test_slice_append(t *testing.T) {
	var a []int
	printSlice("a", a)
	//append works on nil slices.
	a = append(a, 0)
	printSlice("a", a)
	// the slice grows as needed.
	a = append(a, 1)
	printSlice("a", a)

	a = append(a, 2, 3, 4)
	printSlice("a", a)
}

//for 循环的 range 格式可以对 slice 或者 map 进行迭代循环。
func Test_range_(t *testing.T) {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	sum := make([]int, 10)
	for i := range sum {
		sum[i] = 1 << uint(i)
	}
	for _, value := range sum {
		fmt.Printf("%d\n", value)
	}
}

//map 在使用之前必须用 make 而不是 new 来创建；值为 nil 的 map 是空的，并且不能赋值
type Vertexl struct {
	Lat, Long float64
}

func Test_map_(t *testing.T) {
	m := make(map[string]Vertexl)
	m["Bell Labs"] = Vertexl{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}

func Test_map_syntax(t *testing.T) {
	var m = map[string]Vertexl{
		"Bell Labs": Vertexl{
			40.68433, -74.39967,
		},
		"Google": Vertexl{
			37.42202, -122.08408,
		},
	}
	fmt.Println(m)
	var m1 = map[string]Vertexl{
		"Bell Labs": {
			40.68433, -74.39967,
		},
		"Google": {
			37.42202, -122.08408,
		},
	}
	fmt.Println(m1)
}

//插入或修改一个元素：m[key] = elem
//获得元素：elem = m[key]
//删除元素：delete(m, key)
//通过双赋值检测某个键存在：elem, ok = m[key]
//如果 key 在 m 中，`ok` 为 true 。否则， ok 为 `false`，并且 elem 是 map 的元素类型的零值。
//同样的，当从 map 中读取某个不存在的键时，结果是 map 的元素类型的零值。
func Test_map_operate(t *testing.T) {
	m := make(map[string]int)
	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])
	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

}

//函数也是值。
func Test_func_(t *testing.T) {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(3, 4))
}

//Go 函数可以是闭包的。闭包是一个函数值，它来自函数体的外部的变量引用。
//函数可以对这个引用值进行访问和赋值；换句话说这个函数被“绑定”在这个变量上。
//例如，函数 adder 返回一个闭包。每个闭包都被绑定到其各自的 sum 变量上。
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func Test_func_closure(t *testing.T) {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}

}

/*
func fibonacci() func() int {
}

func Test_func_fibonacci(t *testing.T) {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
*/

//-----------------------------------------------------------------

type Vertexf struct {
	X, Y float64
}

func (v *Vertexf) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//Go 没有类。然而，仍然可以在结构体类型上定义方法。
func Test_method_(t *testing.T) {
	v := &Vertexf{
		3, 4,
	}
	fmt.Println(v.Abs())
}

//你可以对包中的 任意 类型定义任意方法，而不仅仅是针对结构体。
//但是，不能对来自其他包的类型或基础类型定义方法

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func Test_method_float(t *testing.T) {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
