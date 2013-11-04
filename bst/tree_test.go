package bst

import (
	//"fmt"
	"math/rand"
	"runtime"
	"sync"
	"testing"
)

var (
	testRoot *Node = NewTestNode()
)

func newTestTree() *Tree {
	return NewTree(testRoot)
}

func newRandomTree() *Tree {
	return NewTree(NewNode(uint64(rand.Uint32()), rand.Int()))
}

func Test_TreeSize(t *testing.T) {
	tree := newTestTree()
	if tree.size != 1 && tree.Size() != 1 {
		t.Fatalf("Expected tree size: %v, actual tree size: %v\n", 1, tree.size)
	}
}

func Test_TreeRootNode(t *testing.T) {
	tree := newTestTree()
	if tree.root != testRoot {
		t.Fatalf("Excpeted tree root %v, actual tree root: %v\n", testRoot, tree.root)
	}
}

func Test_TreeInsert(t *testing.T) {
	tree := newRandomTree()
	insertRandomNodes(tree, 16)
	if tree.Size() != 17 {
		t.Fatalf("Expected tree size: %v, actual tree size: %v\n", 17, tree.Size())
	}

}

func Test_TreeConcurrentInsert(t *testing.T) {
	tree := newRandomTree()
	numGoroutines := runtime.NumCPU()
	numInserts := 1000 // per goroutine
	prev := runtime.GOMAXPROCS(numGoroutines)

	var wg sync.WaitGroup

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			insertRandomNodes(tree, numInserts)
			wg.Done()
		}()
	}

	wg.Wait()

	if tree.Size() != numGoroutines*numInserts+1 {
		t.Fatalf("Expected size of tree: %v, actual size of tree: %v\n", numGoroutines*numInserts+1, tree.Size())
	}

	runtime.GOMAXPROCS(prev)
}

func Test_TreeMin(t *testing.T) {
	tree := NewTree(NewNode(10, ""))
	tree.Insert(NewNode(11, ""))
	tree.Insert(NewNode(9, ""))
	tree.Insert(NewNode(1, ""))
	if tree.Min().Key() != 1 {
		t.Fatalf("Expected min node: %v, actual min node: %v\n", 1, tree.Min().Key())
	}
}

func Test_TreeMax(t *testing.T) {
	tree := NewTree(NewNode(10, ""))
	tree.Insert(NewNode(11, ""))
	tree.Insert(NewNode(9, ""))
	tree.Insert(NewNode(1, ""))
	if tree.Max().Key() != 11 {
		t.Fatalf("Expected min node: %v, actual min node: %v\n", 11, tree.Min().Key())
	}
}

func Test_TreeSingleFind(t *testing.T) {
	tree := newTestTree()
	ok, val := tree.Find(1)
	if !ok || val != "HelloWorld" {
		t.Fatalf("Expected val: %v, actual value: %v\n", testValue, val)
	}
}

func Test_TreeMultiFind(t *testing.T) {
	tree := NewTree(NewNode(10, ""))
	tree.Insert(NewNode(11, "Eleven"))
	tree.Insert(NewNode(9, "Nine"))
	tree.Insert(NewNode(1, "One"))

	ok, val := tree.Find(9)
	if !ok || val != "Nine" {
		t.Fatalf("Expected val: %v, actual value: %v\n", "Nine", val)
	}
}

func insertRandomNodes(tree *Tree, num int) {
	var i int
	for i = 0; i < num; i++ {
		j := rand.Uint32()
		node := NewNode(uint64(j), j)
		tree.Insert(node)
	}
}
