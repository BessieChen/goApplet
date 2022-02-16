package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
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
	reader, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	printFileContents(reader)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
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
	//printFile("abc.txt")

	s := `'
	rjelwkjepidojc
	rejoijc`

	printFileContents(strings.NewReader(s)) //所以我们不仅可以打印 .txt, 还可以打印各种string

}
