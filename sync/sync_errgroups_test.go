package sync

import (
	"golang.org/x/sync/errgroup"

	"context"
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestErrGroup(t *testing.T) {
	gr, ctx := errgroup.WithContext(context.Background())

	for i := 0; i < 10; i++ {
		i := i
		gr.Go(func() error {
			time.Sleep(time.Duration(i) * time.Second)
			if i == 3 {
				fmt.Printf("Err %d\n", i)
				return errors.New("Error")
			}
			fmt.Printf("Norm %d\n", i)
			return nil
		})
	}

	go func() {
		select {
		case <-ctx.Done():
			fmt.Printf("Context done\n")
		}
	}()

	go func() {
		err := gr.Wait()
		fmt.Printf("Group done=%#v\n", err)
	}()

	select {}
}
