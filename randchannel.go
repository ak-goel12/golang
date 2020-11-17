package main

import (
	"math/rand"
	"sync"
	"time"
)

func randint(wg *sync.WaitGroup, c chan int) {
	defer wg.Done()
	//random generator
	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= 10; i++ {
		r := rand.Intn(100)
		c <- r
	}

}
func main() {
	var wg sync.WaitGroup
	//create a buffer of 10
	c := make(chan int, 10) //channel
	//wait group for synchronization
	wg.Add(1)
	//go routine
	go randint(&wg, c)
	wg.Wait()
	close(c) //close the channel
	//print the data in channel
	for item := range c {
		println(item)
	}
}
