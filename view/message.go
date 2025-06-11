package view

import "fmt"

func ShowStartMessage(port int, origin string) {
	fmt.Printf("Starting caching proxy server on http://localhost:%d \n", port)
	fmt.Printf("Forwarding requests to %s \n", origin)
}

func ShowClearMessage() {
	fmt.Println("Cache cleared successfully.")
}
