package main

import (
	"fmt"
	"net/http"
)


func main(){
	fmt.Println("Started")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}