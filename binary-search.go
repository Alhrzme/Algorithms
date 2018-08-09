package main

import (
	"math"
)

func binarySearch(arr []int, searched int, firstElIndex int, lastElIndex int) int {
	if firstElIndex == lastElIndex {
		if arr[firstElIndex] == searched {
			return firstElIndex
		}
		return -1
	}

	middleIndex := int(math.Ceil((float64(lastElIndex) + float64(firstElIndex))/ 2))
	middleVal := arr[middleIndex]

	if middleVal > searched {
		lastElIndex = middleIndex - 1
	} else {
		firstElIndex = middleIndex
	}

	return binarySearch(arr, searched, firstElIndex, lastElIndex)
}
