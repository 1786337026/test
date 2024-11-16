package main

import (
	"fmt"
	"sync"
)

var M,N,K,num int
var wg sync.WaitGroup

func main(){
	M=2000
	N=10
	K=200
	var wg sync.WaitGroup
	var lock sync.Mutex
	wg.Add(1)
	for ;num<=10;{
		num++
		go func() {
			for {
				
			}
			lock.Lock()
			num--
			lock.Unlock()
		}()
	}
	wg.Wait()

}