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
