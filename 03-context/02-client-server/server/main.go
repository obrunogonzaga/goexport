package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("server: handler started")
	defer log.Println("server: handler ended")
	select {
	case <-time.After(5 * time.Second):
		log.Println("server: Request processed")
		w.Write([]byte("Request processed\n"))
	case <-ctx.Done():
		log.Println("server: ", ctx.Err())
	}
}
