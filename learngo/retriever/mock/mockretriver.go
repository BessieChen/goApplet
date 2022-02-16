package mock

import "fmt"

//不需要叫做 mockRetriever, 因为之后调用的时候就已经是 mock.Retriever
type Retriever struct {
	Content string
}

func (r *Retriever) Get(url string) string {
	return r.Content
}

func (r *Retriever) Post(url string,
	form map[string]string) string {
	r.Content = form["contents"]
	return "ok"
}

// 之后打印 mock.Retriever 对象, 就是输出下面的内容
func (r *Retriever) String() string{
	return fmt.Sprintf("Retriever: {Contents=%s\n} ", r.Content)
}

