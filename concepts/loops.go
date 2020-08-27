package main

import (
	"fmt"
)

func main() {
	
	var i int
	for i < 5 {
		fmt.Println(i)
		i++
		if i == 3 {
			continue
		}
		fmt.Println("Continuing...")
	}
	fmt.Println("*******************")
	
	for i := 0 ;i < 5; i++ {
		fmt.Println(i)
	}
	
	fmt.Println("*******************")
	
	slice := []int{40, 41, 42}
	for i, v := range slice {
		fmt.Println(i, v)
	}
	
	fmt.Println("values")
	for _, v := range slice {
		fmt.Println(v)
	}

	fmt.Println("keys")	
	for i, _ := range slice {
		fmt.Println(i)
	}
	
	fmt.Println("*******************")
	
	wellKnowPorts := map[string]int{"http": 80, "https": 443}
	for keys, value := range wellKnowPorts {
		fmt.Println(keys, value)
	}
	
	fmt.Println("**** Panic ***")
	panic("Something bad just happened")
	
}
