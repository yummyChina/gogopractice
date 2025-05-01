package part4

import (
	"fmt"
	"sync"
)

func Comu() {
	var wg sync.WaitGroup
	wg.Add(2)
	ch := make(chan int)
	go func() {
		defer wg.Done()
		defer close(ch)
		for i := 1; i < 10; i++ {
			ch <- i
			fmt.Println("send:", i)
		}
	}()

	go func() {
		defer wg.Done()
		for num := range ch {
			fmt.Println("receive:", num)
		}
	}()
	wg.Wait()

}
