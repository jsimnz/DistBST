package bst

import (
	"testing"
)

var (
	testKey   uint64      = 1
	testValue interface{} = "HelloWorld"

	testRKey   uint64      = 2
	testRValue interface{} = "HelloWorldRight"

	testLKey   uint64      = 3
	testLValue interface{} = "HelloWorldLeft"
)

func NewTestNode() *Node {
	return NewNode(testKey, testValue)
}

func NewLeftTestNode() *Node {
	return NewNode(testLKey, testLValue)
}

func NewRightTestNode() *Node {
	return NewNode(testRKey, testRValue)
}

func Test_NodeCreateTest(t *testing.T) {
	node := NewTestNode()
	if *node.key != 1 {
		t.Fail()
	} else if node.value != "HelloWorld" {
		t.Fail()
	} else if node.left != nil {
		t.Fail()
	} else if node.right != nil {
		t.Fail()
	}
}

func Test_NodeKeyTest(t *testing.T) {
	node := NewTestNode()
	if node.Key() != testKey {
		t.Fail()
	}
}

func Test_NodeValueTest(t *testing.T) {
	node := NewTestNode()
	if node.Value() != testValue {
		t.Fail()
	}
}

func Test_LeftNodeTest(t *testing.T) {
	node := NewTestNode()
	lnode := NewLeftTestNode()
	rnode := NewRightTestNode()

	node.left = lnode
	node.right = rnode

	if node.Right() != rnode {
		t.Fail()
	} else if node.Left() != lnode {
		t.Fail()
	}
}
