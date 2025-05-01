package part4

import (
	"fmt"
	"sync"
)

func Comu1() {
	var wg sync.WaitGroup
	wg.Add(2)
	ch := make(chan int, 5)
	go func() {
		defer wg.Done()
		defer close(ch)
		for i := 1; i < 100; i++ {
			fmt.Println("send:", i)
			ch <- i

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
