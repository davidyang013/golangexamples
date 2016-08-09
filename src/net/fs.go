package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	file, _ := os.Getwd()
	http.Handle("/", http.FileServer(http.Dir("./")))
	log.Printf("Directory [%s] can be accessed via http://127.0.0.1:8081", file)
	http.ListenAndServe(":8081", nil)
}
