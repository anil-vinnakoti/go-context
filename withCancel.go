package main

import (
	"fmt"
	"time"
	"context"
)

func main(){
	ctx, cancel := context.WithCancel(context.Background())
	
	go worker(ctx)

	time.Sleep(time.Second * 3)
	cancel()

	time.Sleep(time.Second)

}

func worker(ctx context.Context){

	for{
		select{
		case <- ctx.Done():
			fmt.Println("worker stopped", ctx.Err())
			return
		default:
			fmt.Println("working...")
			time.Sleep(time.Second)
		}
	}

}