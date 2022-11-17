package main

import "fmt"

func main() {
	var n, k, q int
	fmt.Scan(&n, &k, &q)

	nums := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}

	k = k % n

	for i := 0; i < q; i++ {
		result := 0
		fmt.Scan(&result)
		if (result - k) < 0 {
			fmt.Println(nums[result+n-k])
		} else {
			fmt.Println(nums[result-k])
		}
	}
}
