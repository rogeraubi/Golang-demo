package main

import (
	"fmt"
	"time"
	"context"
)

const shortDuration = 1 * time.Millisecond
func main() {
	d := time.Now().Add(shortDuration)
	ctx1, cancel1 := context.WithDeadline(context.Background(), d)
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx1.Done():
		fmt.Println(ctx1.Err())
	}

	defer cancel1()
		ctx, cancel := context.WithCancel(context.Background())
	    defer cancel()
	    for n:= range gen(ctx) {
		fmt.Println(n)
		if n==5 {
			break
		}
	}	
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()
	select {
	case res := <-c1:
		{
			c11 := make(chan string, 1)
			fmt.Println(res)
			go func() {
				fmt.Println("this is another test channel")
				c11 <- "done"
			}()
			testLog("ok-ok")
			<-c11
			// testLog("late-ok")
			testLog("later-ok")
		}
	case <-time.After(3 * time.Second):fmt.Println("timeout 1")
	}
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}

}
func testLog(msg string) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for n:= range gen(ctx) {
		fmt.Println(n,msg)
		if n==10 {
			break
		}
}
}
func gen(ctx context.Context) <-chan int {
	dst := make(chan int) 
	n :=1 
	go func() {
		for {
			select {
			case <-ctx.Done():
				return 
			case dst <-n:
				n++	
			}
		}
	}()
		return dst
	}
