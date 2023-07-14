package main

import (
	"fmt"
	"time"
)

func main() {
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
	for i := 0; i < 20; i++ {
		time.Sleep(1*time.Second)
		fmt.Println(i,msg)
	}
}
