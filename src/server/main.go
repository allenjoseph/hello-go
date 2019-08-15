package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func logger(url string) {
	fmt.Println(url)
}

type RawHello struct {
	Greeting string
}

func (handler RawHello) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	defer logger(req.URL.Path)
	res.Write([]byte(fmt.Sprintf("%v world", handler.Greeting)))
}

type JsonHello struct {
	Greeting string `json:"hello"`
}

func (handler JsonHello) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	defer logger(req.URL.Path)
	output, err := json.Marshal(handler)
	if err != nil {
		fmt.Println(err)
	} else {
		res.Header().Set("Content-Type", "application/json")
		res.Write(output)
	}
}

func main() {
	http.Handle("/", &RawHello{Greeting: "hello"})
	http.Handle("/json", &JsonHello{Greeting: "world"})
	http.ListenAndServe(":8000", nil)
}
