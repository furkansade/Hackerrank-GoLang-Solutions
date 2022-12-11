package main

import "fmt"

func main() {
	var n, k, t, sum, b int

	fmt.Scan(&n, &k) // 4 1
	for i := 0; i < n; i++ {
		fmt.Scan(&t) // 3 10 2 9
		if i != k {
			sum += t
		}
	}

	fmt.Scan(&b) // 12
	if b == sum/2 {
		fmt.Println("Bon Appetit")
	} else {
		fmt.Println(b - (sum / 2)) // 5
	}

}
