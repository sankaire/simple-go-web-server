package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var counter int
var mutex = &sync.Mutex{}

func echoString(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Kenya")
}

func incrementCounter(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	counter++
	fmt.Fprintf(w, strconv.Itoa(counter))
	mutex.Unlock()
}

func main() {
	http.HandleFunc("/", echoString)
	http.HandleFunc("/increment", incrementCounter)
	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hi")
	})
	log.Fatal(http.ListenAndServe(":8001", nil))
}
