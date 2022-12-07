package main

import "fmt"

func Solution(n int, p int) {

	front := p / 2
	back := (n - p) / 2

	if n%2 == 0 {
		back = (n + 1 - p) / 2
	}

	if front < back {
		fmt.Println(front)
	} else {
		fmt.Println(back)
	}
}

func main() {
	var n, p int
	fmt.Scan(&n, &p)
	Solution(n, p)
}
