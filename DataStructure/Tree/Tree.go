package tree

import "fmt"

type Node struct {
	left  *Node
	data  int
	right *Node
}

type Tree struct {
	root *Node
}

func Tree_create() *Tree {
	new_tree := &Tree{root: nil}
	return new_tree
}

func Tree_add(tree *Tree, value int) {
	new_node := &Node{left: nil, data: value, right: nil}

	if tree.root == nil {
		tree.root = new_node
	} else {
		current := tree.root
		add_node(current, value, new_node)
	}

}

func add_node(current *Node, value int, new_node *Node) {
	if current == nil {
		return
	}

	switch {
	case value < current.data:
		add_node(current.left, value, new_node)
		if current.left == nil {
			current.left = new_node
		}
	case value > current.data:
		add_node(current.right, value, new_node)
		if current.right == nil {
			current.right = new_node
		}
	}
}

func Tree_pre_order(tree *Tree) {
	if tree.root == nil {
		fmt.Println("Empty list")
		return
	}
	pre_order(tree.root)
}

func pre_order(current *Node) {
	if current == nil {
		return
	}

	fmt.Println(current.data)
	pre_order(current.left)
	pre_order(current.right)
}
