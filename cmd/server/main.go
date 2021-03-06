package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type RawHello struct {
	Greeting string
}

func (handler RawHello) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	defer log.Printf(req.URL.Path)
	res.Write([]byte(fmt.Sprintf("%v world", handler.Greeting)))
}

type JsonHello struct {
	Greeting string `json:"hello"`
}

func (handler JsonHello) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	defer log.Printf(req.URL.Path)
	output, err := json.Marshal(handler)
	if err != nil {
		log.Println(err)
	} else {
		res.Header().Set("Content-Type", "application/json")
		res.Write(output)
	}
}

func main() {
	http.Handle("/", &RawHello{Greeting: "hello"})
	http.Handle("/json", &JsonHello{Greeting: "world"})
	http.HandleFunc("/inline", func(res http.ResponseWriter, req *http.Request) {
		defer log.Printf(req.URL.Path)
		res.Write([]byte("Hello World! (inline handle function)"))
	})
	http.ListenAndServe(":8000", nil)
}
