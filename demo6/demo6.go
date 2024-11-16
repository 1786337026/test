package main

import (
	"fmt"
	"sync"
	"time"
)

var M, N, K, num, index, sum int
var wg sync.WaitGroup

func main() {
	M = 2000
	N = 10
	K = 200
	var wg sync.WaitGroup
	var lock sync.Mutex
	ch := make(chan int)
	var flag [30]bool
	for num <= 10 {
		num++
		index++
		wg.Add(1)
		go func(idx int) {
			flag[idx] = false
			for i := 1; i <= K; i++ {
				time.Sleep(time.Millisecond)
				lock.Lock()
				if sum >= M {
					sum += i
					ch <- i
					flag[idx] = true
					lock.Unlock()
					break
				}
				lock.Unlock()
			}
			if flag[idx] {
				fmt.Println("the car-- ", idx, "Moved ", <-ch)
				fmt.Println("______")
			} else {
				sum += 200
				fmt.Println("the car--", idx, " Moved 200")
			}
			wg.Done()
			num--
		}(index)
	}
	wg.Wait()
	close(ch)
}
