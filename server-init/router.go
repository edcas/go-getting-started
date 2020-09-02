package main

import (
	"net/http"
)

type Router struct {
	rules map[string]http.HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		rules: make(map[string]http.HandlerFunc),
	}
}

func (router *Router) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Hello!"))
}