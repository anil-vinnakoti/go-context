package main

import (
	"fmt"
	"net/http"
	"time"
)

func main(){
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r * http.Request){
	ctx := r.Context()

	select{
	case <-time.After(time.Second * 5):
		fmt.Fprintln(w, "Response completed")
		fmt.Println("Response completed")

	case <-ctx.Done():
		fmt.Println("Request cancelled by client")
	}
}