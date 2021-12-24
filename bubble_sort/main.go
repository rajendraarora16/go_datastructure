package main

import "fmt"

func main() {
	array := []int{11, 23, 44, 10, 2, 3, 43, 56, 60}

	fmt.Println(BubbleSort(array))
}

func BubbleSort(arr []int) []int {

	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-j; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	return arr
}
