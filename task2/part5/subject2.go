package part5

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func DoCount2() {
	var counter int32
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddInt32(&counter, 1)
			}
		}()
	}
	wg.Wait()
	fmt.Println(atomic.LoadInt32(&counter))
}
