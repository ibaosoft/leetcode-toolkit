package binary

import (
	"fmt"
)

func Example() {
	// Leetcode
	// 226. 翻转二叉树
	// 期望输入 [4,2,7,1,3,6,9]
	// 期望输出 [4,7,2,9,6,3,1]

	var invertTree func(root *TreeNode) *TreeNode
	invertTree = func(root *TreeNode) *TreeNode {
		if root == nil {
			return root
		}
		root.Left, root.Right = root.Right, root.Left
		invertTree(root.Left)
		invertTree(root.Right)

		return root
	}

	// 定义输入切片
	nums := []int{4, 2, 7, 1, 3, 6, 9}
	// 实际输入
	actualInput := Format2Leetcode(nums)
	fmt.Println("实际输入：", actualInput)

	// 反序列化为二叉树结构
	root := Unmarshal(nums)
	// 翻转二叉树算法函数实现
	invertRoot := invertTree(root)
	// 序列化二叉树为切片
	invertNums := Marshal(invertRoot)

	// 实际输出
	actualOutput := Format2Leetcode(invertNums)
	fmt.Println("实际输出：", actualOutput)

	// Output:
	//
	// 实际输入： [4,2,7,1,3,6,9]
	// 实际输出： [4,7,2,9,6,3,1]
}
