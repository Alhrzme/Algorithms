package main

import "fmt"

func main() {
	fmt.Println("ПЫЩ")
	arr := []int {14,4,3,2,10}
	arr = insertionSort(arr)

	for _, value := range arr {
		fmt.Println(value)
	}

}