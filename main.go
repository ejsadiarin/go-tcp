package main

import (
	"bufio"
	"fmt"
	"go-tcp/client"
	"go-tcp/server"
	"log"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	fmt.Printf("What to start (1 - client, 2 - server): ")
	choice, err := r.ReadString('\n')
	if err != nil {
		fmt.Printf("\nError: %v", err)
		log.Fatal(err)
	}
	choice = strings.TrimSpace(choice)
	switch choice {
	case "1":
		client.StartTCPClient()
	case "2":
		server.StartTCPServer()
	default:
		fmt.Println("Internal server error")
		log.Fatal("Internal server error")
	}
}
