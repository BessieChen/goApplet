package main

import (
	"fmt"
	"unicode/utf8"
)

func lengthOfNonRepeating(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0

	for ind, c := range []byte(s) { //把 s 转成 byte
		lastI, exist := lastOccurred[c]
		if exist && lastI >= start {
			start = lastI + 1
		}
		if ind-start+1 > maxLength {
			maxLength = ind - start + 1
		}
		lastOccurred[c] = ind
	}
	return maxLength
}

func lengthOfNonRepeating2(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0

	for ind, c := range []rune(s) { //把 s 转成 byte
		lastI, exist := lastOccurred[c]
		if exist && lastI >= start {
			start = lastI + 1
		}
		if ind-start+1 > maxLength {
			maxLength = ind - start + 1
		}
		lastOccurred[c] = ind
	}
	return maxLength
}

func main() {
	fmt.Println(lengthOfNonRepeating("abcddbc"))
	fmt.Println(lengthOfNonRepeating2("一二三二一"))

	s := "我爱生活"
	for _, ch := range []byte(s) {
		fmt.Printf("%X ", ch)
	}
	fmt.Println()

	bytes := []byte(s)
	for len(bytes) > 0 {
		c, cSize := utf8.DecodeRune(bytes)
		bytes = bytes[cSize:]            //删除第一个字符
		fmt.Printf("%d %c \n", cSize, c) //打印出来 cSize 是 3, 说明中文是3字节
	}

	for i, ch := range []rune(s) {
		fmt.Printf("(%d: %c) ", i, ch)
	}
}
