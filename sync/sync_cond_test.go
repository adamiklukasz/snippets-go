package sync

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestSyncConditionals(t *testing.T) {
	cond := sync.NewCond(&sync.Mutex{})
	var m []int

	for i := 0; i < 5; i++ {
		go func(i int) {
			cond.L.Lock()

			for len(m) < 3 { // condition to check
				cond.Wait()
				fmt.Printf("try Enabled %d\n", i)
			}

			fmt.Printf("Enable %d\n", i)

			cond.L.Unlock()
		}(i)
	}

	go func() {
		for {
			time.Sleep(1000 * time.Millisecond)
			m = append(m, 0)
			//cond.Broadcast()
			cond.Signal()
		}
	}()

	select {}
}
