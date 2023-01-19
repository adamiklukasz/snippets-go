package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	go server()
	time.Sleep(1 * time.Second)
	//go client()

	select {}
}

func client() {
	conn, resp, err := websocket.DefaultDialer.Dial("ws://127.0.0.1:8080", nil)
	if err != nil {
		fmt.Printf("err=%#v\n", err.Error())
	}

	fmt.Printf("conn=%#v\n", conn)
	fmt.Printf("resp=%#v\n", resp)

	for i := 0; i < 10; i++ {
		err := conn.WriteMessage(websocket.TextMessage, []byte("Hello world"))
		if err != nil {
			fmt.Printf("err=%#v\n", err)
		}
		time.Sleep(1 * time.Second)
	}

	conn.Close()

}

type handler int

func (h handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	u := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	c, err := u.Upgrade(writer, request, nil)
	if err != nil {
		fmt.Printf("err=%#v\n", err.Error())
	}

	c.WriteMessage(websocket.TextMessage, []byte("Hello world"))

	for {
		_, msg, err := c.ReadMessage()
		if err != nil {
			fmt.Printf("err=%#v\n", err.Error())
			return
		}

		fmt.Printf("mes=%#v\n", string(msg))
	}

}

func server() {
	var h handler
	http.ListenAndServe("127.0.0.1:8080", h)
}
