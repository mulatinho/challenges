package main

import (
	"fmt"
)

func roundGrade(num *int) {
	if *num >= 38 {
		multiple := (*num / 5) + 1
		diff := (multiple * 5) - *num

		if diff < 3 {
			*num = (multiple * 5)
		}
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	grades := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%d\n", &grades[i])
		roundGrade(&grades[i])
	}

	for j := 0; j < n; j++ {
		fmt.Println(grades[j])
	}
}
