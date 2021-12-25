package main

import "fmt"

type arrSeries []int

func main() {
	arr := arrSeries{12, 34, 45, 12, 1, 0}
	arr.bubbleSort()
	arr.print()
}
func (arr arrSeries) bubbleSort() []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
func (arr arrSeries) print() {
	fmt.Println(arr)
}
