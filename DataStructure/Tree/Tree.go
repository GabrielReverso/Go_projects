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

//Create a tree
func Tree_create() *Tree {
	new_tree := &Tree{root: nil}
	return new_tree
}

//Add a value in the tree
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
	default:
		if current.left == nil {
			current.left = new_node
		} else if current.right == nil {
			current.right = new_node
		} else {
			add_node(current.left, value, new_node)
			if current.left == nil {
				current.left = new_node
			}
		}
	}
}

//Print in pre-order
func Tree_pre_order(tree *Tree) {
	if tree.root == nil {
		fmt.Println("Empty list")
		return
	}
	fmt.Printf("Pre-order\n")
	fmt.Printf("[ ")
	pre_order(tree.root)
	fmt.Printf("]\n\n")
}

func pre_order(current *Node) {
	if current == nil {
		return
	}

	fmt.Printf("%v ", current.data)
	pre_order(current.left)
	pre_order(current.right)
}

//Print in order
func Tree_order(tree *Tree) {
	if tree.root == nil {
		fmt.Println("Empty list")
		return
	}
	fmt.Printf("Order\n")
	fmt.Printf("[ ")
	order(tree.root)
	fmt.Printf("]\n\n")
}

func order(current *Node) {
	if current == nil {
		return
	}
	order(current.left)
	fmt.Printf("%v ", current.data)
	order(current.right)
}

//Print in post-order
func Tree_post_order(tree *Tree) {
	if tree.root == nil {
		fmt.Println("Empty list")
		return
	}
	fmt.Printf("Post-order\n")
	fmt.Printf("[ ")
	post_order(tree.root)
	fmt.Printf("]\n\n")
}

func post_order(current *Node) {
	if current == nil {
		return
	}
	post_order(current.left)
	post_order(current.right)
	fmt.Printf("%v ", current.data)
}
