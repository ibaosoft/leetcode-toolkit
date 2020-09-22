package binary

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestFormat2Leetcode(t *testing.T) {
	for i, anonymous := range []struct {
		input  []int
		expect string
	}{
		{
			input:  []int{},
			expect: "[]",
		},
		{
			input:  []int{1},
			expect: "[1]",
		},
		{
			input:  []int{1, 2, 3},
			expect: "[1,2,3]",
		},
		{
			input:  []int{1, 2, 3, 4, 5, 6, 7},
			expect: "[1,2,3,4,5,6,7]",
		},
		{
			input:  []int{3, 9, 20, Null, Null, 15, 7},
			expect: "[3,9,20,null,null,15,7]",
		},
		{
			input:  []int{5, 4, 8, 11, Null, 13, 4, 7, 2, Null, Null, 5, 1},
			expect: "[5,4,8,11,null,13,4,7,2,null,null,5,1]",
		},
	} {
		if actual := Format2Leetcode(anonymous.input); actual != anonymous.expect {
			t.Errorf("Format2Leetcode func format result not equal expect, index is %d, input is %v, expect is %s, actual is %s",
				i,
				anonymous.input,
				anonymous.expect,
				actual)
		}
	}
}
func TestUnmarshal(t *testing.T) {
	for i, anonymous := range []struct {
		input  []int
		expect *TreeNode
	}{
		{
			input:  []int{},
			expect: nil,
		},
		{
			input:  []int{1},
			expect: &TreeNode{Val: 1},
		},
		{
			input:  []int{1, 2, 3},
			expect: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
		},
		{
			input:  []int{1, 2, 3, 4, 5, 6, 7},
			expect: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 7}}},
		},
		{
			input:  []int{3, 9, 20, Null, Null, 15, 7},
			expect: &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}},
		},
		{
			input:  []int{5, 4, 8, 11, Null, 13, 4, 7, 2, Null, Null, 5, 1},
			expect: &TreeNode{Val: 5, Left: &TreeNode{Val: 4, Left: &TreeNode{Val: 11, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 2}}}, Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 13}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 1}}}},
		},
	} {
		actual := Unmarshal(anonymous.input)
		actualBytes, _ := json.Marshal(actual)
		expectBytes, _ := json.Marshal(anonymous.expect)
		if string(actualBytes) != string(expectBytes) {
			t.Errorf("Unmarshal func result not equal expect, index is %d, input is %v, expect is %s, actual is %s",
				i,
				anonymous.input,
				expectBytes,
				actualBytes)
		}
	}
}

func TestMarshal(t *testing.T) {
	for i, anonymous := range []struct {
		input  *TreeNode
		expect []int
	}{
		{
			input:  nil,
			expect: []int{},
		},
		{
			input:  &TreeNode{Val: 1},
			expect: []int{1},
		},
		{
			input:  &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
			expect: []int{1, 2, 3},
		},
		{
			input:  &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 7}}},
			expect: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			input:  &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}},
			expect: []int{3, 9, 20, Null, Null, 15, 7},
		},
		{
			input:  &TreeNode{Val: 5, Left: &TreeNode{Val: 4, Left: &TreeNode{Val: 11, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 2}}}, Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 13}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 1}}}},
			expect: []int{5, 4, 8, 11, Null, 13, 4, 7, 2, Null, Null, 5, 1},
		},
	} {
		actual := Marshal(anonymous.input)
		if fmt.Sprint(actual) != fmt.Sprint(anonymous.expect) {
			t.Errorf("Unmarshal func result not equal expect, index is %d, input is %v, expect is %v, actual is %v",
				i,
				anonymous.input,
				anonymous.expect,
				actual)
		}
	}
}

func TestDepth(t *testing.T) {
	for i, anonymous := range []struct {
		input  *TreeNode
		expect int
	}{
		{
			input:  nil,
			expect: 0,
		},
		{
			input:  &TreeNode{Val: 1},
			expect: 1,
		},
		{
			input:  &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
			expect: 2,
		},
		{
			input:  &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 7}}},
			expect: 3,
		},
		{
			input:  &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}},
			expect: 3,
		},
		{
			input:  &TreeNode{Val: 5, Left: &TreeNode{Val: 4, Left: &TreeNode{Val: 11, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 2}}}, Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 13}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 1}}}},
			expect: 4,
		},
	} {
		if depth := Depth(anonymous.input); depth != anonymous.expect {
			t.Errorf("Depth func result not equal expect, index is %d, input is %v, expect is %d, actual is %d",
				i,
				anonymous.input,
				anonymous.expect,
				depth)
		}
	}
}
