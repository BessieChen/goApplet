package main

import (
	"fmt"
	"io/ioutil"
	"learngo/infra"
	"net/http"
)

func retrieve(url string) []byte{
	//url = "https://www.baidu.com"
	response, err := http.Get(url)
	if err != nil{
		panic(err)
	}
	defer response.Body.Close() 	//程序跑完之后, response需要关闭
	bytes, err2 := ioutil.ReadAll(response.Body) //工程中, 每个函数都要处理出错的情况
	if err2 != nil{
		panic(err)
	}
	return bytes
}

func retrieve2(url string) string{
	//url = "https://www.baidu.com"
	response, err := http.Get(url)
	if err != nil{
		panic(err)
	}
	defer response.Body.Close() 	//程序跑完之后, response需要关闭
	bytes, err2 := ioutil.ReadAll(response.Body) //工程中, 每个函数都要处理出错的情况
	if err2 != nil{
		panic(err)
	}
	return string(bytes)
}

func getRetriever() infra.Retriever{	//返回的是类型
	return infra.Retriever{}
	//return testing.Retriever{} 不能这么写
}

func getRetrieverInterface() retrieverInterface{	//返回的是类型
	return infra.Retriever{} 	//之后需要谁家的实现{不管是infra还是testing}, 就直接改这里的
	//return testing.Retriever{}
}

type retrieverInterface interface{
	Get(string) string
}

func main() {
	//fmt.Printf("%s\n", retrieve("https://www.baidu.com"))
	//fmt.Println(retrieve2("https://www.baidu.com"))

	//通过调用的形式:
	//r := infra.Retriever{} 	//package 名 + struct 名
	//fmt.Println(r.Get("https://www.baidu.com"))
	//
	//var r2 infra.Retriever
	//r2 = getRetriever()	//这样就解耦了, retriever 的类型
	//fmt.Println(r2.Get("https://www.baidu.com"))

	//假设我们想使用 testing.Retriever, 需要把所有的 infra.Retriever 都改成 testing.Retriever, 都很麻烦
	var r3 retrieverInterface = getRetrieverInterface()
	fmt.Println(r3.Get("https://www.baidu.com"))
}

