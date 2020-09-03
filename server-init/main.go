package main

import "fmt"

func main()  {
	server := NewServer(":5000")
	err := server.Listen()

	if err != nil {
		fmt.Println(err)
	}
}