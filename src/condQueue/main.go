package main

import (
	"sync"
	"time"
)



func main() {
	c := sync.NewCond(&sync.Mutex{}}
	queue := make([]interface{}, 0, 10)
	// c.Wait is a blocking call 
	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay)
		c.L.Lock()
		queue = queue[1:]
		fmt.Println("Removed from the queue")
		c.L.Unlock()
		c.Signal()
	}

	for i := 0; i < 10; i++ {
		c.L.Lock()
		for len(queue) == 2 {
			c.Wait()
		}
		fmt.Println("Addong to the queue")
		queue = append(queue, struct{}{})
		go removeFromQueue(1 * time.Second)
		c.L.Lock()
	}

}
