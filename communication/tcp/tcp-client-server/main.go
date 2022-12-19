package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func main() {
	go server()

	for i := 0; i < 10; i++ {
		go client()
	}

	select {}
}

func client() {
	dial, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		return
	}

	defer dial.Close()

	_, err = dial.Write([]byte("Hello world"))
	if err != nil {
		fmt.Printf("err=%#v\n", err)
	}

	dial.Close()
}

func server() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf("err=%#v\n", err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			fmt.Printf("err=%#v\n", err)
		}

		go func(c net.Conn) {
			conn.SetDeadline(time.Now().Add(10 * time.Second))
			defer conn.Close()

			sc := bufio.NewScanner(conn)
			for sc.Scan() {
				ln := sc.Text()
				fmt.Printf("[server] received=%#v\n", ln)
			}

			fmt.Printf("[server] connection closed\n")
		}(conn)
	}
}
