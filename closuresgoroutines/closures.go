package main 

import (
	"fmt"
_	"sync"
)
func main() {
	// var wg sync.WaitGroup
	done :=make(chan bool)
	values :=[]string{"a","b","c"}

    i :=0;
	for _, v:= range values { 
		go func(v string) {
			fmt.Println(v)
			// defer wg.Done()
			done <- true	
			}(v) 
			i++
			// wg.Add(1)
		}

		for _= range values {
		// for i := 0; i < 5; i++ {
			<-done
		}
		// }
		// wg.Wait()
}
