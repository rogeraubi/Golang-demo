package main

import (
	"context"
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)

	}
}
func main() {
	type favContextKey string
	f("direct")
	k := favContextKey("language")
	ctx := context.WithValue(context.Background(), k, "GO")

	go f("goroutines")

	go func(msg string, t context.Context) {
		for i := 0; i < 2; i++ {
			fmt.Println(msg, ":", i, t.Value(k))
			time.Sleep(time.Second)
		}
	}("going", ctx)
	time.Sleep(time.Second)
	fmt.Println("done")
}
