package main

import (
	"fmt"
	"sort"
)

func main() {

	arr := []int{2, 1, 4, 76, 765, 43, 22, 10}
	fmt.Println(BinarySearch(arr, 10))
}

func BinarySearch(arr []int, searchVal int) bool {

	low := 0
	high := len(arr) - 1

	sort.Ints(arr)

	for low <= high {
		median := (low + high) / 2

		if arr[median] < searchVal {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(arr) || arr[low] != searchVal {
		return false
	}

	return true
}
