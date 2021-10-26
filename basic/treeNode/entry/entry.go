package main

import (
	tree "basic/treeNode"
	"fmt"
)

func main() {
	var root tree.TreeNode //记得要手动写 package 的名字: tree.
	root = tree.TreeNode{Value: 2}
	root.Left = &tree.TreeNode{1, nil, nil}
	fmt.Println(root)

}
