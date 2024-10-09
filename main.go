package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// Use the flag package to parse command line arguments
	serverAddr := flag.String("server", "localhost:8080", "Address of the server to connect to")
	flag.Parse()

	// Step 1: Connect to the server
	fmt.Println("Attempting to connect to the server...")
	conn, err := net.Dial("tcp", *serverAddr)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}
	defer conn.Close()
	fmt.Println("Connected to server")

	// Step 2: Create a reader to handle server output
	serverReader := bufio.NewReader(conn)
	clientReader := bufio.NewReader(os.Stdin)

	// Step 3: Read and display server's request for username
	fmt.Println("Waiting for username prompt from server...")
	prompt, err := serverReader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading from server:", err)
		return
	}
	fmt.Print("Server prompt received: ", prompt)

	// Step 4: Read user input for username
	fmt.Print("Enter username: ")
	username, _ := clientReader.ReadString('\n')
	username = strings.TrimSpace(username)

	// Send the username to the server
	fmt.Println("Sending username to server...")
	conn.Write([]byte(username + "\n"))

	// Step 5: Read and display server's request for password
	fmt.Println("Waiting for password prompt from server...")
	prompt, err = serverReader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading from server:", err)
		return
	}
	fmt.Print("Server prompt received: ", prompt)

	// Step 6: Read user input for password
	fmt.Print("Enter password: ")
	password, _ := clientReader.ReadString('\n')
	password = strings.TrimSpace(password)

	// Send the password to the server
	fmt.Println("Sending password to server...")
	conn.Write([]byte(password + "\n"))

	// Step 7: Read the server's response (authentication result)
	response, err := serverReader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading response from server:", err)
		return
	}

	// Step 8: Display the server's response
	fmt.Println("Server response:", strings.TrimSpace(response))
}
