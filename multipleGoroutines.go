package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	count :=  3
	for i:=1;i<= count; i++{
	wg.Add(1)
	
	go func(id int) {
		defer wg.Done()
		worker(ctx, id)
		}(i)
	}

	// cancel immediately (no sleep needed)
	time.Sleep(time.Second * 3)
	cancel()

	wg.Wait()
	fmt.Println("main exiting")
}

func worker(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("worker %d stopped\n", id)
			return
		default:
			fmt.Printf("worker %d working...\n", id)
			time.Sleep(time.Second) // simulate work
		}
	}
}