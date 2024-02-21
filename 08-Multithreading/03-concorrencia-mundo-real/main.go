package main

import (
	"fmt"
	"net/http"
	"sync"
)

var number uint64 = 0

func main() {
	m := sync.Mutex{}
	http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m.Lock()
		number++
		m.Unlock()
		w.Write([]byte("Voce é o visitante de número " + fmt.Sprint(number) + " da página!"))
	}))
}
