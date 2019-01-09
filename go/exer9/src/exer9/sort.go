package exer9

import (
	// "fmt"
	"math/rand"
)

// Partition the slice arr around a random pivot (in-place), and return the pivot location.
func partition(arr []float64) int {
	// Adapted from https://stackoverflow.com/a/15803401/6871666
	left := 0
	right := len(arr) - 1

	// Choose random pivot
	pivotIndex := rand.Intn(len(arr))

	// Stash pivot at the right of the slice
	arr[pivotIndex], arr[right] = arr[right], arr[pivotIndex]

	// Move elements smaller than the pivot to the left
	for i := range arr {
		if arr[i] < arr[right] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	// Place the pivot after the last-smaller element
	arr[left], arr[right] = arr[right], arr[left]
	return left
}

func InsertionSort(arr []float64) {
	// TODO: implement insertion sort
	var x float64
	for i := 0; i < len(arr); i++ {
		x = arr[i]
		k := i
		for j := 0; j < i; j++ {
			if x < arr[j] {
				k = j
				j = i
			}
		}
		if i > 0 && i < len(arr) && k+1 < len(arr) && k+1 <= i {
			// fmt.Println(arr)
			copy(arr[k+1:i+1], arr[k:i])
			// copy(arr[1:2], arr[0:2])
			// fmt.Println(arr)
			arr[k] = x
			// fmt.Printf("k:%d,i:%d\n", k, i)
			// fmt.Println(arr)
		}

	}

}

const insertionSortCutoff = 10

func QuickSort(arr []float64) {
	// TODO: implement Quicksort:
	//   do nothing for length < 2
	//   do insertion sort for length < insertionSortCutoff
	//   do Quicksort otherwise.
	// TODO: decide on a good value for insertionSortCutoff
	// fmt.Printf("len:%d\n", len(arr))
	ln := len(arr)
	if ln > insertionSortCutoff {

		n := partition(arr)
		// fmt.Printf("n:%d\n", n)
		if n > 0 {
			QuickSort(arr[:n])
		}
		if n+1 < len(arr) {
			QuickSort(arr[n+1:])
		}

	} else if ln >= 2 && ln <= insertionSortCutoff {
		InsertionSort(arr)
	}
}
