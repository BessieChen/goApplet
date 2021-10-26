package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	default:
		panic("unsupported operation" + op)
	}
}

func evalPanic(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	default:
		//panic("unsupported operation" + op)
		return 0, fmt.Errorf("unsupported operation: %s\n", op)
	}
}

func div(a, b int) (int, int) {
	return a / b, a % b
}

func div2(a, b int) (q, r int) {
	q = a / b
	r = a % b
	//return q, r 	也是可以的
	return
}

func div3(a, b int) (q, r int) {
	return a / b, a % b
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("print %s(%d, %d)\n", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func main() {
	a, b := div(3, 5)
	fmt.Println(a, b)

	a, b = div2(4, 6)
	fmt.Println(a, b)

	_, b = div3(6, 9)
	fmt.Println(b)

	if res, err := evalPanic(1, 2, "x"); err == nil {
		fmt.Println(res)
	} else {
		fmt.Println(err)
	}
	d := apply(pow, 1, 2)
	fmt.Println(d)

	d = apply(func(a int, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 3, 4)
	fmt.Println(d)

	fmt.Println(sum(1, 2, 3, 4, 5))
}
