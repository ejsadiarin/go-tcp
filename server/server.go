package server

import (
	"fmt"
	"log"
	"net"
)

func StartTCPServer() {
	/***** server *****/
	// create socket
	// accept socket connection request
	// once established, the server should listen...
	addr, err := net.ResolveTCPAddr("tcp", ":9777")
	if err != nil {
		fmt.Printf("\nError: %v", err)
		log.Fatal(err)
	}
	ln, err := net.ListenTCP("tcp", addr)
	if err != nil {
		fmt.Printf("\nError: %v", err)
		log.Fatal(err)
	}
	fmt.Println("Listening on port :9777")
	conn, err := ln.Accept()
	if err != nil {
		fmt.Printf("\nError: %v", err)
	}
	buf := make([]byte, 1024)
	defer ln.Close()
	for {
		// read
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("\nError: %v", err)
			log.Fatal(err)
		}
		fmt.Printf("\nReceived data %s from %s", string(buf[:n]), addr)
	}
}
