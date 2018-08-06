package main

import (
	"math"
	//"fmt"
)

func mergeSort(arr []int) []int {
	if len(arr) > 1 {
		midIndex := int(math.Floor(float64(len(arr)) / 2.0))

		var firstHalf []int
		var secondHalf []int
		for i := 0; i < len(arr); i++ {
			if i < midIndex {
				firstHalf = append(firstHalf, arr[i])
			} else {
				secondHalf = append(secondHalf, arr[i])
			}
		}

		firstHalf = mergeSort(firstHalf)
		secondHalf = mergeSort(secondHalf)

		firstHalfIndex := 0
		secondHalfIndex := 0
		for i := 0; i < len(arr); i++ {
			if  firstHalfIndex < len(firstHalf) && (secondHalfIndex >= len(secondHalf) || firstHalf[firstHalfIndex] <= secondHalf[secondHalfIndex]) {
				arr[i] = firstHalf[firstHalfIndex]
				firstHalfIndex++
			} else {
				arr[i] = secondHalf[secondHalfIndex]
				secondHalfIndex++
			}
		}
	}

	return arr
}
