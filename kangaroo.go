package main

import (
	"fmt"
)

func main() {
	var x1, v1, x2, v2 int
	fmt.Scan(&x1, &v1, &x2, &v2)

	jump := 0
	match := 0

	for {
		if jump > 10000 {
			break
		}

		x1 += v1
		x2 += v2

		if x1 == x2 {
			match++
			break
		}

		//fmt.Printf("jump: %d, x1:%.2d, x2:%.2d\n", jump, x1, x2)
		jump++
	}

	if match > 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
