package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	var q int
	fmt.Scanf("%d%d\n", &n, &q)
	// fmt.Println(n, q)
	var a [25]int
	for i := 1; i <= n; i++ {
		fmt.Scanf("%d", &a[i])
	}

	// for i := 1; i <= n; i++ {
	// 	fmt.Println(a[i])
	// }
	sort.Ints(a[:])

	var num int

	for t := 1; t <= q; t++ {
		fmt.Scanf("\n%d", &num)
		var sum, tot int = 0, 0
		var flag bool = false
		for i := 24; i >= 24-n+1; i-- {
			sum += a[i]
			tot++
			if sum >= num {
				fmt.Println(tot)
				flag = true
				break
			}
		}
		if flag == false {
			fmt.Println("-1")
		}
	}
}
