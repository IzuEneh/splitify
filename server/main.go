package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Starting server...")
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	filePath := os.Getenv("STATIC_FILE_PATH")
	http.Handle("/", http.FileServer(http.Dir(filePath)))
	http.ListenAndServe("0.0.0.0:8080", nil)
	fmt.Println("Running...")
}
