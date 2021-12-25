//
//	Selection sort algorithm..
//

package main

import "fmt"

type arrays []int

func main() {

	arr := arrays{}
	arr = []int{32, 21, 1, 0, 44, 32, 2, 21, 56, 6, 34, 23}

	arr.selectionSort()
	arr.printArr()
}

func (arr arrays) selectionSort() []int {

	for i := 0; i < len(arr); i++ {
		minIndex := i
		for j := i; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		// Swapping..
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
	return arr
}

func (arr arrays) printArr() {
	fmt.Println(arr)
}
