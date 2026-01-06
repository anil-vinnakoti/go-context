package main

import (
	"fmt"
	"time"
	"context"
)

func main() {
	parentCtx, parentCancel := context.WithCancel(context.Background())
	defer parentCancel()

	childCtx, childCancel := context.WithCancel(parentCtx)
	defer childCancel()

	go func() {
		for {
			select {
			case <-childCtx.Done():
				fmt.Println("child stopped:", childCtx.Err())
				return
			default:
				fmt.Println("working on something...")
			}
		}
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("canceling parent")
	parentCancel()

	time.Sleep(time.Second)
}
