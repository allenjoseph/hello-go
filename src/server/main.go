package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type rawHello struct {
	greeting string
}

func (handler rawHello) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte(fmt.Sprintf("%v world", handler.greeting)))
}

type jsonHello struct {
	Greeting string `json:"hello"`
}

func (handler jsonHello) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	output, err := json.Marshal(handler)
	if err != nil {
		fmt.Println(err)
	} else {
		res.Header().Set("Content-Type", "application/json")
		res.Write(output)
	}
}

func main() {
	http.Handle("/", &rawHello{greeting: "hello"})
	http.Handle("/json", &jsonHello{Greeting: "world"})
	http.ListenAndServe(":8000", nil)
}
