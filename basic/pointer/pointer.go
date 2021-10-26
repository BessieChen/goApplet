package main

import "fmt"

func f1() {
	var a int = 2
	var pa *int = &a
	*pa = 3
	fmt.Println(a)
}

//因为只有值传递, 所以要改变的话, 就要传入指针{传入地址}
func swap(a, b *int) {
	*b, *a = *a, *b
}

func swap2(a, b int) (int, int) {
	return b, a
}

func main() {
	f1()
	a := 1
	b := 2
	fmt.Println(a, b)

	//method1:
	swap(&a, &b)
	fmt.Println(a, b)

	//method2:
	a, b = swap2(a, b)
	fmt.Println(a, b)
}
