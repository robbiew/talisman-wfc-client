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
	// Define command-line flags
	serverAddr := flag.String("server", "", "Address of the server to connect to (e.g., localhost:8080)")

	// Parse command-line flags
	flag.Parse()

	// Ensure that the required flags are set
	if *serverAddr == "" {
		fmt.Println("Error: --server flag is required")
		flag.Usage()
		os.Exit(1)
	}

	// Connect to the server
	conn, err := net.Dial("tcp", *serverAddr)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}
	defer conn.Close()

	// Create readers for the server and user input
	serverReader := bufio.NewReader(conn)
	clientReader := bufio.NewReader(os.Stdin)

	// Step 1: Handle server's username and password prompts
	for {
		// Read the server's message (username or password prompt)
		prompt, err := serverReader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from server:", err)
			return
		}
		fmt.Print(prompt)

		// If the server sends "Authentication successful", break out of this loop and continue
		if strings.Contains(prompt, "Authentication successful") {
			break
		}

		// Get user input for the username or password
		userInput, _ := clientReader.ReadString('\n')
		userInput = strings.TrimSpace(userInput)

		// Send the user input to the server
		conn.Write([]byte(userInput + "\n"))
	}

	// Step 2: Continuously read updates from the server after authentication
	for {
		update, err := serverReader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading updates from server:", err)
			break
		}
		fmt.Print(update)
	}
}
