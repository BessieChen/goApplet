package main

import "fmt"

func grade(g int) string {
	res := ""
	switch {
	case g < 0:
		panic(fmt.Sprintf("Wrong score: %d", g))
	case g < 10:
		res = "a"
	case g < 20:
		res = "b"
	case g < 100:
		res = "c"
	default:
		res = "d"
	}
	return res
}

func main() {
	fmt.Println(grade(101))
	fmt.Println(grade(-1))
}
