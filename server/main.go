package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting server...")
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.ListenAndServe("0.0.0.0:8080", nil)
}
