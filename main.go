package main

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

func init()  {
	log.SetFormatter(&log.JSONFormatter{})
}
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Info("Send info")
		log.Error("Send Error")
		w.Write([]byte("send logs to stdout and stderr"))
	})

	log.Fatal(http.ListenAndServe(":5002", nil))
}
