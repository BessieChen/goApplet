package infra

import (
	"io/ioutil"
	"net/http"
)

type Retriever struct {}

/**
1. 因为这个 struct 里面没有东西, 所以直接写 (Retriever), 不需要变量名
2. 都是大写, 所以是public
 */
func (Retriever) Get(url string) string{
		//url = "https://www.baidu.com".
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


