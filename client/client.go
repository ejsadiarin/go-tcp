package client

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func StartTCPClient() {
	/***** client *****/
	// create socket
	clconn, err := net.Dial("tcp", "localhost:9777")
	if err != nil {
		fmt.Printf("\nError: %v", err)
		log.Fatal(err)
	}
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("Enter message: ")
		msg, err := reader.ReadString('\n')
		msg = strings.TrimSpace(msg)
		if err != nil {
			fmt.Printf("\nError: %v", err)
			log.Fatal(err)
		}
		_, err = clconn.Write([]byte(msg))
		if err != nil {
			fmt.Printf("\nError: %v", err)
			log.Fatal(err)
		}
	}
}
