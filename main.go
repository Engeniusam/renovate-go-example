package main

import (
	"fmt"

	"github.com/gorilla/websocket"
)

func main() {
	fmt.Println("Hello, Renovate!")

	// Creating a basic websocket dialer to use the imported package
	var dialer websocket.Dialer
	fmt.Printf("Dialer: %+v\n", dialer)
}
