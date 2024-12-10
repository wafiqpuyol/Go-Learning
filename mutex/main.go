package main

import (
	"fmt"
	"sync"
)

type like struct {
	postName string
	count    int
	mu       sync.Mutex
}

func incrementLike(l *like, wg *sync.WaitGroup) {
	l.mu.Lock()
	wg.Add(1)
	defer func() {
		l.mu.Unlock()
		fmt.Println(l.count)
		wg.Done()
	}()
	l.count++
}

func main() {
	var wg sync.WaitGroup
	likes := like{
		postName: "Go",
		count:    0,
	}

	for i := 0; i < 5; i++ {
		go incrementLike(&likes, &wg)
	}
	fmt.Println(likes.count)
	wg.Wait()
}
