package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
			os.Exit(-1)
		}
	}()

	// create an dialer
	conn, err := net.Dial("tcp", "localhost:5000")
	check(err)

	// read buf from stdin
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name: ")
	clientName, _ := inputReader.ReadString('\n')
	trimmedName := strings.Trim(clientName, "\r\n")

	// send request to server
	for {
		fmt.Println("Enter data to send to server ( Type Q for quit ): ")
		input, _ := inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\r\n")
		if trimmedInput == "Q" {
			return
		}
		_, err := conn.Write([]byte(trimmedName + " Says: " + trimmedInput))
		check(err)
		fmt.Println("Mesage Sent Successfully.")
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
