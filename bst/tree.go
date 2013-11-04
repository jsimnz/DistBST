package bst

import (
	"errors"
	"fmt"
	"sync"
)

const (
	_ Traversal = iota
	InOrder
	PreOrder
	PostOrder

	_ childDir = iota
	leftDir
	rightDir

	_ actionType = iota
	insertAction
	findAction
)

type Traversal int
type childDir int
type actionType int

type VisitFunc func(node *Node)
type actionFunc func(node *Node)

var (
	emptyActionFunc actionFunc = func(node *Node) {}
)

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

func (t *Tree) String() string {
	t.Lock()
	defer t.Unlock()
	return fmt.Sprintf("{root %v, size: %v}", t.root, t.size)
}

func (t *Tree) Size() int {
	return t.size
}

// Insert a Node into the Tree
func (t *Tree) Insert(node *Node) error {
	t.Lock()
	defer t.Unlock()

	if t.search(t.root, node, insertAction, emptyActionFunc) != nil {
		// there is already a node with that key, see searcht.
		return errors.New("Duplicate node")
	}

	t.size++
	return nil
}

func (t *Tree) search(root, node *Node, act actionType, actionFn actionFunc) error {
	if root.Key() == node.Key() {
		if act == findAction {
			actionFn(root)
			return nil
		} else {
			return errors.New("Node already exists")
		}
	} else if root.Key() < node.Key() {
		if root.right != nil {
			return t.search(root.right, node, act, actionFn)
		}
		if act == findAction {
			return errors.New("Node not found")
		}
		root.right = node
		return nil
	} else if root.Key() > node.Key() {
		if root.left != nil {
			return t.search(root.left, node, act, actionFn)
		}
		if act == findAction {
			return errors.New("")
		}
		root.left = node
	}

	return nil
}

// Find a Node in the tree, if it isnt found return NotExistError
func (t *Tree) Find(key uint64) (bool, interface{}) {
	t.Lock()
	defer t.Unlock()

	node := NewNode(key, nil)

	var results interface{}
	err := t.search(t.root, node, findAction, func(node *Node) {
		results = node.Value()
	})

	if err != nil {
		return false, nil
	}

	return true, results //Node not found
}

// Delete a Node in a tree and rebalance
func (t *Tree) Delete(node *Node) error {
	t.Lock()
	defer t.Unlock()

	
	

	return errors.New("Node not found")
}

/*func (t *Tree) oneSubTreeDelete(parent, node *Node, dir childDir) {
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
}*/

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
