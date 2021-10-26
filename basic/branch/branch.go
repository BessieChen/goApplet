package main

import (
	"fmt"
	"io/ioutil"
)

func func1() {
	const filename = "abc.txt"
	res, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", res)
	}
}

func func2() {
	const filename = "abc.txt"
	if content, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", content)
	}
}

func main() {

}
