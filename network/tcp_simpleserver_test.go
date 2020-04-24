package network

import (
	"bufio"
	"fmt"
	"net"
	"testing"
	"time"
)

func TestTCPSimpleServer(t *testing.T) {
	go startTCPServer()
	time.Sleep(100 * time.Millisecond)
	go startTCPClient("cli 1")
	go startTCPClient("cli 2")

	select {}
}

func startTCPServer() {
	ln, _ := net.Listen("tcp", ":8080")
	for {
		conn, err := ln.Accept()
		fmt.Printf("err=%#v\n", err)

		go func(conn net.Conn) {
			fmt.Printf("Handle connection...\n")
			for {
				str, _ := bufio.NewReader(conn).ReadString('\n')
				fmt.Printf("Received %#v\n", string(str))
			}
		}(conn)
	}
}

func startTCPClient(text string) {
	conn, _ := net.Dial("tcp", "127.0.0.1:8080")

	i := 1
	for {
		i++
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("Sending...\n")
		_, _ = fmt.Fprint(conn, fmt.Sprintf("%s %d\n", text, i))
	}
}
