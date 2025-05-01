package part2

import (
	"fmt"
	"sync"
	"time"
)

func TwoRoutine() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i += 2 {
			fmt.Println("奇数:", i)
			time.Sleep(800 * time.Millisecond)
		}
	}()

	go func() {
		// defer wg.Done()
		for i := 2; i <= 10; i += 2 {
			fmt.Println("偶数:", i)
			time.Sleep(1 * time.Second)
		}
	}()

	wg.Wait()

}
