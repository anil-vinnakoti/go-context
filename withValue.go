package main

import (
	"fmt"
	"context"
)

type userIDKeyType string

func main(){

	var userIDKey userIDKeyType = "userID"
	ctx := context.WithValue(context.Background(), userIDKey, "28031998")
	
	printUser(ctx)
}

func printUser(ctx context.Context){
	fmt.Println("User ID:", ctx.Value(userIDKeyType("userID")))
}