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
	ch2 := make(chan int, 10)
	var flag [100]bool
	go func() {
		for {
			ch2 <- 1
			index++
			wg.Add(1)
			go func(idx int) {
				flag[idx] = false
				for i := 1; i <= K; i++ {
					sum++
					//time.Sleep(time.Millisecond)
					fmt.Printf("\b\b\b%d", sum)
					//time.Sleep(time.Millisecond * 100)
					lock.Lock()

					// fmt.Println("sum:", sum)

					if sum >= M {
						ch <- i
						flag[idx] = true
						lock.Unlock()
						break
					}
					lock.Unlock()
				}
				//			fmt.Println("sum2:", sum)
				if flag[idx] {
					// fmt.Println("out")
					fmt.Println("the car-- ", idx, "Moved ", <-ch)
					// fmt.Println("______")
				} else {
					//				fmt.Println("sum1:", sum)
					fmt.Println("the car--", idx, " Moved 200")
				}
				wg.Done()
				<-ch
			}(index)
		}
	}()
	time.Sleep(time.Second)
	wg.Wait()
	close(ch)
}
