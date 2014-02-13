package dist

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

type testKey int

func (k testKey) Value() int {
	return int(k)
}

func init() {
	rand.Seed(int64(time.Now().Nanosecond()))
}

func createSingleTree() *Tree {
	return NewDistTree(1, func(a, b interface{}) bool {
		return a.(testKey) < b.(testKey)
	})
}

func createMultiTree(num int) *Tree {
	return NewDistTree(num, func(a, b interface{}) bool {
		return a.(testKey) < b.(testKey)
	})
}

func Benchmark_SingleTree100Insert(b *testing.B) {

	for i := 0; i < b.N; i++ {
		// Setup
		b.StopTimer()
		t := createSingleTree()
		wg := new(sync.WaitGroup)
		b.StartTimer()
		for j := 0; j < 100; j++ {
			go func() {
				key := rand.Int()
				item := key
				t.Insert(testKey(key), item)
				wg.Done()
			}()
			wg.Wait()
		}
	}
}

func Benchmark_SingleTree100Delete(b *testing.B) {

	// Delete 100 keys
	for i := 0; i < b.N; i++ {
		// Setup
		b.StopTimer()
		t := createSingleTree()
		var keys []int

		wg := new(sync.WaitGroup)

		// Insert 100 keys and record them
		for j := 0; j < 100; j++ {

			key := rand.Int()
			item := key
			t.Insert(testKey(key), item)
			keys = append(keys, key)
		}
		b.StartTimer() // restart timer
		wg.Add(100)
		for _, key := range keys {
			go func() {
				t.Delete(testKey(key))
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func Benchmark_SingleTree100Get(b *testing.B) {

	// Setup
	b.StopTimer()
	t := createSingleTree()
	var keys []int

	wg := new(sync.WaitGroup)

	// Insert 100 keys and record them
	for j := 0; j < 100; j++ {

		key := rand.Int()
		item := key
		t.Insert(testKey(key), item)
		keys = append(keys, key)
	}
	b.StartTimer() // restart timer

	// Delete 100 keys
	for i := 0; i < b.N; i++ {
		wg.Add(100)
		for _, key := range keys {
			go func() {
				t.Get(testKey(key))
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func Benchmark_SingleTree100Exists(b *testing.B) {
	// Setup
	b.StopTimer()
	t := createSingleTree()
	var keys []int

	wg := new(sync.WaitGroup)

	// Insert 100 keys and record them
	for j := 0; j < 100; j++ {

		key := rand.Int()
		item := key
		t.Insert(testKey(key), item)
		keys = append(keys, key)
	}
	b.StartTimer() // restart timer

	// Delete 100 keys
	for i := 0; i < b.N; i++ {
		wg.Add(100)
		for _, key := range keys {
			go func() {
				t.Exists(testKey(key))
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func Benchmark_5MultiTree100Inserts(b *testing.B) {
	// Setup
	b.StopTimer()
	t := createMultiTree(5) // Creata a distributed tree backed by 10 trees
	wg := new(sync.WaitGroup)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		for j := 0; j < 100; j++ {
			go func() {
				key := rand.Int()
				item := key
				t.Insert(testKey(key), item)
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func Benchmark_10MultiTree100Inserts(b *testing.B) {
	// Setup
	b.StopTimer()
	t := createMultiTree(10) // Creata a distributed tree backed by 10 trees
	wg := new(sync.WaitGroup)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		for j := 0; j < 100; j++ {
			go func() {
				key := rand.Int()
				item := key
				t.Insert(testKey(key), item)
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func Benchmark_20MultiTree100Inserts(b *testing.B) {
	// Setup
	b.StopTimer()
	t := createMultiTree(10) // Creata a distributed tree backed by 10 trees
	wg := new(sync.WaitGroup)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		for j := 0; j < 100; j++ {
			go func() {
				key := rand.Int()
				item := key
				t.Insert(testKey(key), item)
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func Benchmark_25MultiTree100Inserts(b *testing.B) {
	// Setup
	b.StopTimer()
	t := createMultiTree(10) // Creata a distributed tree backed by 10 trees
	wg := new(sync.WaitGroup)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		for j := 0; j < 100; j++ {
			go func() {
				key := rand.Int()
				item := key
				t.Insert(testKey(key), item)
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func Benchmark_5MultiTree100Delete(b *testing.B) {
	// Setup
	b.StopTimer()
	t := createMultiTree(5)
	var keys []int
	wg := new(sync.WaitGroup)

	// Insert 100 keys and record them
	for j := 0; j < 100; j++ {
		key := rand.Int()
		item := key
		t.Insert(testKey(key), item)
		keys = append(keys, key)
	}
	b.StartTimer() // restart timer

	// Delete 100 keys
	for i := 0; i < b.N; i++ {
		for _, key := range keys {
			go func() {
				t.Delete(testKey(key))
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func Benchmark_10MultiTree100Delete(b *testing.B) {
	// Setup
	b.StopTimer()
	t := createMultiTree(10)
	var keys []int
	wg := new(sync.WaitGroup)

	// Insert 100 keys and record them
	for j := 0; j < 100; j++ {
		key := rand.Int()
		item := key
		t.Insert(testKey(key), item)
		keys = append(keys, key)
	}
	b.StartTimer() // restart timer

	// Delete 100 keys
	for i := 0; i < b.N; i++ {
		for _, key := range keys {
			go func() {
				t.Delete(testKey(key))
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func Benchmark_20MultiTree100Delete(b *testing.B) {
	// Setup
	b.StopTimer()
	t := createMultiTree(20)
	var keys []int
	wg := new(sync.WaitGroup)

	// Insert 100 keys and record them
	for j := 0; j < 100; j++ {
		key := rand.Int()
		item := key
		t.Insert(testKey(key), item)
		keys = append(keys, key)
	}
	b.StartTimer() // restart timer

	// Delete 100 keys
	for i := 0; i < b.N; i++ {
		for _, key := range keys {
			go func() {
				t.Delete(testKey(key))
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func Benchmark_25MultiTree100Delete(b *testing.B) {
	// Setup
	b.StopTimer()
	t := createMultiTree(25)
	var keys []int
	wg := new(sync.WaitGroup)

	// Insert 100 keys and record them
	for j := 0; j < 100; j++ {
		key := rand.Int()
		item := key
		t.Insert(testKey(key), item)
		keys = append(keys, key)
	}
	b.StartTimer() // restart timer

	// Delete 100 keys
	for i := 0; i < b.N; i++ {
		for _, key := range keys {
			go func() {
				t.Delete(testKey(key))
				wg.Done()
			}()
		}
		wg.Wait()
	}
}
