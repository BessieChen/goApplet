package main

import (
	"fmt"
	"learngo/retriever/mock"
	real2 "learngo/retriever/real"
	"time"
)

//以下都是接口 -----------
type RetrieverInterface interface{
	Get(url string) string	//不需要添加 func 关键字, 因为 interface 里面都是函数
}

type PosterInterface interface{
	Post(url string, form map[string] string) string
}

//这个就是接口的组合
type RetrieverAndPosterInterface interface {
	Get(url string) string
	Post(url string, form map[string]string) string
}
//以上都是接口 -----------

const url = "https://www.baidu.com"
//以下是调用接口的函数, 其实就是填写参数是什么 -------
func download(r RetrieverInterface) string {
	return r.Get(url)
}

func post(p PosterInterface) {
	p.Post("https://www.baidu.com",
		map[string]string{			//这里相当于是直接在 括号里面生成了参数 map
		"name":"bessie",
		"course":"learngo",
		})
}

func session(rp RetrieverAndPosterInterface) string {
	rp.Post(url, map[string]string{"contents": "fake com"})
	return rp.Get(url)
}
//以上是调用接口的函数, 其实就是填写参数是什么 -------








func main() {
	//var r Retriever
	//fmt.Println(download(r))	//出错, 因为不知道 r 是什么

	//
	//var r RetrieverInterface
	//r = mock.Retriever{"this is a fake .com"}
	//fmt.Println(download(r))
	//
	//fmt.Println(download(mock.Retriever{"this is fake"}))
	//r = real2.Retriever{}	//好神奇, 我定义的是 real, 他要我写的是 real2
	//fmt.Println(download(r))

	var r RetrieverInterface
	r = &mock.Retriever{"this is a fake .com"}	//这个是拷贝, 把 mock.Retriver 这个对象拷贝到了 r 里面
	fmt.Printf("> %T %v\n", r, r)
		//T: type{类型}, v: value{值}, r里面的东西
		//类型: mock.Retriever
		//值: {this is a fake .com}
	inspect(r)
	//type assertion: 类型的检查
	if mockRetriever, ok := r.(*mock.Retriever); ok{		//之前说了 r 里面是拷贝, 市一大块内存, 这个内存里面就有一个 mock.Retriever 对象
		fmt.Println("is mock retriever, ", mockRetriever.Content)
	}else{
		fmt.Println("is mock retriever")
	}

	r = real2.Retriever{		//这个是拷贝, 把 real.Retriver 这个对象拷贝到了 r 里面
		UserAgent: "Mozilla/5.0",
		Timeout: time.Minute,
	}	//好神奇, 我定义的是 real, 他要我写的是 real2
	fmt.Printf("> %T %v\n", r, r)
		//real.Retriever
		//{Mozilla/5.0 1m0s}
	inspect(r)
	if raalRetriever, ok := r.(real2.Retriever); ok{		//之前说了 r 里面是拷贝, 市一大块内存, 这个内存里面就有一个 mock.Retriever 对象
		fmt.Println("is real retriever, ", raalRetriever.UserAgent)
	}else{
		fmt.Println("is real retriever")
	}

	r = &real2.Retriever{		//这个是指针, 把 real.Retriver 这个对象的地址, 给了 r, 其实 r 就不是装很多东西的内存了, 仅仅是地址
		UserAgent: "Mozilla/5.0",
		Timeout: time.Minute,
	}	//好神奇, 我定义的是 real, 他要我写的是 real2
	fmt.Printf("%T %v\n", r, r)
		//*real.Retriever
		//&{Mozilla/5.0 1m0s}
	inspect(r)
	if raalRetriever, ok := r.(real2.Retriever); ok{		// r 只是一个指针, 既然是指针, 就是存着目标对象的地址, 即: &目标对象
		fmt.Println("is real retriever, ", raalRetriever.UserAgent)
	}else{
		fmt.Println("is real retriever pointer")
	}
	if raalRetriever, ok := r.(*real2.Retriever); ok{		// r 只是一个指针, 既然是指针, 就是存着目标对象的地址, 即: &目标对象. 但是这么写是错的, 正确的写法: r是目标对象的指针, 即 *real2.Retriever. 适应吧
		fmt.Println("is real retriever, ", raalRetriever.UserAgent)
	}else{
		fmt.Println("is real retriever pointer")
	}

	fmt.Println(session(&mock.Retriever{Content: "test rp session()"})) //这个没错, 因为 mock 包, 实现了 Post() 和 Get(), 但是没有修改值. 因为是值接受者, 修改为指针接受者就可以修改值
	//fmt.Println(session((real2.Retriever{"test rp session()"}))) // 这一句会报错, 因为 real2 这个包, 没有实现 Post()


}

func inspect(r RetrieverInterface) {
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents: ", v.Content)
	case real2.Retriever:
		fmt.Println("UserAgent: ", v.UserAgent)
	case *real2.Retriever:
		fmt.Println("UserAgent: ", v.UserAgent)
	}
}