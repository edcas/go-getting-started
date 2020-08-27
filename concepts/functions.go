package functions

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
	port := 3000
	startWebServer(port, 2)
	fmt.Println("**************************")
	_, err := startWebServerError(port, 2)
	fmt.Println(err)
	fmt.Println("**************************")
	pt, err := startWebServerError(port, 2)
	fmt.Println(pt, err)
}

func startWebServer(port int, numberOfRetries int) {
	fmt.Println("Starting server...")
	fmt.Println("Server started on port", port)
	fmt.Println("Number of retries", numberOfRetries)
}

func startWebServerError(port int, numberOfRetries int) (int, error) {
	fmt.Println("Starting server...")
	fmt.Println("Server started on port", port)
	fmt.Println("Number of retries", numberOfRetries)
	return port, errors.New("Something went wrong")
}
