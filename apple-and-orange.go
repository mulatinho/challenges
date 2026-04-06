package main

import (
	"fmt"
)

func main() {
	var s, t, a, b, m, n int
	res := [2]int{0, 0}

	fmt.Scan(&s, &t)
	fmt.Scan(&a, &b)
	fmt.Scan(&m, &n)

	applesdist := make([]int, m)
	orangedist := make([]int, n)

	for i := 0; i < m; i++ {
		fmt.Scan(&applesdist[i])
		v := a + applesdist[i]

		if v >= s && v <= t {
			res[0]++
		}
	}

	for i := 0; i < n; i++ {
		fmt.Scan(&orangedist[i])
		v := b + orangedist[i]

		if v >= s && v <= t {
			res[1]++
		}
	}

	fmt.Printf("%d\n%d\n", res[0], res[1])
}
