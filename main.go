package main

import "fmt"

func main() {
	fmt.Println("ПЫЩ")
	arr := []int {14,4,3,2,10}
	arr = mergeSort(arr)

	for _, value := range arr {
		fmt.Println(value)
	}

}