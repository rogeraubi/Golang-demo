package main 
import (
	"fmt"
)

func main() {
	var msg1 string
	messages :=make(chan string,2)


	go func() {
		messages <-"ping"
		messages <-"ping1"
	}()

	msg1 =<-messages
	// fmt.Println(msg1)

	msg1 =<-messages
   
	fmt.Println(msg1)
	fmt.Println(<-messages)
}