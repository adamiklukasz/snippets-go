package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"os"
	"strings"
)

//go:embed input.txt
var embeddedFile string

func scaninput(r io.Reader) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		s := scanner.Text()
		s2 := strings.Split(s, " ")

		fmt.Printf("s=%#v\n", s2)
	}
}

func scanFromString() {
	t := `
	heello 10 dfsdf
	sdfdsf as 32424
	sdfdsf 45 sdfdsf
	`

	scaninput(strings.NewReader(t[1:]))
}

func scanFromEmbeddedString() {
	scaninput(strings.NewReader(embeddedFile))
}

func scanFromStdIn() {
	scaninput(os.Stdin)
}

const filePath = "D:\\Workspace\\Go\\snippets-go\\io\\line-filtering-with-reader\\input.txt"

func scanFromFile() {
	f, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	fmt.Printf("err=%#v\n", err)
	scaninput(f)
}

func main() {
	scanFromString()
	//scanFromStdIn()
	scanFromFile()
	scanFromEmbeddedString()
}
