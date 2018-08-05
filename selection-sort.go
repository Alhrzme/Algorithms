package main

func selectionSort(arr []int) []int {
	for i, value := range arr {
		if i == len(arr) - 1 {
			break
		}

		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}

		if i != minIndex {
			arr[i] = arr[minIndex]
			arr[minIndex] = value
		}
	}

	return arr;
}