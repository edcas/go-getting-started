package main

import (
	"fmt"
	"net/http"
)

func CheckAuth() Middleware {
	return func(handlerFunc http.HandlerFunc) http.HandlerFunc {
		return func(writer http.ResponseWriter, request *http.Request) {
			flag := false
			fmt.Println("Checking Authentication")
			if flag {
				handlerFunc(writer, request)
			} else {
				return
			}
		}
	}
}
