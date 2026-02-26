package main

import (
	"fmt"
	"sort"
)

var binary string = "binary search algorithim"

func is_bina_search(n []int, t int) int {

	left := 0
	right := len(n) - 1

	for left <= right {

		mid := (left + right) / 2

		if n[mid] == t {
			return mid
		} else if t > n[mid] {
			left = mid + 1
		} else if t < n[mid] {
			right = mid - 1
		}
	}
	return -1
}

func main() {

	n := []int{3, 25, 10, 25, 8, 10, 47, 25}

	sort.Ints(n)

	t := is_bina_search(n, 10)

	fmt.Println(t)
}
