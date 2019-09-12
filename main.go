package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("I am a superman: %s", "John")
		w.Write([]byte("send logs to stdout and stderr"))
	})

	log.Fatal(http.ListenAndServe(":5002", nil))
}
