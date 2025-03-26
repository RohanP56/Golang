# Sync Waitgroup
- **Using go routine we can execute multiple task concurrently. But some functions are execute early and some functions takes times. so we are not waiting to execute all functions.**
- To handle this problem there's a term called `sync.WaitGroup` it's a synchronization premitive in Go, that is used to wait for a collection of goroutines to finish their execution.
- It allows you to coordinate the execution of multiple goroutines and ensure that they all complete before continuing with the rest of the program.

``` Go
package main

import (
	"fmt"
	"sync"
)

func worker(i int, wg *sync.WaitGroup) {
	defer wg.Done() // after every function call, go will mark that iteration is "Done"
	fmt.Printf("Worker %d has started Working\n", i)
	fmt.Printf("Worker %d has ended is Work\n", i)

}

func main() {
	var wg sync.WaitGroup // Task list

	for i := range 3 {
		wg.Add(1) // Increment weight group counter (adding check list)
		go worker(i, &wg)
	}
	wg.Wait() // go routine will wait to execute all function
	fmt.Println("Workers work completed")
}
```