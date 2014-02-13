package dist

import (
	"github.com/stathat/treap"
	"sync"
)

// The backend identfier key
type Key interface {
	// Return an int represent the keys value
	Value() int
}

// Item represents values in the tree
type Item interface{}

// A function designed to compare keys
//type LessFunc func(a, b interface{}) bool

// The frontend of the entire distributed collection of trees
// that allows the backend trees to be used as a single logical tree
type Tree struct {
	// Number of backend trees
	numBackends int

	// Backends
	backends []treeBackend
}

// A bst backend
type treeBackend struct {
	// Tree
	tree *treap.Tree

	// Mutex
	sync.RWMutex
}

// Creates a new Distributed tree frontend, with num many backends
func NewDistTree(num int, lessfn treap.LessFunc) *Tree {
	tree := &Tree{
		numBackends: num,
	}

	// Create the initial collection of tree backends
	for i := 0; i < num; i++ {
		t := treap.NewTree(lessfn)
		backend := treeBackend{tree: t}
		tree.backends = append(tree.backends, backend)
	}

	return tree
}

// Insert a Item into the Tree via a Key
func (t *Tree) Insert(key Key, item Item) {
	backend := t.getBackend(key)

	// Lock the backend
	backend.Lock()
	defer backend.Unlock()

	backend.tree.Insert(key, item)
}

// Delete an item from the Tree
func (t *Tree) Delete(key Key) {
	backend := t.getBackend(key)

	// Lock the backend
	backend.Lock()
	defer backend.Unlock()

	backend.tree.Delete(key)
}

// Check to see if an item exists
func (t *Tree) Exists(key Key) bool {
	backend := t.getBackend(key)

	// Lock the backend
	backend.RLock()
	defer backend.RUnlock()

	return backend.tree.Exists(key)

}

// Get an item from the tree
func (t *Tree) Get(key Key) Item {
	backend := t.getBackend(key)

	// Lock the backend
	backend.RLock()
	defer backend.RUnlock()

	return backend.tree.Get(key)
}

// Get the appropriate backend for the current key
func (t *Tree) getBackend(key Key) treeBackend {
	id := key.Value() % t.numBackends
	return t.backends[id]
}
