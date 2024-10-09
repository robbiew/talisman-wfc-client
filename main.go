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

	// Step 2: Connect to the server
	conn, err := net.Dial("tcp", *serverAddr)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}
	defer conn.Close()

	// Get username and password input from the user
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Print("Enter Password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	// Send the username and password to the server
	conn.Write([]byte(username + "\n"))
	conn.Write([]byte(password + "\n"))

	// Step 5: Read the server's response
	serverReader := bufio.NewReader(conn)
	response, _ := serverReader.ReadString('\n')

	// Print the server's response (success or failure)
	fmt.Println("Server response:", strings.TrimSpace(response))
}
