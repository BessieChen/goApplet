package tree

import "fmt"

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

//自定义工厂函数
func createNode(value int) *TreeNode {
	return &TreeNode{value, nil, nil} //cpp中返回的是局部变量的地址, 这个是典型错误. 但是go可以, 老师没有解释为什么
}

func main() {
	var root TreeNode
	root = TreeNode{Value: 3}
	root.Left = &TreeNode{}
	root.Right = &TreeNode{2, nil, nil}
	root.Right.Left = new(TreeNode) //new出来的东西, 本身返回的就是指针

	//创建 treeNode切片
	nodes := []TreeNode{
		{Value: 3},
		{},
		{Value: 3, Left: nil, Right: &root},
		{Value: 4, Left: nil, Right: &root},
	}
	fmt.Println(nodes)

	var root2 *TreeNode
	root2 = createNode(3)
	fmt.Println(root2)
}
