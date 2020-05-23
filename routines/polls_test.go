package routines

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"
)

var counter int64
var chLimit = make(chan struct{}, 20)

func DoSth(i int) {
	chLimit <- struct{}{}
	atomic.AddInt64(&counter, 1)

	fmt.Printf("Routine %d %d\n", i, counter)
	time.Sleep(200 * time.Millisecond)

	atomic.AddInt64(&counter, -1)
	<-chLimit
}

func TestPolls(t *testing.T) {
	for i := 0; i < 100; i++ {
		go DoSth(i)
	}

	select {}
}
