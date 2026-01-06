package main

import (
	"fmt"
	"time"
	"context"
)

func main(){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 2)
	defer cancel()

	select{
	case <-time.After(time.Second * 3):
		fmt.Println("work done")
	case <-ctx.Done():
		fmt.Println("ended", ctx.Err())

	}
}