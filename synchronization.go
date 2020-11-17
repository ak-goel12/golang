package main

import (
	"fmt"
	"sync"
	"time"
)

//wait group is passed as a pointer to the function
func worker(id int, wg *sync.WaitGroup) {

	defer wg.Done() //make sure that the program run without error

	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
}
