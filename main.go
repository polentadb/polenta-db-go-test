package main

import (
	"fmt"
	"net"
)

func main() {
	// Connect to the server
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer conn.Close()

	// Send data to the server
	data := []byte("Hello, Server!")
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	response := make([]byte, 1024)
	n, err := conn.Read(response)
	if err != nil {
		return
	}
	fmt.Println(string(response[:n]))

	// Read and process data from the server
	// ...
}
