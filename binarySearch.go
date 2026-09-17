package main

import "fmt"

func binarySearch(slice []int, target int) int {
	low := 0
	high := len(slice) - 1

	for low <= high {
		mid := (low + high) / 2
		guess := slice[mid]

		if guess == target {
			return mid
		}
		if guess < target {
			low = mid + 1
		}
		if guess > target {
			high = mid - 1

		}
	}

	return -1
}

func testBinarySearch() {
	// binarySearch
	slice := []int{1, 2, 4, 5}
	target := 3

	result := binarySearch(slice, target)

	if result != -1 {
		fmt.Println("Found target", target, "at index", result)
	} else {
		fmt.Println("Target not found")
	}
}
