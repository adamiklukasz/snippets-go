package main

import (
	"embed"
	_ "embed"
	"fmt"
	"io"
)

//go:embed embDir/*.*
//go:embed embDir/subDir/*.*
var folder embed.FS

func handle(p string) {
	f, err := folder.Open(p)
	fmt.Printf("err=%#v\n", err)

	b, err := io.ReadAll(f)
	fmt.Printf("err=%#v\n", err)
	fmt.Printf("b=%#v\n", string(b))
}

func main() {
	handle("embDir/f1.txt")
	handle("embDir/subDir/f3.txt")
}
