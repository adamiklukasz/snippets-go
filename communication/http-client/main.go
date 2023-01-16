package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	g, err := http.Get("https://gobyexample.com/http-client")
	fmt.Printf("err=%#v\n", err)

	os.Stdout.ReadFrom(g.Body)
	g.Body.Close()
}
