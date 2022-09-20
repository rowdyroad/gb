package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	appCtx, cancel := context.WithTimeout(context.Background(), time.Second)

	ctx := context.WithValue(appCtx, "key", "key")
	ctx2 := context.WithValue(ctx, "key2", "key2")

	fmt.Println(ctx2.Value("key"))

	go func() {
		time.Sleep(2 * time.Second)

		fmt.Println("Cancel by func")
		cancel()
	}()

	<-ctx2.Done()
}
