package bst

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}
type BST struct {
	Root *Node
}

func (tree *BST) IsEmpty() bool {
	return tree.Root == nil
}

func (tree *BST) Insert(v int) {
	tree.insert(v, tree.Root)
}

func (tree *BST) Has(v int) bool {
	_, target := tree.FindNode(v)
	return target != nil
}

func (tree *BST) insert(v int, node *Node) {

	if tree.IsEmpty() {
		tree.Root = &Node{Val: v}
		return
	}

	if v < node.Val && node.Left == nil {
		node.Left = &Node{Val: v}
		return
	}
	if v > node.Val && node.Right == nil {
		node.Right = &Node{Val: v}
		return
	}

	if v < node.Val && node.Left != nil {
		tree.insert(v, node.Left)
		return
	}

	if v > node.Val && node.Right != nil {
		tree.insert(v, node.Right)
	}

}

func (tree *BST) Remove(v int) {
	// if we can't find the node, nothing to do ¯\_(ツ)_/¯
	parent, target := tree.FindNode(v)
	if target == nil {
		return
	}
	// if target is a leaf, just remove it: make his parent's Left/Right point to nil
	if tree.isLeaf(target) {
		tree.replaceChild(parent, target, nil)
		return
	}
	// if target has just one child: make its parent point to that unique child

	if target.Left == nil {
		tree.replaceChild(parent, target, target.Right)
		return
	}
	if target.Right == nil {
		tree.replaceChild(parent, target, target.Left)
		return
	}

	// Now that we reached this line, we are sure the target has two children, we need to:
	// __ 1- Find the smallest of its right subtree (or biggest of its left tree)
	// __ 2- Set its value to value of the node with the value we want to remove
	// __ 3- Remove the node found in the step above
	smallestBiggerNodeParent, smallestBiggerNode := target, target.Right
	for ; smallestBiggerNode.Left != nil; smallestBiggerNodeParent, smallestBiggerNode = smallestBiggerNode, smallestBiggerNode.Left {

	}
	target.Val = smallestBiggerNode.Val
	tree.replaceChild(smallestBiggerNodeParent, smallestBiggerNode, nil)

}

func (tree *BST) isLeaf(node *Node) bool {
	return !tree.IsEmpty() && node != nil && node.Left == nil && node.Right == nil
}

func (tree *BST) FindNode(v int) (p *Node, t *Node) {
	var (
		parent *Node = nil
		target       = tree.Root
	)
	for {
		if target == nil {
			return nil, nil
		}

		if target.Val == v {
			return parent, target
		}

		parent = target

		if v < target.Val {
			target = target.Left

			continue

		}

		target = target.Right
	}

}

func (tree *BST) replaceChild(parent *Node, oldChild *Node, newChild *Node) error {
	if parent == nil {
		tree.Root = newChild
		return nil
	}
	if parent.Left == oldChild {
		parent.Left = newChild
		return nil
	}
	if parent.Right == oldChild {
		parent.Right = newChild
		return nil
	}

	return fmt.Errorf("the given parent node and the old child do not match")

}
