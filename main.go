package main

import (
	"fmt"
	"github.com/fluent/fluent-logger-golang/fluent"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Student struct {
	Name string `msg:"name"`
	IdNo int    `msg:"id_no"`
	Age int     `msg:"age"`
}

func init()  {
	//log.SetFormatter(&log.JSONFormatter{})
}
func main() {

	logger,err :=fluent.New(fluent.Config{FluentHost:"fluentd", FluentPort:24224})
	fmt.Println(logger.BufferLimit)
	if err != nil{
		fmt.Println("error creating fluentd instance", err)
		return
	}

	defer logger.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Info("test")

		tag := "fluentd-log-demo"
		data := map[string]interface{}{
			"name": "John",
			"id_no": 1234,
			"age": 25,
		}
		err := logger.Post(tag,data)

		if err != nil {
			fmt.Printf("Got error: %v", err)
		}

		_, _  = w.Write([]byte("OK"))
	})

	log.Fatal(http.ListenAndServe(":5002", nil))
}
