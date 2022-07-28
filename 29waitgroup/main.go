package main

import (
	"fmt"
	"sync"
)

var x = 0

func add(wg *sync.WaitGroup, m *sync.Mutex) {

	m.Lock()
	x = x + 1
	m.Unlock()
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var m sync.Mutex

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		//go add(&wg)
		go add(&wg, &m)
	}

	wg.Wait()

	fmt.Println(x)
}

package main

import (
	"fmt"
	"sync"
)

var x = 0

func increment(wg *sync.WaitGroup) {
	x = x + 1
	wg.Done()
}
func main() {
	var w sync.WaitGroup
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w)
	}
	w.Wait()
	fmt.Println("Main X:", x)
}
