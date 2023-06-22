package main
import (
	"fmt"
	"sync"
	"time"
)
func worker(id int) {
	fmt.Print("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}
func main() {
   var wg sync.WaitGroup
   for i:=1; i<=5; i++ {
	wg.Add(1)
	j:= i
	go func(j int) {
		defer wg.Done()
		worker(j)
	}(j)
   }
   wg.Wait()
}