package main

import (
	"fmt"
)

func main() {
	var n, sum, a1, a2 int
	var a int
	fmt.Scanf("%d\n", &n)

	for i := 1; i <= n; i++ {
		fmt.Scanf("%d", &a)
		sum += a
	}
	a1 = sum - (sum/n)*n
	a2 = n - a1
	fmt.Println(a1 * a2)
}
