package main

import "fmt"

func main()  {
	server := NewServer(":5000")
	server.Handle("/", HandlerRoot)
	server.Handle("/home", HandlerHome)
	err := server.Listen()

	if err != nil {
		fmt.Println(err)
	}
}