package main

import (
	"golang.org/x/sync/errgroup"

	"context"
	"errors"
	"fmt"
	"time"
)

func main() {
	errGroup, ctxErrGroup := errgroup.WithContext(context.Background())

	for i := 0; i < 10; i++ {
		i := i
		errGroup.Go(func() error {
			time.Sleep(time.Duration(i) * time.Second)
			if i == 3 {
				fmt.Printf("Err %d\n", i)
				return errors.New("Error")
			}

			if ctxErrGroup.Err() != nil {
				fmt.Printf("No point to continue %d\n", i)
				return nil
			}

			fmt.Printf("Normal execution %d\n", i)
			return nil
		})
	}

	go func() {
		select {
		case <-ctxErrGroup.Done():
			fmt.Printf("Context of errorGroup done\n")
		}
	}()

	go func() {
		err := errGroup.Wait()
		fmt.Printf("Group done=%#v\n", err)
	}()

	select {}
}
