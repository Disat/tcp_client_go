package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8082")
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter message: ")
		scanner.Scan()
		input := scanner.Text()

		_, err := conn.Write([]byte(input))
		if err != nil {
			fmt.Println("Error sending message:", err)
			return
		}

		response := make([]byte, 1024)
		_, err = conn.Read(response)
		if err != nil {
			fmt.Println("Error receiving response:", err)
			return
		}
		fmt.Println("Server response:", string(response))
	}
}
