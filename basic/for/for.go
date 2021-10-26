package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func convertToBinary(n int) string {
	res := ""
	if n <= 0 {
		return res
	}
	for ; n > 0; n >>= 1 {
		low := n & 1 //low := n % 2
		res = strconv.Itoa(low) + res
	}
	return res
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func foreverLoop() {
	for {
		fmt.Println("abc")
	}
}

func main() {
	fmt.Println(convertToBinary(8), convertToBinary(16), convertToBinary(13))
	printFile("abc.txt")

}
