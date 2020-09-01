package main

import (
	"fmt"
)

type node struct {
	val   int
	left  *node
	right *node
}

type bTree struct {
	root *node
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func newNode(val int) *node {
	return &node{val, nil, nil}
}

func inorder(n *node) {
	if n != nil {
		inorder(n.left)
		fmt.Print(n.val, " ")
		inorder(n.right)
	}
}

func insert(root *node, val int) *node {
	if root == nil {
		return newNode(val)
	}

	if val < root.val {
		root.left = insert(root.left, val)
	} else {
		root.right = insert(root.right, val)
	}

	return root
}

/*
 *
 */

func inorderSuccessor(root *node) *node {
	cur := root
	for cur.left != nil {
		cur = cur.left
	}
	return cur
}

func bstDelete(root *node, val int) *node {
	if root == nil {
		return nil
	}

	if val < root.val {
		root.left = bstDelete(root.left, val)
	} else if val > root.val {
		root.right = bstDelete(root.right, val)
	} else if root.val == val {
		// NOTE:This is the node to delete.
		// node with one child
		if root.left == nil {
			root = root.right
		} else if root.right == nil {
			root = root.left
		} else {
			// node with two children
			d := inorderSuccessor(root)
			root.val = d.val
			root.right = bstDelete(root.right, val)
		}
	}

	return root
}

/* 递归计算数的深度
 *
 *
 */
func calculateDepth(n *node, depth int) int {
	if n == nil {
		return depth
	}
	return max(calculateDepth(n.left, depth+1), calculateDepth(n.right, depth+1))
}

func (t *bTree) depth() int {
	return calculateDepth(t.root, 0)
}

func main() {
	t := &bTree{nil}
	inorder(t.root)
	t.root = insert(t.root, 10)
	t.root = insert(t.root, 20)
	t.root = insert(t.root, 15)
	t.root = insert(t.root, 30)
	t.root = insert(t.root, 1)
	fmt.Print(t.depth(), "\n")
	inorder(t.root)
	fmt.Print("\n")
	t.root = bstDelete(t.root, 10)
	inorder(t.root)
	fmt.Print("\n")
	t.root = bstDelete(t.root, 30)
	inorder(t.root)
	fmt.Print("\n")
	t.root = bstDelete(t.root, 15)
	inorder(t.root)
	fmt.Print("\n")
	t.root = bstDelete(t.root, 20)
	inorder(t.root)
	fmt.Print("\n")
	fmt.Print(t.depth(), "\n")
}
