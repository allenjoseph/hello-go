package main

import (
	"fmt"
	"net/http"
)

type reqHandler struct {
	greeting string
}

func (handler reqHandler) ServeHTTP(resWriter http.ResponseWriter, req *http.Request) {
	resWriter.Write([]byte(fmt.Sprintf("%v world", handler.greeting)))
}

func main() {
	http.Handle("/", &reqHandler{greeting: "hello"})
	http.ListenAndServe(":8000", nil)
}
