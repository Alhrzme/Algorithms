package main


func insertionSort(arr []int) []int {
	for i, value := range arr {
		currentMax := value
		j := i - 1
		for ; j >= 0 && arr[j] > currentMax; j-- {
			arr[j + 1] = arr[j]
		}
		arr[j + 1] = currentMax
	}

	return arr
}