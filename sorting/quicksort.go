package main

import (
	"fmt"
)

func main() {
	unsortedArray := []uint{444, 2, 2, 5, 21, 131, 523, 3, 31, 5, 51, 924, 3413, 31, 24, 64, 24, 13, 775, 427, 55}
	//sort.Slice(unsortedArray, func(i, j int) bool {
	//	return unsortedArray[i] < unsortedArray[j]
	//})
	//fmt.Printf("%v\n", unsortedArray)

	quicksort(unsortedArray)

	fmt.Printf("%v\n", unsortedArray)

}

func quicksort(arr []uint) {
	fmt.Printf("%v\n", arr)

	if len(arr) <= 1 {
		return
	}

	pindex := len(arr) - 1
	pivot := arr[pindex]

	for i, j := 0, pindex-1; i < j; {
		if arr[i] <= pivot {
			i++
		} else if arr[j] >= pivot {
			j--
		}

		if arr[i] >= pivot && arr[j] <= pivot {
			swap(arr, i, j)
		}

		if i == j && arr[j] >= pivot {
			swap(arr, j, pindex)
			pindex = j
		}
	}

	left := arr[:pindex]
	right := arr[pindex:]
	quicksort(left)
	quicksort(right)

}

func swap(arr []uint, i int, j int) {
	buffer := arr[i]
	arr[i] = arr[j]
	arr[j] = buffer
}
