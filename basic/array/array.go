package main

import "fmt"

func printArr(arr []int) {
	arr[0] = 100
	//这个 []int 是切片, 这个就可以真的修改值!!!
}

func printArr2(arr [5]int) {
	arr[0] = 100
	//而且, 即便传入的参数是 [5]arr, 也不是指针, 依旧是值传递{会拷贝!!}, 和cpp完全不同
}

func main() {
	var arr [5]int
	arr2 := [3]int{1, 2, 3}          //如果是冒号, 一定要给具体的数值
	arr3 := [...]int{2, 3, 4, 2, 42} // ...代表编译器帮我们数
	arr4 := []int{2, 3, 4, 2, 42}    //这个是切片
	fmt.Println(arr, arr2, arr3, arr4)

	//var grid = [3][4]int	//错的!!
	var grid [3][4]int //3行4列
	fmt.Println(grid)
	fmt.Println("----------------")

	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}
	fmt.Println("----------------")

	for i := range arr3 {
		fmt.Println(arr3[i])
	}
	fmt.Println("----------------")

	for i, v := range arr3 {
		fmt.Println(i, v)
	}

	fmt.Println("----------------")

	//go 是强类型语言!! 真的 []int 和 [5]int 是完全不同的, 不会说 [5]int就能放在[]int 里面, 依旧放不了!!
	//printArr(arr):错
	//printArr(arr2):错
	//printArr(arr3):错
	printArr(arr4) //这个值被修改了!!

	printArr2(arr) //这个值没有被修改!!
	//printArr2(arr2):错
	printArr2(arr3) //这个值没有被修改!!
	//printArr2(arr4):错

	fmt.Println(arr, arr3, arr4)

}
