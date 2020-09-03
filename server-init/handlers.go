package main

import (
	"fmt"
	"net/http"
)

func HandlerRoot(writer http.ResponseWriter, request *http.Request)  {
	writer.Write([]byte("Hello"))
}

func HandlerHome(writer http.ResponseWriter, request *http.Request)  {
	_, err := writer.Write([]byte("Home"))

	if err != nil {
		fmt.Println(err)
	}
}
