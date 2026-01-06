package main

import (
	"fmt"
	"time"
	"context"
)

func main(){
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*2))
	defer cancel()
	go worker(ctx)

	time.Sleep(time.Second *4)
	
}


func worker(ctx context.Context){
	for{
		select{
		case <-ctx.Done():
			fmt.Println("ended...", ctx.Err())
			return
		default:
			fmt.Println("working...")
			time.Sleep(time.Second)
		}
	}

}