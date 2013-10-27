package bst

import (
	"errors"
	"sync"
)

const (
	_ Traversal = iota
	InOrder
	PreOrder
	PostOrder

	_ action = iota
	insertAction
	findAction

	_ childDir = iota
	leftDir
	rightDir
)

type Traversal int
type action int
type childDir int

type VisitFunc func(node *Node)

// A Binary Search Tree
type Tree struct {
	root *Node
	size int
	sync.RWMutex
}

// Create a new tree with the root as the given node
func NewTree(node *Node) *Tree {
	tree := &Tree{root: node, size: 1}
	return tree
}

// Insert a Node into the Tree
func (t *Tree) Insert(node *Node) error {
	t.Lock()
	defer t.Unlock()

	root := t.root
	insertRoot, _ := t.search(nil, root, node)
	if insertRoot == nil { //Found the place to insert the node
		t.size++
		insertRoot = node
		return nil
	}

	// else, there is already a node with that key, see searcht.
	return errors.New("Duplicate node")
}

// Find a Node in the tree, if it isnt found return NotExistError
func (t *Tree) Find(node *Node) *Node {
	t.Lock()
	defer t.Unlock()

	root := t.root
	foundNode, _ := t.search(nil, root, node)
	if foundNode != nil { //We found our node
		return foundNode
	}

	return nil //Node not found
}

// Delete a Node in a tree and rebalance
func (t *Tree) Delete(node *Node) error {
	t.Lock()
	defer t.Unlock()

	root := t.root
	foundNode, parent := t.search(nil, root, node)
	if foundNode != nil { //We found our node, now DELETE IT!!

		//Apply the delete algorithm
		numChildren := 0
		if foundNode.left != nil {
			numChildren++
		}
		if foundNode.right != nil {
			numChildren++
		}

		switch numChildren {
		case 0:
			foundNode = nil
			t.size--
			return nil
		case 1:
			if parent.left == foundNode {
				t.oneSubTreeDelete(parent, foundNode, leftDir)
			} else {
				t.oneSubTreeDelete(parent, foundNode, rightDir)
			}
			t.size--
			return nil
		case 2:
			t.twoSubTreeDelete(foundNode)
			t.size--
			return nil
		}

	}

	return errors.New("Node not found")
}

func (t *Tree) oneSubTreeDelete(parent, node *Node, dir childDir) {
	var newChild *Node
	if node.left != nil {
		newChild = node.left
	} else {
		newChild = node.right
	}

	switch dir {
	case leftDir:
		parent.left = newChild
		break
	case rightDir:
		parent.right = newChild
		break
	}
}

func (t *Tree) twoSubTreeDelete(node *Node) {
	//Get the min of the right subtree
	rightTree := NewTree(node.right)
	minNode := rightTree.Min()

	// replace the current node with the min node of the right subtree
	node.key = minNode.key
	node.value = minNode.value

	// delete the duplicate node (aka the minNode from above)
	rightTree.Delete(minNode)
}

// Recursivly iterate through the tree with a given key, apply action when node or position is found
func (t *Tree) search(parent, root, node *Node) (*Node, *Node) {

	// Return the root if its the root we're looking for, so the action can be applied to it
	if root == nil || root.key == node.key {
		return root, parent
	}

	if root.key > node.key {
		parent = root
		root = root.left
		return t.search(parent, root, node)
	} else if root.key < node.key {
		parent = root
		root = root.right
		return t.search(parent, root, node)
	}

	return nil, nil
}

// Traverse the tree and apply the Visit function on each Node
func (t *Tree) Traverse(traverse Traversal, visit VisitFunc) {
	t.Lock()
	defer t.Unlock()

	root := t.root
	switch traverse {
	case PostOrder:
		t.doPostOrder(root, visit)
	case PreOrder:
		t.doPreOrder(root, visit)
	case InOrder:
		t.doInOrder(root, visit)
	}
}

// PostOrder: Left, Right, Root
func (t *Tree) doPostOrder(root *Node, visit VisitFunc) {
	if root == nil {
		return
	}

	t.doPostOrder(root.left, visit)
	t.doPostOrder(root.right, visit)
	visit(root)
}

// PreOrder: Root, Left, Right
func (t *Tree) doPreOrder(root *Node, visit VisitFunc) {
	if root == nil {
		return
	}

	visit(root)
	t.doPreOrder(root.left, visit)
	t.doPreOrder(root.right, visit)
}

// InOrder: Left, Root, right
func (t *Tree) doInOrder(root *Node, visit VisitFunc) {
	if root == nil {
		return
	}

	t.doInOrder(root.left, visit)
	visit(root)
	t.doInOrder(root.right, visit)
}

func (t *Tree) Min() *Node {
	t.Lock()
	defer t.Unlock()

	root := t.root
	for root.left != nil {
		root = root.left
	}

	return root
}

func (t *Tree) Max() *Node {
	t.Lock()
	defer t.Unlock()

	root := t.root
	for root.right != nil {
		root = root.right
	}

	return root
}
