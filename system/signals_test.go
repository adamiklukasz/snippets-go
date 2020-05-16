package system

import (
	_ "bytes"
	"fmt"
	"os"
	"os/signal"
	"testing"
)

func TestSignals(t *testing.T) {

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill)

	for i := 0; i <= 1024; i++ {
		select {
		case s := <-c:
			fmt.Printf("Singal %s\n", s.String())
			return
		default:
		}
	}
}
