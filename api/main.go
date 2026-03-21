package main

import (
	"api-development/router"
	"log"
	"net/http"
)

func main() {
	r := router.New()
	log.Println("started server at port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Failed to start the server")
	}
}
