// Binary search algorithm in golang

package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {

	var searchVal string
	arr := []int{23, 3, 12, 4, 5, 1, 6, 34, 22, 2, 65, 3}
	sort.Ints(arr)
	fmt.Println("Array:", arr)
	// Reading input value from terminal....
	fmt.Println("Enter your input value to search:")
	_, err := fmt.Scanln(&searchVal)
	if err != nil {
		fmt.Println("There is an error:", err)
		os.Exit(1)
	}

	// Converting string value to integer value...
	// Notice here the base "10"..
	//
	// We can check typeOf via fmt.Printf("TypeOf: %T", value)

	intInputVal, err := strconv.ParseInt(searchVal, 10, 64)
	if err != nil {
		fmt.Println("There is an error with ParseInt:", err)
		os.Exit(1)
	}

	valExist, pos := binarySearch(arr, int(intInputVal))

	fmt.Println("Value exists:", valExist, "at position:", pos)
}

func binarySearch(arr []int, searchVal int) (bool, int) {
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
		return false, 0
	}

	return true, low
}
