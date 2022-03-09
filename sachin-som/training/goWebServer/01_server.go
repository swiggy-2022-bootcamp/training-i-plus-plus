package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("starting the server...")

	// Listen to request
	listerner, err := net.Listen("tcp", "localhost:5000")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		conn, err := listerner.Accept()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		go doServerStuff(conn)
	}
}

func doServerStuff(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading: ", err.Error())
			return // terminate program
		}
		fmt.Printf("Received data:\n %v\n", string(buf))
	}
}
