package bst

import (
	"sync"
)

/*
A Node represents a node in the tree
*/
type Node struct {
	key   uint64
	value interface{}
	left  *Node
	right *Node
	sync.RWMutex
}

// Create a new node
func NewNode(key uint64, value interface{}) *Node {
	node := &Node{key: key, value: value}
	return node
}

// Get the node's key
func (n *Node) Key() uint64 {
	return n.key
}

// Get the node's value
func (n *Node) Value() interface{} {
	return n.value
}

// Get the node's left node
func (n *Node) Left() *Node {
	return n.left
}

// Get the node's rigth node
func (n *Node) Right() *Node {
	return n.right
}
