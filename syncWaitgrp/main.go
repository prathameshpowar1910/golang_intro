package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup) {
	//do some work
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)
	//some task is happening here
	fmt.Printf("Worker %d done\n", id)

}

func main() {

	var wg sync.WaitGroup
	//start 3 worker goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	//wait for all workers to complete
	wg.Wait()
	fmt.Println("All workers complete")
}
