package binary

import (
	"bytes"
	"container/list"
	"fmt"
	"math"
)

// TreeNode Definition for a binary tree node
type TreeNode struct {
	Val   int       `json:"val"`
	Left  *TreeNode `json:"left,omitempty"`
	Right *TreeNode `json:"right,omitempty"`
}

// Null 自定义类型`Null[MinInt32]`标识的元素解释为nil节点
const Null int = math.MinInt32

// Unmarshal 将整形切片反序列化为标准二叉树结构对象，其中自定义类型`Null[MinInt32]`标识的元素解释为nil节点
//
// 输入：nums = [5,4,8,11,null,13,4,7,2,null,null,5,1]
//
// 输出（图形化结构）:
//
//                5
//              /   \
//             4     8
//            / \   / \
//           11 nl 13  4
//          / \   / \  / \
//         7   2 nl nl 5  1
//
func Unmarshal(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	root := new(TreeNode)
	root.Val = nums[0]

	queue := list.New()
	queue.PushFront(root)

	for i := 0; queue.Len() > 0 && i < len(nums); {
		element := queue.Back()
		queue.Remove(element)

		cur := element.Value.(*TreeNode)
		if cur.Val == Null {
			continue
		}

		if i++; i < len(nums) {
			left := new(TreeNode)
			left.Val = nums[i]
			if left.Val != Null {
				cur.Left = left
			}
			queue.PushFront(left)
		}

		if i++; i < len(nums) {
			right := new(TreeNode)
			right.Val = nums[i]
			if right.Val != Null {
				cur.Right = right
			}
			queue.PushFront(right)
		}
	}

	return root
}

// Marshal 将标准二叉树结构对象序列化为整形切片，其中nil节点会按自定义类型`Null[MinInt32]`标识
func Marshal(root *TreeNode) (res []int) {
	depth := Depth(root)

	var dfs func(depth int, nodes ...*TreeNode)
	dfs = func(curDepth int, nodes ...*TreeNode) {
		if curDepth == 1 && len(nodes) == 1 && nodes[0] == nil {
			return
		}
		if len(nodes) <= 0 {
			return
		}

		nexts := []*TreeNode{}
		for _, node := range nodes {
			if node == nil {
				res = append(res, Null)
				continue
			}

			res = append(res, node.Val)
			if curDepth < depth {
				nexts = append(nexts, node.Left)
				nexts = append(nexts, node.Right)
			}
		}
		dfs(curDepth+1, nexts...)
	}

	dfs(1, root)
	return
}

// Depth 获取二叉树最大深度
func Depth(root *TreeNode) int {
	var depthF func(root *TreeNode) float64
	depthF = func(root *TreeNode) float64 {
		if root == nil {
			return 0
		}

		return math.Max(depthF(root.Left), depthF(root.Right)) + 1
	}

	return int(depthF(root))
}

// Format2Leetcode 根据 Leetcode 测试用例的输入/输出数据格式进行格式化
func Format2Leetcode(nums []int) string {
	buffer := bytes.NewBuffer(nil)
	buffer.WriteByte('[')
	for i, num := range nums {
		if num == math.MinInt32 {
			buffer.WriteString("null")
		} else {
			buffer.WriteString(fmt.Sprint(num))
		}
		if i < len(nums)-1 {
			buffer.WriteByte(',')
		}
	}
	buffer.WriteByte(']')

	return string(buffer.Bytes())
}
