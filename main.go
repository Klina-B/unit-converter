package main

import (
	"fmt"
	"net/http"

	"github.com/Klina-B/unit-converter/handlers"
)


func main(){
	fmt.Println("Started")
	http.HandleFunc("/", handlers.LengthHandler)
	http.ListenAndServe(":8080", nil)
}