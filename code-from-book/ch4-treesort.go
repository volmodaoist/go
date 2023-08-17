package main

import "fmt"

type TreeNode struct {
	value       int
	left, right *TreeNode
}

func Sort(value []int) []int {
	var root *TreeNode
	for _, v := range value {
		root = tree_add(root, v)
	}
	return appendValue(value[:0], root)
}

func tree_add(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{value: val}
	}
	if val < root.value {
		root.left = tree_add(root.left, val)
	} else {
		root.right = tree_add(root.right, val)
	}
	return root
}

func appendValue(values []int, root *TreeNode) []int {
	if root != nil {
		values = appendValue(values, root.left)
		values = append(values, root.value)
		values = appendValue(values, root.right)
	}
	return values
}

func main() {
	values := [...]int{3, 2, 4, 6, 9, 2, 1}
	sorted_values := Sort(values[:])
	fmt.Println(sorted_values)
	for _, v := range sorted_values {
		fmt.Println(v)
	}
}
