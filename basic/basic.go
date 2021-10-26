package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var aa = 3

//错误: bb := 4, 因为它不能在函数外使用

//2-2
//var (
//	b = 3
//	d = "def"
//	c = false
//)

func varaibleZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableIntialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, d = 3, 4, true, "def"
	fmt.Println(a, b, c, d)
}

func variableShort() {
	a, b, c, d := 3, 4, true, "def"
	b = 5
	fmt.Println(a, b, c, d)
}

func euler() {
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))

	fmt.Println(1 + cmplx.Pow(math.E, 1i*math.Pi)) //e^i*pi + 1 == 0
	fmt.Println(1 + cmplx.Exp(1i*math.Pi))         //e^i*pi + 1 == 0
	fmt.Printf("%.2f\n", 1+cmplx.Exp(1i*math.Pi))  //e^i*pi + 1 == 0
}

func triangle() {
	var a, b int = 3, 4
	var c int
	//c = math.Sqrt(a + b)
	//c = math.Sqrt(float64(a + b))
	c = int(math.Sqrt(float64(a + b)))
	fmt.Println(c)
}

func consts() {
	const filename = "abc.txt"
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a + b))
	fmt.Println(filename, c)
}

func enums() {
	const (
		cpp    = 0
		java   = 1
		golang = 2
	)
	fmt.Println(cpp, java, golang)
}

func enums2() {
	const (
		cpp = iota
		_
		python
		golang
		java
	)

	const (
		b = 1 << (1 * iota)
		c
		d
		e
	)
	fmt.Println(cpp, java)
	fmt.Println(b, c, d, e)
}

func main() {
	fmt.Println("hello world")
	//varaibleZeroValue()
	//variableIntialValue()
	//variableTypeDeduction()
	//variableShort()
	//
	//euler()
	//triangle()

	enums2()
}
