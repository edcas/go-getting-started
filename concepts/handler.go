package main

import (
	"net/http"
)

type fooHandler struct {
	Message string
}

func (handler *fooHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte(handler.Message))
}

func barHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("bar called"))
}

func main() {
	http.Handle("/foo", &fooHandler{Message: "foo called"})
	http.HandleFunc("/bar", barHandler)
	http.ListenAndServe(":5000", nil)
}
