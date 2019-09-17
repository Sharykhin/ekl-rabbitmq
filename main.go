package main

import (
	log "github.com/sirupsen/logrus"
	"fmt"
	"net/http"
)

type Student struct {
	Action string `json:"action"`
	UserId string    `json:"user_id"`
	Message string     `json:"message"`
}

func init()  {
	log.SetFormatter(&log.JSONFormatter{})
}
func main() {

	//logger,err :=fluent.New(fluent.Config{FluentHost:"fluentd", FluentPort:24224})
	//fmt.Println(logger.BufferLimit)
	//if err != nil{
	//	fmt.Println("error creating fluentd instance", err)
	//	return
	//}
	//
	//defer logger.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//log.Info("test")
		fmt.Println(Student{"UserCreatd", "22233", "Ha Ha Ha"})
		//_, err := os.Stdout.Write([]byte(`{"app": "simplebank", "key":"value"}`))
		//if err !=nil {
		//	panic(err)
		//}
		//tag := "fluentd-log-demo"
		//data := map[string]interface{}{
		//	"name": "John",
		//	"id_no": 1234,
		//	"age": 25,
		//}
		//err := logger.Post(tag,data)
		//
		//if err != nil {
		//	fmt.Printf("Got error: %v", err)
		//}

		_, _  = w.Write([]byte("OK"))
		return
	})

	log.Fatal(http.ListenAndServe(":5002", nil))
}
