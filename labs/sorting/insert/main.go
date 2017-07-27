package main

import "fmt"

func insert(arr []int) {
	for i := 1; i < len(arr); i++ {
		value := arr[i]
		for j := i - 1; j >= 0; j-- {
			if value < arr[j] {
				arr[j+1], arr[j] = arr[j], value
			} else {
				break
			}
		}
	}
}

func main() {
	arr := []int{6, 5, 3, 2, 1, 9}
	fmt.Println(arr)
	insert(arr)
	fmt.Println(arr)
}
