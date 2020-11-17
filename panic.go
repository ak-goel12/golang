package main

import (
	"fmt"
	"sync"
	"time"
)

func cleanUp(wg *sync.WaitGroup) {
	defer wg.Done() //make sure that the program run without error

	if r := recover(); r != nil {
		fmt.Println("Recover in cleanup:", r)
	}
}

//wait group is passed as a pointer to the function
func worker(id int, wg *sync.WaitGroup) {

	defer cleanUp(wg)
	for i := 0; i < 3; i++ {

		fmt.Printf("Worker %d starting\n", id)
		time.Sleep(time.Second)
		fmt.Printf("Worker %d done\n", id)
		if i == 2 {
			panic("Oh dear, a 2")
		}
	}
}

func main() {

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
}
